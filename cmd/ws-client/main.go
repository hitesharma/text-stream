package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

const (
	SERVER_URL = "ws://localhost:8088/ws"
)

func main() {
	// Connect to the WebSocket server
	log.Printf("Connecting to %s...", SERVER_URL)
	conn, _, err := websocket.DefaultDialer.Dial(SERVER_URL, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer conn.Close()
	log.Println("Connected to server.")

	// Channel to listen for system interrupts | SIGINT
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	done := make(chan struct{})

	// Listen for messages from the server
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}
			log.Printf("Received: %s", message)
		}
	}()

	// Handle interrupts and cleanup
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("Received interrupt signal. Closing connection...")

			// Gracefully close the WebSocket connection by sending a close message
			// and then waiting for the server to close the connection
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error during closing handshake:", err)
				return
			}

			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
