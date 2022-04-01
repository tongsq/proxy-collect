package bootstrap

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tongsq/go-lib/logger"
	"proxy-collect/config"
	"proxy-collect/dao"
	"proxy-collect/dto"
)

func StartApiServer() {
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
		city := c.Query("city")
		var durationI int64
		duration := c.Query("duration")
		if duration != "" {
			durationI, err = strconv.ParseInt(duration, 10, 64)
		}
		var list []dto.ProxyDto
		nowTime := time.Now().Unix()
		for _, proxy := range proxies {
			if city != "" && !strings.Contains(proxy.City, city) {
				continue
			}
			if durationI > 0 && ((nowTime - proxy.ActiveTime) < durationI) {
				continue
			}
			list = append(list, dto.NewProxyDto(proxy))
		}
		c.JSON(200, gin.H{
			"data":  list,
			"code":  0,
			"count": len(list),
		})
	})
	r.Run(fmt.Sprintf("%s:%s", config.Get().Api.Host, config.Get().Api.Port))
}
