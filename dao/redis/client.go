package redis

import redis_client "github.com/tongsq/go-lib/redis-client"

var Client *redis_client.RedisClient = &redis_client.RedisClient{
	MaxIdle:   10,
	MaxActive: 20,
	Network:   "tcp",
	Address:   "127.0.0.1:6379",
}
