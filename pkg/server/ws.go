package server

import (
	"github.com/hitesharma/text-stream/internal/handler"
	"github.com/hitesharma/text-stream/internal/provider"
	"github.com/hitesharma/text-stream/pkg/infra/websocket"
)

func NewTextServer() *websocket.Server {
	// Initialize providers and provider manager
	providersList := []provider.InferenceProvider{
		provider.Provider1,
		provider.Provider2,
		provider.Provider3,
	}
	providerManager := provider.NewProviderManager(providersList)

	// Allocate server with custom handler
	h := handler.NewTextHandler(providerManager)
	server := websocket.AllocateServer(h)
	return server
}
