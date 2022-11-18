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

	var gameState *game.GameState

	for {
	S:
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true

			gameState = game.NewGame()
			gameState.Map = game.CreateBaseMap()
			gameState.Players = pool.createPlayers(gameState.Map)

			for otherClient := range pool.Clients {
				if client.ID == otherClient.ID {
					otherClient.Conn.WriteJSON(Message{Type: "START_GAME", Body: "me", GameState: gameState, Creator: client.ID})
				}
				otherClient.Conn.WriteJSON(Message{Type: "START_GAME", Body: strconv.Itoa(len(pool.Clients)), GameState: gameState, Creator: client.ID})
			}

			break S
		case client := <-pool.Unregister:
			delete(pool.Clients, client)

			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: "USER_LEFT", Body: strconv.Itoa(len(pool.Clients))})
			}
			break S
		case message := <-pool.Broadcast:
			currentPlayerIndex := gameState.FindPlayer(message.Creator)
			player := &gameState.Players[currentPlayerIndex]

			switch message.Type {
			case "START_GAME":
				gameState = game.NewGame()
				gameState.Map = game.CreateBaseMap()
				gameState.Players = pool.createPlayers(gameState.Map)
			case "PLAYER_MOVE":
				if !player.IsAlive() {
					break S
				}
				//update player movement
				player.Move(message.Body)
				lostLive := gameState.CheckIfPlayerDied(player)

				if lostLive {
					//let reborn
					go message.MonstersReborn(pool, gameState, []int{currentPlayerIndex})
				}
				pickedUpPowerUp := player.PickedUpPowerUp(&gameState.PowerUps)

				if pickedUpPowerUp {
					message.Body = "PICKED_UP_POWERUP"
				}
			case "PLAYER_DROPPED_BOMB":
				if player.BombsLeft <= 0 || !player.IsAlive() {
					break S
				}
				player.DropBomb()

				go message.BombExploded(pool)

			case "BOMB_EXPLODED":
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
				player.ExplosionComplete()

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
