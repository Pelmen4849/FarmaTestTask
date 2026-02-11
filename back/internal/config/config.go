package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	ServerPort string
	Timeout    int

	JWTSecret string
}

func Load() *Config {
	// Загружаем .env файл, если он существует
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment")
	}

	cfg := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "farm_site"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),

		ServerPort: getEnv("SERVER_PORT", "8080"),
		Timeout:    getEnvAsInt("SERVER_TIMEOUT", 30),

		JWTSecret: getEnv("JWT_SECRET", "default_secret_change_me"),
	}

	// Проверяем обязательные поля
	if cfg.DBPassword == "" {
		log.Fatal("DB_PASSWORD is required")
	}
	if cfg.JWTSecret == "default_secret_change_me" {
		log.Fatal("JWT_SECRET must be set in production")
	}

	return cfg
}

// Helper: получить строковое значение переменной или значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Helper: получить целочисленное значение переменной
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
