package websocket

import (
	g "bomberman-dom/game"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool // holds channels for communicating in websocket connection
}

type Message struct {
	Type      string       `json:"type"`
	Creator   string       `json:"creator"`
	Body      string       `json:"body"`
	GameState *g.GameState `json:"gameState"`
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

func (m *Message) ExplosionComplete(pool *Pool){
	time.Sleep(900 * time.Millisecond)
	m.Type = "EXPLOSION_COMPLETED"
	pool.Broadcast <- *m
}

func (m *Message) BombExploded(pool *Pool){
	time.Sleep(3000 * time.Millisecond)
	m.Type = "BOMB_EXPLODED"
	pool.Broadcast <- *m
}

func (m *Message) MonstersReborn(pool *Pool, gameState *g.GameState,monstersLostLives []int){
	if len(monstersLostLives) == 0 {return}
	time.Sleep(3000 * time.Millisecond)
	m.Type = "PLAYER_REBORN"
	gameState.LetMonstersReborn(monstersLostLives)
	pool.Broadcast <- *m	
}

func (m *Message) UpdateMap(pool *Pool, gameState *g.GameState, destroyedBlocks []int){
	if len(destroyedBlocks) == 0 {return}
	time.Sleep(1000 * time.Millisecond)
	m.Type = "MAP_UPDATE"
	gameState.Map = g.DestroyBlocks(gameState.Map, destroyedBlocks)
	gameState.RevealPowerUps(destroyedBlocks)

	pool.Broadcast <- *m
}
