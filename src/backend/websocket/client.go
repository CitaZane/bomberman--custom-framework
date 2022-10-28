package websocket

import "github.com/gorilla/websocket"

type Client struct {
	ID   string
	conn *websocket.Conn //ws connection
	send chan []byte     //sedn channel for outgoing messages
}

type Clients struct {
	Connections map[*Client]bool
}

func (c *Clients) UnregisterClient(client *Client) {
	if _, ok := c.Connections[client]; ok {
		delete(c.Connections, client)
	}
}

func (c *Clients) RegisterNewClient(client *Client) {
	c.Connections[client] = true
}

func NewClient(conn *websocket.Conn, ID string) *Client {
	return &Client{
		ID:   ID,
		conn: conn,
		send: make(chan []byte, 256),
	}
}
