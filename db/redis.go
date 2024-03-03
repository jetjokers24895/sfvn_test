package db

import (
	"fmt"
	"sfvn_test/config"

	"github.com/redis/rueidis"
)

func InitRedisClient(cfg *config.Config) rueidis.Client {

	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{fmt.Sprintf("localhost:%v", cfg.Redis.Port)}})
	if err != nil {
		panic(err)
	}
	return client
}
