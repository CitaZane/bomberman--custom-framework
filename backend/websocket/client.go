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
	Type int    `json:"type"`
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
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}

		// send created message to broadcast channel
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}