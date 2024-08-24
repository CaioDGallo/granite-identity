package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress    string
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	Env              string
	MessageBrokerURL string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		ServerAddress:    getEnv("SERVER_ADDRESS", ":8080"),
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "5432"),
		DBUser:           getEnv("DB_USER", "postgres"),
		DBPassword:       getEnv("DB_PASSWORD", "password"),
		DBName:           getEnv("DB_NAME", "accounts_db"),
		Env:              getEnv("ENV", "development"),
		MessageBrokerURL: getEnv("MESSAGE_BROKER_URL", "amqp://guest:guest@172.17.0.1:5672/"),
	}

	return config, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		intValue, err := strconv.Atoi(value)
		if err == nil {
			return intValue
		}
		log.Printf("Invalid value for %s, using default: %d", key, fallback)
	}
	return fallback
}

func getEnvAsDuration(key string, fallback time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		durationValue, err := time.ParseDuration(value)
		if err == nil {
			return durationValue
		}
		log.Printf("Invalid duration format for %s, using default: %s", key, fallback.String())
	}
	return fallback
}
