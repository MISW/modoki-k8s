package config

import (
	"os"
	"strings"
)

type EnvConfig struct {
	APIKeys []string
	Domain  string
}

func ReadEnv() *EnvConfig {
	apiKeys := strings.Split(os.Getenv("MODOKI_API_KEY"), ",")
	domain := os.Getenv("MODOKI_APP_DOMAIN")

	return &EnvConfig{
		APIKeys: apiKeys,
		Domain:  domain,
	}
}
