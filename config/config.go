package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB         *DBConfig
	HTTP       *HTTPConfig
	Coingeckor *CoingeckorConfig
	Redis      *RedisConfig
}

func NewConfig() *Config {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Println("Error loading .env file at " + ".env." + env)
	}

	return &Config{
		DB:         LoadDBConfig(),
		HTTP:       LoadHTTPConfig(),
		Coingeckor: LoadCoingeckorConfig(),
		Redis:      LoadRedisConfig(),
	}
}
