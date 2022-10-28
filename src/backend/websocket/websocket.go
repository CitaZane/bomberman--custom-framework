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

func SocketHandler(clients *Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("Error during connection upgradation:", err)
			return
		}

		client := NewClient(conn, "ye")
		clients.RegisterNewClient(client)

		go client.Writer()
		go client.Reader(clients)
	}
}

/* -------------------------------------------------------------------------- */
/*                    basic reader and writer for websocket conn              */
/* -------------------------------------------------------------------------- */
// define a writer which will send
// new messages to our WebSocket endpoint
func (client *Client) Writer() {
	for {
		message, ok := <-client.send
		if !ok {
			log.Println("err on writing message")
			return
		}
		w, err := client.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			return
		}
		_, err = w.Write(message)
		if err != nil {
			log.Println("Line 91", err)
			return
		}
		if err := w.Close(); err != nil {
			log.Println("Line 95", err)
			return
		}
	}
}

// define a reader which will listen for
// new messages being sent to our WebSocketendpoint
// Unregister client when client disconnect
func (client *Client) Reader(clients *Clients) {
	defer clients.UnregisterClient(client)
	for {
		// read in a message
		_, _, err := client.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// log.Println(msg)
	}
}
