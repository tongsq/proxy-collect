package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"proxy-collect/model"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/all", func(c *gin.Context) {
		var proxies []model.Proxy
		model.DB.Where("status=?", 1).Find(&proxies)
		fmt.Printf("count:%d, cap: %d\n", len(proxies), cap(proxies))
		list := []string{}
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
