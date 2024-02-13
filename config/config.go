// File: config/config.go

package config

import (
    "os"
)

// Config holds all configuration for the application.
// Add other configuration fields as needed.
type Config struct {
    KafkaBroker  string
    JaegerURL    string
    ServerPort   string
}

// LoadConfig loads the application configuration from environment variables
// or fallbacks to default values.
func LoadConfig() (*Config, error) {
    return &Config{
        KafkaBroker:  getEnv("KAFKA_BROKER", "localhost:9092"),
        JaegerURL:    getEnv("JAEGER_URL", "http://localhost:14268/api/traces"),
        ServerPort:   getEnv("SERVER_PORT", "3000"),
    }, nil
}

// getEnv reads an environment variable or returns a default value if not set.
func getEnv(key, defaultValue string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        value = defaultValue
    }
    return value
}
