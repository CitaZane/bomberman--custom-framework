package api

import (
	"net/http"
	"strconv"
)

func QueuePlayerCount(clientsCount int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:1234")
		w.Write([]byte(strconv.Itoa(clientsCount)))
	}
}
