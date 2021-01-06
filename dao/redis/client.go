package redis

import (
	redis_client "github.com/tongsq/go-lib/redis-client"
	"proxy-collect/config"
)

var Client *redis_client.RedisClient = &redis_client.RedisClient{
	MaxIdle:   config.Get().Redis.MaxIdle,
	MaxActive: config.Get().Redis.MaxActive,
	Network:   config.Get().Redis.Network,
	Address:   config.Get().Redis.Address,
	Password:  config.Get().Redis.Password,
}
