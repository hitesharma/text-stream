package provider

import (
	"errors"
	"math/rand"
)

// InferenceProvider is a function type for providers
type InferenceProvider func() (string, error)

func Provider1() (string, error) {
	if rand.Float32() < 0.1 {
		return "", errors.New("Provider 1 failed")
	}
	return "Provider 1 response", nil
}

func Provider2() (string, error) {
	if rand.Float32() < 0.2 {
		return "", errors.New("Provider 2 failed")
	}
	return "Provider 2 response", nil
}

func Provider3() (string, error) {
	if rand.Float32() < 0.3 {
		return "", errors.New("Provider 3 failed")
	}
	return "Provider 3 response", nil
}
