package configs

import (
	"os"
)

type Config struct {
	DatabaseURL string
	JWTSecret   string
	Environment string
	Port        string      
}

func Load() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "debt_tracker.db"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key-change-this-in-production"),
		Environment: getEnv("ENVIRONMENT", "development"),
		Port:        getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}