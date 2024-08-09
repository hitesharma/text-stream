package provider

import (
	"fmt"
	"math/rand"
)

// InferenceProvider is a function type for providers
type InferenceProvider func(string) (string, error)

// inferenceProviderHosts is an array that serves as a centralized collection
// of inference server hostnames. To add a new inference server, simply add
// its hostname to this array, unless additional configuration is required.
var inferenceProviderHosts = []string{
	"Provider 1",
	"Provider 2",
	"Provider 3",
}

// ProviderStub simulates an inference provider that processes requests
// based on a given host name. It randomly fails with a 10% chance,
// returning an error response. Otherwise, it returns a fixed response
// message including the host name. This stub can be expanded to take
// user input and generate more dynamic responses if needed.
func ProviderStub(hostName string) (string, error) {
	if rand.Float32() < 0.1 {
		return "", fmt.Errorf("%s failed", hostName)
	}
	return fmt.Sprintf("%s response", hostName), nil
}

// The init function initializes the providerManager by populating its
// providers slice with InferenceProvider functions. It iterates over
// the inferenceProviderHosts array, creating a closure for each hostname.
// This closure captures the current hostname and returns an InferenceProvider
// function that will use this hostname when calling ProviderStub.
// By using a closure, each provider in the providerManager.providers slice
// is correctly associated with its respective hostname.
func init() {
	for _, v := range inferenceProviderHosts {
		// Instead of calling ProviderStub(v), we create a closure
		// that captures v and returns an InferenceProvider
		providerManager.providers = append(providerManager.providers, func(host string) InferenceProvider {
			return func(input string) (string, error) {
				return ProviderStub(host)
			}
		}(v))
	}
}
