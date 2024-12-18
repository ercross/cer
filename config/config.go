package config

import (
	"context"
	"fmt"
	"os"
)

type Config struct {
	ApiPort         string
	ProviderAApiKey string
	ProviderBApiKey string
}

const defaultApiPort = "15001"

func LoadConfig(awsCredential string) (*Config, error) {
	secrets, err := fetchSecret("cadana-aws-secret")
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	cfg := &Config{
		ProviderAApiKey: secrets["provider_a_api_key"],
		ProviderBApiKey: secrets["provider_b_api_key"],
	}

	apiPort := os.Getenv("API_PORT")
	if len(apiPort) == 0 {
		apiPort = defaultApiPort
	}
	cfg.ApiPort = apiPort
	return cfg, nil
}

// FetchSecret retrieves a secret from AWS Secrets Manager
func fetchSecret(secretName string) (map[string]string, error) {
	ctx := context.Background()

	asm, err := newAWSSecretManager(ctx)
	if err != nil {
		return nil, fmt.Errorf("error instantiating a new instance of aws secret manager: %w", err)
	}
	return asm.getSecretValue(ctx, secretName)
}
