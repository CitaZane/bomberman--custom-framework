package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// holds game state to send it to all players
type Game struct {
	Players []string `json:"players"`
	created bool
}

var game Game

func SocketHandler(pool *Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("Error during connection upgradation:", err)
			return
		}

		username := r.URL.Query()["username"][0]

		if !game.created {
			game.created = true
		}
		game.Players = append(game.Players, username)

		client := Client{
			ID:   username,
			Conn: conn,
			Pool: pool,
		}

		pool.Register <- &client

		client.Read()

	}
}
