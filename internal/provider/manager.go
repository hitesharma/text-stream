package provider

import (
	"log"
	"time"
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

// GetCurrentProvider gets current provider based on the set current Id
func (m *ProviderManager) GetCurrentProvider() InferenceProvider {
	return m.providers[m.currentId]
}

// SwitchProvider switches to the next available provider in the pool
// in case of a failure. This implementation randomly selects a provider
// from the pool, excluding the current one.
// TODO: Consider implementing a mechanism that tracks providers with recent
// errors, allowing the system to deprioritize them and select them last.
// Update the logic of SwitchProvider accordingly.
func (m *ProviderManager) SwitchProvider() {
	m.currentId = (m.currentId + 1) % len(m.providers)
	log.Printf("Switched to provider %d", m.currentId+1)
}

// MonitorAndSwitch monitors the current provider and switches to a different
// provider if an error is encountered during the inference process.
// This implementation could be optimized by incorporating a dedicated health check
// mechanism from the provider, which would reduce the need for performing full
// inference tasks just to assess the provider's health.
func (m *ProviderManager) MonitorAndSwitch() {
	// Check every 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for range ticker.C {
			// Get current provider and check for it's health
			_, err := m.GetCurrentProvider()()
			if err != nil {
				log.Printf("Error detected: %v, switching provider", err)
				m.SwitchProvider()
			}
		}
	}()
}
