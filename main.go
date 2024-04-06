package main

import (
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {

	// Create a chat room instance
	Cr := newRoom("Test")

	http.Handle("/", http.FileServer(http.Dir("webpage")))

	// Create websocket connection
	http.Handle("/ws", websocket.Handler(Cr.handleWS))

	http.ListenAndServe(":3000", nil)
}
