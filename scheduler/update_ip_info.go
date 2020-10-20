package scheduler

import (
	"fmt"
	"proxy-collect/component/logger"
	"proxy-collect/model"
	"proxy-collect/service"
	"time"
)

type UpdateIpInfo struct {
}

func (s UpdateIpInfo) Run() {
	logger.Success("update ip info start run")
	var proxies []model.Proxy
	model.DB.Where("status=? and country=?", 1, "").Find(&proxies)
	logger.Info(fmt.Sprintf("count:%d, cap: %d\n", len(proxies), cap(proxies)))

	for _, proxy := range proxies {
		ipInfoDto := service.ProxyService.GetIpInfo(proxy.Host, proxy.Port)
		logger.Info("get ip info:", ipInfoDto)
		time.Sleep(3 * time.Second)
		if ipInfoDto == nil {
			logger.Error("get ip info fail")
			continue
		}
		proxy.Country = ipInfoDto.Country
		proxy.Isp = ipInfoDto.Isp
		proxy.Region = ipInfoDto.Region
		proxy.City = ipInfoDto.City
		model.DB.Save(&proxy)
		logger.Success("update ip detail success:" + proxy.Host + ":" + proxy.Port)
	}
}
