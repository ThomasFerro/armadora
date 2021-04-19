package config

import (
	"log"
	"os"
)

type ConfigurationKey string

var ConfigurationsWithFallback = map[ConfigurationKey]string{
	"ALLOWED_ORIGIN":                   "http://localhost:8081",
	"PORT":                             "80",
	"MONGO_URI":                        "mongodb://localhost:27017",
	"MONGO_DATABASE_NAME":              "armadora",
	"MONGO_EVENT_COLLECTION_NAME":      "events",
	"MONGO_PARTY_COLLECTION_NAME":      "parties",
	"MONGO_PROJECTION_COLLECTION_NAME": "projections",
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
