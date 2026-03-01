package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func Init() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Error read .env file")
		return nil, err
	}
	port, err := strconv.Atoi(getEnv("DB_PORT", "8080"))
	log.Println(port)
	if err != nil {
		log.Println("Cannot get port")
		return nil, err
	}

	return &Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     port,
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", ""),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
