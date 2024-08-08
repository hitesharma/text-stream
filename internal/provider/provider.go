package provider

import (
	"errors"
	"math/rand"
)

// InferenceProvider is a function type for providers
type InferenceProvider func() (string, error)

func ProviderA() (string, error) {
	if rand.Float32() < 0.1 {
		return "", errors.New("Provider A failed")
	}
	return "Provider A response", nil
}

func ProviderB() (string, error) {
	if rand.Float32() < 0.2 {
		return "", errors.New("Provider B failed")
	}
	return "Provider B response", nil
}

func ProviderC() (string, error) {
	if rand.Float32() < 0.3 {
		return "", errors.New("Provider C failed")
	}
	return "Provider C response", nil
}
