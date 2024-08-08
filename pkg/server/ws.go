package server

import (
	"github.com/hitesharma/text-stream/internal/handler"
	"github.com/hitesharma/text-stream/internal/provider"
	"github.com/hitesharma/text-stream/pkg/infra/websocket"
)

type TextServer struct {
}

func NewTextServer() *websocket.Server {
	// Initialize providers and provider manager
	providersList := []provider.InferenceProvider{
		provider.ProviderA,
		provider.ProviderB,
		provider.ProviderC,
	}
	providerManager := provider.NewProviderManager(providersList)

	// Allocate server with custom handler
	h := handler.NewTextHandler(providerManager)
	server := websocket.AllocateServer(h)
	return server
}
