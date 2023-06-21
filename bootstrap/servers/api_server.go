package servers

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
	router := r.Use(TokenAuthMiddleware)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/all", func(c *gin.Context) {
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
		var durationMin, durationMax int64
		duration := c.Query("duration")
		if duration != "" {
			if strings.Contains(duration, "-") {
				strs := strings.Split(duration, "-")
				if len(strs) >= 2 {
					durationMin, _ = strconv.ParseInt(strs[0], 10, 64)
					durationMax, _ = strconv.ParseInt(strs[1], 10, 64)
				}
			} else {
				durationMin, _ = strconv.ParseInt(duration, 10, 64)
			}
		}
		proto := c.Query("proto")
		count := c.Query("count")
		countNum, _ := strconv.Atoi(count)
		var list []dto.ProxyInfoDto
		nowTime := time.Now().Unix()
		for _, proxy := range proxies {
			if city != "" && !strings.Contains(proxy.City, city) {
				continue
			}
			if durationMin > 0 && ((nowTime - proxy.ActiveTime) < durationMin) {
				continue
			}
			if durationMax > 0 && ((nowTime - proxy.ActiveTime) > durationMax) {
				continue
			}
			if proto != "" && proto != proxy.Proto {
				continue
			}
			list = append(list, dto.NewProxyDto(proxy))
			if countNum > 0 && len(list) >= countNum {
				break
			}
		}
		c.JSON(200, gin.H{
			"data":  list,
			"code":  0,
			"count": len(list),
		})
	})
	err := r.Run(fmt.Sprintf("%s:%s", config.Get().Api.Host, config.Get().Api.Port))
	if err != nil {
		panic("start api server fail:" + err.Error())
	}
}

func TokenAuthMiddleware(c *gin.Context) {
	tokenSecret := config.Get().Api.Token
	if tokenSecret != "" {
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}
		if token != tokenSecret {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "API token required",
			})
			c.Abort()
			return
		}
	}
	c.Next()
}
