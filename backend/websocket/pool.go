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
	BombExploded chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:     make(chan *Client),
		Unregister:   make(chan *Client),
		Clients:      make(map[*Client]bool),
		Broadcast:    make(chan Message),
		BombExploded: make(chan Message),
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
			currentPlayerIndex := gameState.FindPlayer(message.Creator)
			currentPlayer := gameState.Players[currentPlayerIndex]
			if message.Type == "PLAYER_MOVE" {
				// fmt.Println(message.Creator)
				gameState.Players[currentPlayerIndex].Move(message.Body)
			} else if message.Type == "PLAYER_DROPPED_BOMB" {
				if gameState.Players[currentPlayerIndex].BombsLeft <= 0 {
					break S
				}

				gameState.Players[currentPlayerIndex].Bombs = append(gameState.Players[currentPlayerIndex].Bombs, game.Bomb{X: currentPlayer.X, Y: currentPlayer.Y})
				gameState.Players[currentPlayerIndex].BombsLeft--

				go startBombTimer(pool.BombExploded, message)

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
		case message := <-pool.BombExploded:
			currentPlayerIndex := gameState.FindPlayer(message.Creator)
			gameState.Players[currentPlayerIndex].BombsLeft++
			gameState.Players[currentPlayerIndex].Bombs = gameState.Players[currentPlayerIndex].Bombs[1:]

		}
		// fmt.Println("Size of Connection Pool: ", len(pool.Clients))

	}
}

func startBombTimer(ch chan Message, msg Message) {
	time.Sleep(1000 * time.Millisecond)
	ch <- msg
}
