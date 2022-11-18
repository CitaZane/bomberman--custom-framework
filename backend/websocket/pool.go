package websocket

import (
	"bomberman-dom/game"
	"fmt"
	"strconv"
	"time"
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
				} // allow player to move only if has not died

				//update player movement
				player.Move(message.Body)

				lostLive := gameState.CheckIfPlayerDied(player)
				if lostLive {
					go func() {
						time.Sleep(3000 * time.Millisecond)
						message.Type = "PLAYER_REBORN"
						gameState.Players[currentPlayerIndex].Movement = game.RightStop
						pool.Broadcast <- message
					}()
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
				go func() {
					time.Sleep(3000 * time.Millisecond)
					message.Type = "BOMB_EXPLODED"
					pool.Broadcast <- message
				}()

			case "BOMB_EXPLODED":
				destroyedBlocks, explosion := player.MakeExplosion(gameState.Map)
				player.BombExplosionComplete()
				monstersLostLives := gameState.CheckIfSomebodyDied(&explosion)

				if len(monstersLostLives) != 0 {
					go func() {
						time.Sleep(3000 * time.Millisecond)
						message.Type = "PLAYER_REBORN"
						for _, i := range monstersLostLives { //reset the movement
							gameState.Players[i].Movement = game.RightStop
						}

						pool.Broadcast <- message
					}()
				}
				if len(destroyedBlocks) != 0 {
					go func() { //trigger map update
						time.Sleep(1000 * time.Millisecond)
						message.Type = "MAP_UPDATE"
						gameState.Map = game.DestroyBlocks(gameState.Map, destroyedBlocks)

						// check if destroyed block index match with powerup block index
						for _, blockIndex := range destroyedBlocks {
							for _, powerUp := range game.GeneratedPowerUps {
								if blockIndex == powerUp.Tile {
									// fmt.Println("blockIndex", blockIndex)
									// fmt.Println("powerUp", powerUp)
									gameState.PowerUps = append(gameState.PowerUps, powerUp)
								}
							}
						}

						pool.Broadcast <- message
					}()
				}
				go func() { //trigger end of explosion
					time.Sleep(900 * time.Millisecond)
					message.Type = "EXPLOSION_COMPLETED"
					pool.Broadcast <- message
				}()

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
