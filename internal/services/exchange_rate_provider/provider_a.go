package exchange

import (
	"fmt"
	"math/rand/v2"
)

type ProviderA struct {
	ApiKey string
}

func NewProviderA(apiKey string) *ProviderA {
	return &ProviderA{
		ApiKey: apiKey,
	}
}

func (a *ProviderA) FetchExchangeRate(pair string) (float64, error) {
	// make rest call to ProviderA endpoint
	sleep()

	// Introduce a random chance of failure
	if rand.Float64() < 0.2 {
		return 0, fmt.Errorf("failed to fetch exchange rate: random error occurred")
	}

	randomFloat := rand.Float64()

	return randomFloat, nil
}

func (a *ProviderA) Name() string {
	return "ProviderA"
}
