package config

import "context"

type awsSecretManager struct {
	config int
}

func newAWSSecretManager(ctx context.Context) (*awsSecretManager, error) {
	return &awsSecretManager{
		config: 0,
	}, nil
}

func (s *awsSecretManager) getSecretValue(ctx context.Context, secretName string) (map[string]string, error) {
	return map[string]string{
		"provider-a-api-key": "provider-a-test-api-key",
		"provider-b-api-key": "provider-b-test-api-key",
		"another-key":        "another-key",
	}, nil
}
