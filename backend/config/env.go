package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DBUrl     string
	JWTSecret string
}

var Env *Config

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, using system environment variables")
	}

	Env = &Config{
		Port:      getEnv("PORT", "8080"),
		DBUrl:     getEnv("DB_URL", ""),
		JWTSecret: getEnv("JWT_SECRET", ""),
	}

	// Validate required environment variables
	if Env.DBUrl == "" {
		log.Fatal("DB_URL is not set in environment variables")
	}
	if Env.JWTSecret == "" {
		log.Fatal("JWT_SECRET is not set in environment variables")
	}

	log.Println("Environment variables loaded successfully")
}

// getEnv retrieves environment variable or returns default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
