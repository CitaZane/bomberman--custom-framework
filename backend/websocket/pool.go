package websocket

import (
	"bomberman-dom/game"
	"fmt"
	"strconv"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) createPlayers(gameMap []int) []game.Player {
	keys := make([]game.Player, len(pool.Clients))

	i := 0
	for client := range pool.Clients {
		keys[i] = game.CreatePlayer(client.ID, i, gameMap)
		i++
	}

	return keys
}

func (pool *Pool) Start() {

	var gameState = game.NewGame()

	for {
	S:
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true

			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: "JOIN_QUEUE", Body: strconv.Itoa(len(pool.Clients))})
			}

		case client := <-pool.Unregister:
			delete(pool.Clients, client)

			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: "USER_LEFT", Body: strconv.Itoa(len(pool.Clients))})
			}
		case message := <-pool.Broadcast:

			/*
				had to move these because before the "START_GAME" message we dont have players yet
			*/

			// currentPlayerIndex := gameState.FindPlayer(message.Creator)
			// 	player := &gameState.Players[currentPlayerIndex]

			switch message.Type {
			case "START_GAME":
				gameState.Map = game.CreateBaseMap()
				gameState.Players = pool.createPlayers(gameState.Map)

			case "PLAYER_MOVE":
				currentPlayerIndex := gameState.FindPlayer(message.Creator)
				player := &gameState.Players[currentPlayerIndex]
				if !player.IsAlive() || gameState.State != game.Play {
					break S
				}
				//update player movement
				player.Move(message.Body, message.Delta)

				if lostLive := gameState.CheckIfPlayerDied(player); lostLive {
					go message.MonstersReborn(pool, gameState, []int{currentPlayerIndex})
				}
				if pickedUpPowerUp := player.PickedUpPowerUp(&gameState.PowerUps); pickedUpPowerUp {
					message.Body = "PICKED_UP_POWERUP"

				}

			case "PLAYER_DROPPED_BOMB":
				currentPlayerIndex := gameState.FindPlayer(message.Creator)
				player := &gameState.Players[currentPlayerIndex]

				if player.BombsLeft <= 0 || !player.IsAlive() || player.Invincible{
					break S
				}
				player.DropBomb()

				go message.BombExploded(pool)

			case "BOMB_EXPLODED":
				currentPlayerIndex := gameState.FindPlayer(message.Creator)
				player := &gameState.Players[currentPlayerIndex]

				destroyedBlocks, explosion := player.MakeExplosion(gameState.Map)
				player.BombExplosionComplete()
				monstersLostLives := gameState.CheckIfSomebodyDied(&explosion)
				//trigger monster-reborn
				go message.MonstersReborn(pool, gameState, monstersLostLives)
				//trigger map update
				go message.UpdateMap(pool, gameState, destroyedBlocks)
				//trigger end of explosion
				go message.ExplosionComplete(pool)

			case "EXPLOSION_COMPLETED":
				currentPlayerIndex := gameState.FindPlayer(message.Creator)
				player := &gameState.Players[currentPlayerIndex]

				player.ExplosionComplete()
			case "PLAYER_AUTO_MOVE":
				currentPlayerIndex := gameState.FindPlayer(message.Creator)
				player := &gameState.Players[currentPlayerIndex]

				//update player movement wthouth obstacles
				done := player.AutoMove(message.Body)
				message.Type = "PLAYER_MOVE"
				if !done {
					go message.AutoGuideWinner(pool, message.Creator)
				} else {
					gameState.State = game.Finish
				}
			}

			if gameState.State == game.GameOver {
				message.ActivateGameOverScreen(pool, gameState)
				gameState.State = game.WalkOfFame
				var winner = gameState.FindWinner()
				go message.AutoGuideWinner(pool, gameState.Players[winner].Name)
			}

			message.GameState = gameState
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}

		}

	}
}
