package config

import (
	"os"
	"log"
)

type ConfigurationKey string

var ConfigurationsWithFallback = map[ConfigurationKey]string{
	"ALLOWED_ORIGIN": "http://localhost:8081",
	"PORT": "80",
	"EVENT_STORE_URL": "http://eventstore:2113",
	"EVENT_STORE_USERNAME": "admin",
	"EVENT_STORE_PASSWORD": "changeit",
}

func GetConfiguration(configurationKey ConfigurationKey) string {
	fallback := ConfigurationsWithFallback[configurationKey]
	configurationFromEnv := os.Getenv(string(configurationKey))
	if configurationFromEnv == "" {
		configurationFromEnv = fallback
		log.Printf(
			"No configuration found in env variables for %v, falling back to %v",
			configurationKey,
			fallback,
		)
	}
	return configurationFromEnv
}