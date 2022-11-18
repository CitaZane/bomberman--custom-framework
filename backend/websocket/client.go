package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool // holds channels for communicating in websocket connection
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
		var msg Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			return
		}
		// send created message to broadcast channel
		c.Pool.Broadcast <- msg
		// fmt.Printf("Text message Received: %+v\n", msg)
	}
}
