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
		
		if isUnique := pool.UsernameUnique(username); !isUnique{
			closeMessage := websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "Username already taken") 	
			conn.WriteMessage(websocket.CloseMessage,closeMessage)
		 	conn.Close()
			return
		}

		client := Client{
			ID:   username,
			Conn: conn,
			Pool: pool,
		}

		pool.Register <- &client

		client.Read()

	}
}
