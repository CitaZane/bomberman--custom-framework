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

func SocketHandler(pool *Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("Error during connection upgradation:", err)
			return
		}

		username := r.URL.Query()["username"][0]

		client := Client{
			ID:   username,
			Conn: conn,
			Pool: pool,
		}

		pool.Register <- &client

		client.Read()

	}
}
