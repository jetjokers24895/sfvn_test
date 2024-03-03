package config

import "os"

type RedisConfig struct {
	Port string
}

func LoadRedisConfig() *RedisConfig {
	return &RedisConfig{
		Port: os.Getenv("REDIS_PORT"),
	}
}
