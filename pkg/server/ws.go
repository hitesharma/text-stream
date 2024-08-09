package server

import (
	"github.com/hitesharma/text-stream/internal/handler"
	"github.com/hitesharma/text-stream/internal/provider"
	"github.com/hitesharma/text-stream/pkg/infra/websocket"
)

func NewTextServer() *websocket.Server {
	providerManager := provider.GetProviderManager()
	h := handler.NewTextHandler(providerManager)
	server := websocket.AllocateServer(h)
	return server
}
