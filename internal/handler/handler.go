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

	// Start monitoring and switching providers
	h.providerManager.MonitorAndSwitch()

	for {
		response, err := h.providerManager.GetCurrentProvider()()
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

		time.Sleep(1 * time.Second)
	}
}
