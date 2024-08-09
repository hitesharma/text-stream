package handler

import (
	"log"
	"time"

	ws "github.com/gorilla/websocket"
	"github.com/hitesharma/text-stream/internal/provider"
	"github.com/hitesharma/text-stream/pkg/infra/websocket"
)

// TextHandler implements the Handler interface
type TextHandler struct {
	providerManager *provider.ProviderManager
}

// NewTextHandler initializes a new TextHandler
func NewTextHandler(manager *provider.ProviderManager) *TextHandler {
	return &TextHandler{providerManager: manager}
}

// ServeWsConn handles WebSocket connection and streams data
func (h *TextHandler) ServeWsConn(ctx *websocket.Context) {
	defer ctx.Conn.Close()

	// Set up a close handler
	ctx.Conn.SetCloseHandler(func(code int, text string) error {
		log.Printf("Connection closed: code=%d, text=%s", code, text)
		return nil
	})

	// Create a channel to signal when the connection is closed
	closeSignal := make(chan struct{})
	go func() {
		defer close(closeSignal)
		for {
			_, _, err := ctx.Conn.ReadMessage()
			if err != nil {
				// Handle the error and close the connection
				if ws.IsCloseError(err, ws.CloseGoingAway, ws.CloseAbnormalClosure) {
					log.Printf("Client connection closed abnormally:: %v", err)
				} else if ws.IsCloseError(err, ws.CloseNormalClosure) {
					log.Printf("Client connection closed normally:: %v", err)
				} else {
					log.Printf("Error reading message: %v", err)
				}
				// Exit the goroutine and signal closure
				return
			}
		}
	}()

	// Start monitoring and switching providers
	h.providerManager.MonitorAndSwitch()

	for {
		select {
		case <-closeSignal:
			// Stop processing if the close signal is received
			log.Println("Connection closed, stopping processing")
			return

		default:
			response, err := h.providerManager.RunCurrentProvider()
			if err != nil {
				log.Printf("Error: %v", err)
				h.providerManager.SwitchProvider()
				continue
			}

			// Send message to the WebSocket client
			err = ctx.Conn.WriteMessage(ws.TextMessage, []byte(response))
			if err != nil {
				log.Println("Error writing to WebSocket:", err)
				break
			}

			// Temporary: Add time delay to mimic message processing.
			time.Sleep(1 * time.Second)
		}
	}
}
