package main

import (
	ws "bomberman-dom/websocket"
	"fmt"
	"log"
	"net/http"
)

func main() {
	wsClients := ws.Clients{Connections: make(map[*ws.Client]bool)}

	http.HandleFunc("/ws", ws.SocketHandler(&wsClients))

	fmt.Printf("Server started at http://localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
