package config

import (
	"os"
	"strings"
)

type EnvConfig struct {
	APIKeys []string
}

func ReadEnv() *EnvConfig {
	apiKeys := strings.Split(os.Getenv("MODOKI_API_KEY"), ",")

	return &EnvConfig{
		APIKeys: apiKeys,
	}
}
