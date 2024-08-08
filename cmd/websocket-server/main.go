package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hitesharma/text-stream/pkg/server"
)

const (
	WS_PORT = 8088
)

func main() {
	srv := server.NewTextServer()
	// Start WebSocket server
	http.Handle("/ws", srv)
	log.Printf("WebSocket server started on :%d", WS_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", WS_PORT), nil))
}
