package websocket

import (
	"bomberman-dom/game"
	"fmt"
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

func (pool *Pool) createPlayers() []game.Player {
	keys := make([]game.Player, len(pool.Clients))

	i := 0
	for client := range pool.Clients {
		keys[i] = game.CreatePlayer(client.ID, i)
		// keys[i] = game.Player{X: 0, Y: 0, Name: client.ID, Speed:1}
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
			gameState.Players = pool.createPlayers()

			for otherClient := range pool.Clients {
				if client.ID == otherClient.ID {
					err := otherClient.Conn.WriteJSON(Message{Type: "JOIN_QUEUE", GameState: gameState})
					if err != nil {
						fmt.Println("JSON MARSHAL ERR", err)
					}
				} else {
					otherClient.Conn.WriteJSON(Message{Type: "NEW_USER", GameState: gameState})
				}
			}
			break S
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			gameState.Players = pool.createPlayers()

			// fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: "USER_LEFT", GameState: gameState})
			}
			break S
		case message := <-pool.Broadcast:
			// fmt.Println("Sending message to all clients in Pool")

			if message.Type == "PLAYER_MOVE" {
				currentPlayerIndex := gameState.FindPlayer(message.Creator)
				gameState.Players[currentPlayerIndex].Move( message.Body )
			}

			message.GameState = gameState

			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}

		// fmt.Println("Size of Connection Pool: ", len(pool.Clients))

	}
}
