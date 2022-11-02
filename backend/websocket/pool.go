package websocket

import (
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

func (pool *Pool) getClientNames() []string {
	keys := make([]string, len(pool.Clients))

	i := 0
	for client := range pool.Clients {
		keys[i] = client.ID
		i++
	}

	return keys
}

func (pool *Pool) Start() {
	// listen for every message in every pools channel
	for {
		// select will execute whichever channel sends data first
		// channels can only send data when they have received it
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			game.Players = pool.getClientNames()

			for otherClient := range pool.Clients {
				if client.ID == otherClient.ID {
					err := otherClient.Conn.WriteJSON(Message{Type: "INIT_ROOM", GameState: game})
					if err != nil {
						fmt.Println("JSON MARSHAL ERR", err)
					}
				} else {
					otherClient.Conn.WriteJSON(Message{Type: "NEW_USER", GameState: game})
				}
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			game.Players = pool.getClientNames()

			// fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: "USER_LEFT", GameState: game})
			}
			break
		case message := <-pool.Broadcast:
			// fmt.Println("Sending message to all clients in Pool")
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}

		fmt.Println("Size of Connection Pool: ", len(pool.Clients))

	}
}
