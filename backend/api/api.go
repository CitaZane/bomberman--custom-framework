package api

import (
	"bomberman-dom/websocket"
	"fmt"
	"net/http"
	"strconv"
)

func QueuePlayerCount(pool *websocket.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:1234")
		fmt.Println("ClientsCount", len(pool.Clients))
		w.Write([]byte(strconv.Itoa(len(pool.Clients))))
	}
}
