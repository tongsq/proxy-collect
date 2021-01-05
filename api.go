package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tongsq/go-lib/logger"
	"proxy-collect/dao"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/all", func(c *gin.Context) {
		proxies, err := dao.ProxyDao.GetActiveList()
		if err != nil {
			logger.Error("get active proxy fail", logger.Fields{"err": err})
			c.JSON(200, gin.H{
				"data":  []string{},
				"code":  0,
				"count": 0,
			})
			return
		}
		logger.FInfo("count:%d, cap: %d\n", len(proxies), cap(proxies))
		var list []string
		for _, proxy := range proxies {
			str := proxy.Host + ":" + proxy.Port
			list = append(list, str)
		}
		c.JSON(200, gin.H{
			"data":  list,
			"code":  0,
			"count": len(list),
		})
	})
	r.Run("0.0.0.0:8090")
}
