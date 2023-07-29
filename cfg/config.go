package cfg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Config struct {
	Postgres DBConfig
}

var configs *Config

func GetAppPort() string {
	return getEnvOrDefault("APP_PORT", "8081")
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func GetConfig() DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println(err, "Couldn't connect to db")
	}
	configs := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	return configs
}
