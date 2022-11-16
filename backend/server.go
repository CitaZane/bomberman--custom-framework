package main

import (
	"bomberman-dom/game"
	ws "bomberman-dom/websocket"
	"fmt"
	"log"
	"net/http"
)

func main() {
	pool := ws.NewPool()
	game := game.GameState{
		Bombs:    make([]game.Bomb, 0),
		PowerUps: make([]*game.PowerUp, 0),
	}

	go pool.Start(&game)

	http.HandleFunc("/ws", ws.SocketHandler(pool))
	fmt.Printf("Server started at http://localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
