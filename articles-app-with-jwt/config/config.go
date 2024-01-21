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
		AppPort: get("APP_PORT", "8000"),
	}
}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		Host:     get("DB_HOST", "localhost"),
		Port:     get("DB_PORT", "5432"),
		Username: get("DB_USERNAME", "postgres"),
		Password: get("DB_PASSWORD", "password"),
		Database: get("DB_DATABASE", "articles_app"),
		SSLMode:  get("DB_SSLMODE", "disable"),
		TimeZone: get("DB_TIMEZONE", "UTC"),
	}
}

func GetJWTConfig() *JWTConfig {
	return &JWTConfig{
		Secret: get("JWT_SECRET", "secret"),
		Expiry: get("JWT_EXPIRY", ""), // @TODO: use expiry from ENV
	}
}

func get(key string, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultVal
}
