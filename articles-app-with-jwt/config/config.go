package config

import (
	"os"
)

type AppConfig struct {
	AppPort string
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
	TimeZone string
}

type JWTConfig struct {
	Secret string
	Expiry string
}

func GetAppConfig() *AppConfig {
	return &AppConfig{
		AppPort: Get("APP_PORT", "8000"),
	}
}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		Host:     Get("DB_HOST", "localhost"),
		Port:     Get("DB_PORT", "5432"),
		Username: Get("DB_USERNAME", "postgres"),
		Password: Get("DB_PASSWORD", "password"),
		Database: Get("DB_DATABASE", "articles_app"),
		SSLMode:  Get("DB_SSLMODE", "disable"),
		TimeZone: Get("DB_TIMEZONE", "UTC"),
	}
}

func GetJWTConfig() *JWTConfig {
	return &JWTConfig{
		Secret: Get("JWT_SECRET", "secret"),
		Expiry: Get("JWT_EXPIRY", ""), // @TODO: use expiry from JWT
	}
}

func Get(key string, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultVal
}
