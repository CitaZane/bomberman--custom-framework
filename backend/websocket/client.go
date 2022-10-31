package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool // holds channels for communicating in websocket connection
}

type Message struct {
	Type string `json:"type"`
	Body string `json:"body"`
}

// keep listening for messages from websocket
func (c *Client) Read() {
	defer func() {
		// unregister client by sending the client to the unregister channel
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {

		// if we get a message, we will read it here
		err := c.Conn.ReadJSON()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: "NEW_USER", Body: string(p)}
		fmt.Println("Got a message")
		// send created message to broadcast channel
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
