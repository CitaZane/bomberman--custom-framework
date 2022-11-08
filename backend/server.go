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

	go pool.Start(&game.State)

	http.HandleFunc("/ws", ws.SocketHandler(pool))
	fmt.Printf("Server started at http://localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
