package websocket

import (
	"bomberman-dom/game"
	"fmt"
	"strconv"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    []*Client
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    []*Client{},
		Broadcast:  make(chan Message),
	}
}
func (pool *Pool) RemoveClient(clientGoingAway *Client) {
	for i, client := range pool.Clients {
		if client.ID == clientGoingAway.ID {
			pool.Clients = append(pool.Clients[:i], pool.Clients[i+1:]...)
		}
	}
}

func (pool *Pool) createPlayers() []game.Player {
	keys := []game.Player{}

	i := 0
	for _, client := range pool.Clients {
		keys = append(keys, game.CreatePlayer(client.ID, i))
		if i == 3 {
			break
		} //max 4 players
		i++
	}

	return keys
}
func (pool *Pool) UsernameUnique(username string) bool {
	for _, client := range pool.Clients {
		if client.ID == username {
			return false
		}
	}
	return true
}

func (pool *Pool) Start() {

	var gameState = game.NewGame()
	var playerNames = make([]string, 0) // playerNames is for sending player names to lobby without creating actual players

	for {
	S:
		select {
		case client := <-pool.Register:
			pool.Clients = append(pool.Clients, client)
			if gameState.State == game.Lobby {

				playerNames = append(playerNames, client.ID)

				message := Message{Type: "JOIN_QUEUE", PlayerNames: playerNames}
				for _, client := range pool.Clients {
					client.Conn.WriteJSON(message)
				}
			} else {
				message := Message{
					Type:      "JOIN_SPECTATOR",
					Body:      strconv.Itoa(len(pool.Clients) - len(gameState.Players)), //-th in next queue
					GameState: gameState,
				}
				client.Conn.WriteJSON(message)
			}

		case client := <-pool.Unregister:
			message := Message{}
			if gameState.State != game.Lobby && gameState.IsPlayer(client.ID) {
				currentPlayerIndex := gameState.FindPlayer(client.ID)
				go message.PlayerLeftGame(pool, currentPlayerIndex, gameState)
			} else {
				message.Type = "USER_LEFT"
				message.Body = strconv.Itoa(len(pool.Clients) - 1)
				for i, name := range playerNames {
					if name == client.ID {
						playerNames = append(playerNames[:i], playerNames[i+1:]...)
					}
				}
				message.PlayerNames = playerNames

			}
			pool.RemoveClient(client)

			for _, client := range pool.Clients {
				client.Conn.WriteJSON(message)
			}
		case message := <-pool.Broadcast:

			if gameState.State == game.Lobby {
				if message.Type == "START_GAME" {
					gameState.StartGame()
					gameState.Players = pool.createPlayers()
				}
			} else {
				currentPlayerIndex := gameState.FindPlayer(message.Creator)
				if !gameState.IsPlayer(message.Creator) { //this is a watcher do not register moves
					break S
				}
				player := &gameState.Players[currentPlayerIndex]
				switch message.Type {
				case "PLAYER_MOVE":
					if !player.IsAlive() || gameState.State != game.Play {
						break S
					}
					//update player movement
					player.Move(message.Body, message.Delta, gameState)

					if lostLive := gameState.CheckIfPlayerDied(player); lostLive {
						go message.MonstersReborn(pool, gameState, []int{currentPlayerIndex})
					}
					if pickedUpPowerUp := player.PickedUpPowerUp(&gameState.PowerUps); pickedUpPowerUp {
						message.Body = "PICKED_UP_POWERUP"
					}

				case "PLAYER_DROPPED_BOMB":
					if player.BombsLeft <= 0 || !player.IsAlive() || player.Invincible {
						break S
					}
					player.DropBomb()

					go message.BombExploded(pool)

				case "BOMB_EXPLODED":
					destroyedBlocks, explosion := player.MakeExplosion(gameState.Map)
					player.BombExplosionComplete()
					monstersLostLives := gameState.CheckIfSomebodyDied(&explosion)
					//trigger side effects
					go message.MonstersReborn(pool, gameState, monstersLostLives)
					go message.UpdateMap(pool, gameState, destroyedBlocks)
					go message.ExplosionComplete(pool)

				case "EXPLOSION_COMPLETED":
					player.ExplosionComplete()
				case "PLAYER_AUTO_MOVE":
					//update player movement wthouth obstacles
					done := player.AutoMove(message.Body)
					message.Type = "PLAYER_MOVE"
					if !done {
						go message.AutoGuideWinner(pool, message.Creator)
					} else {
						gameState.State = game.Finish
					}

				}

			}
			if gameState.State == game.GameOver {
				message.ActivateGameOverScreen(pool, gameState)
				gameState.State = game.WalkOfFame
				go message.AutoGuideWinner(pool, gameState.FindWinner())
			}
			if gameState.State == game.Finish {
				message.Type = "FINISH"
				gameState.FinishGame()
			}
			message.GameState = gameState
			for _, client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println("WEBSOCKET ERROR: ", err)
					gameState.FinishGame()
				}
			}

		}

	}
}
