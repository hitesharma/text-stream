package provider

import (
	"log"
)

type ProviderManager struct {
	providers []InferenceProvider
	currentId int
}

func NewProviderManager(providers []InferenceProvider) *ProviderManager {
	return &ProviderManager{
		providers: providers,
		currentId: 0,
	}
}

// GetCurrentProvider gets current provider based on the set current id
func (m *ProviderManager) GetCurrentProvider() InferenceProvider {
	return m.providers[m.currentId]
}

// SwitchProvider switches between providers in case of failure
func (m *ProviderManager) SwitchProvider() {
	m.currentId = (m.currentId + 1) % len(m.providers)
	log.Printf("Switched to provider %d", m.currentId+1)
}
