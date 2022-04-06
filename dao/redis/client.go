package redis

import (
	redis_client "github.com/tongsq/go-lib/redis-client"
	"proxy-collect/config"
)

var client *redis_client.RedisClient

func Client() *redis_client.RedisClient {
	if client == nil {
		client = &redis_client.RedisClient{
			MaxIdle:   config.Get().Redis.MaxIdle,
			MaxActive: config.Get().Redis.MaxActive,
			Network:   config.Get().Redis.Network,
			Address:   config.Get().Redis.Address,
			Password:  config.Get().Redis.Password,
		}
	}
	return client
}
