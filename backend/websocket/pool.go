package websocket

import (
	"bomberman-dom/game"
	"fmt"
	"strconv"
	"time"
)

type Pool struct {
	Register     chan *Client
	Unregister   chan *Client
	Clients      map[*Client]bool
	Broadcast    chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:     make(chan *Client),
		Unregister:   make(chan *Client),
		Clients:      make(map[*Client]bool),
		Broadcast:    make(chan Message),
	}
}

func (pool *Pool) createPlayers() []game.Player {
	keys := make([]game.Player, len(pool.Clients))

	i := 0
	for client := range pool.Clients {
		keys[i] = game.CreatePlayer(client.ID, i)
		i++
	}

	return keys
}

func (pool *Pool) Start(gameState *game.GameState) {
	// listen for every message in every pools channel
	for {
		// select will execute whichever channel sends data first
		// channels can only send data when they have received it
	S:
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true

			gameState.Map = game.CreateBaseMap()
			gameState.Players = pool.createPlayers()

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
				gameState.Map = game.CreateBaseMap()
				gameState.Players = pool.createPlayers()
			case "PLAYER_MOVE": //update player movement
				player.Move(message.Body)
			case "PLAYER_DROPPED_BOMB" : 
				if player.BombsLeft <= 0 {
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
				gameState.CheckIfSomebodyDie(&explosion) 
				if len(destroyedBlocks)!=0{
					go func() { //trigger map update
							time.Sleep(600 * time.Millisecond)
							message.Type = "MAP_UPDATE"	
							message.GameState.Map = game.DestroyBlocks(gameState.Map, destroyedBlocks)
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
			case "MAP_UPDATE":
				gameState.Map = message.GameState.Map
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
