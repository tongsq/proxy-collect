package dao

import (
	"proxy-collect/dao/redis"
)

var ProxyDao proxyDaoInterface = redis.NewRedisProxyDao()
