package config

import "os"

type CoingeckorConfig struct {
	Domain string
	ApiKey string
}

func LoadCoingeckorConfig() *CoingeckorConfig {
	return &CoingeckorConfig{
		Domain: os.Getenv("COINGECKOR_DOMAIN"),
		ApiKey: os.Getenv("COINGECKOR_APIKEY"),
	}
}
