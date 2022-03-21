package scheduler

import (
	"fmt"
	"time"

	"github.com/tongsq/go-lib/logger"
	"proxy-collect/dao"
	"proxy-collect/service/ip"
)

type UpdateIpInfo struct {
}

func (s UpdateIpInfo) Run() {
	logger.FSuccess("update ip info start run")
	proxies := dao.ProxyDao.GetNeedUpdateInfoList()
	logger.FInfo(fmt.Sprintf("count:%d, cap: %d\n", len(proxies), cap(proxies)))
	for _, proxy := range proxies {
		//use online ip database
		ipInfoDto := ip.GetIpInfoByIp138(proxy.Host, proxy.Port)
		logger.Info("get ip info:", logger.Fields{"ipInfoDto": ipInfoDto})
		time.Sleep(5 * time.Second)
		if ipInfoDto == nil {
			logger.FError("get ip info fail")
			continue
		}
		proxy.Country = ipInfoDto.Country
		proxy.Isp = ipInfoDto.Isp
		proxy.Region = ipInfoDto.Region
		proxy.City = ipInfoDto.City
		err := dao.ProxyDao.Save(&proxy)
		if err == nil {
			logger.Success("update ip detail success", logger.Fields{"host": proxy.Host, "port": proxy.Port})
		}
	}
}
