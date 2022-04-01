package dao

import (
	"proxy-collect/config"
	"proxy-collect/dao/database"
	"proxy-collect/dao/redis"
)

var ProxyDao proxyDaoInterface

func LoadDao() {
	c := config.Get().Dao
	if c == "redis" {
		ProxyDao = redis.NewRedisProxyDao()
	} else {
		ProxyDao = database.NewMysqlProxyDao()
	}
}
