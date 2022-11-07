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

			// fmt.Println("Sending message to all clients in Pool", message.Type)

			if message.Type == "PLAYER_MOVE" {
				fmt.Println(message.Creator)
				currentPlayerIndex := gameState.FindPlayer(message.Creator)
				gameState.Players[currentPlayerIndex].Move(message.Body)
			} else if message.Type == "START_GAME" {
				gameState.Map = game.CreateBaseMap()
				gameState.Players = pool.createPlayers()
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
