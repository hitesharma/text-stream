package provider

import (
	"log"
	"time"
)

type ProviderManager struct {
	providers []InferenceProvider
	currentId int
}

var providerManager = &ProviderManager{}

func GetProviderManager() *ProviderManager {
	providerManager.currentId = 0
	return providerManager
}

// RunCurrentProvider executes the current provider function based on the
// set current ID in the ProviderManager.
func (m *ProviderManager) RunCurrentProvider() (string, error) {
	provider := m.providers[m.currentId]
	currentHost := inferenceProviderHosts[m.currentId]

	return func(currentHost string) (string, error) {
		return provider(currentHost)
	}(currentHost)
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
			// Run current provider to check for it's health
			_, err := m.RunCurrentProvider()
			if err != nil {
				log.Printf("Error detected: %v, switching provider", err)
				m.SwitchProvider()
			}
		}
	}()
}
