package scheduler

import (
	"github.com/tongsq/go-lib/logger"
	"proxy-collect/global"
	"proxy-collect/service"
)

type GetIp struct {
}

func (s GetIp) Run() {
	logger.FSuccess("collect ip start run")
	//pool := component.NewTaskPool(100)
	//defer pool.Close()
	serviceList := []service.GetProxyInterface{
		service.GetProxyXila,
		service.GetProxyNima,
		service.GetProxyKuai,
		service.GetProxyData5u,
		service.GetProxy66ip,
		//service.GetProxyGuoBanjia,
		service.GetProxyCoderBusy,
		service.GetProxyIp3366,
		service.GetProxyIpJiangXianLi,
		service.GetProxy89Ip,
		service.GetProxy7Yip,
		service.GetProxyProxyList,
		service.GetProxyZdaye,
		service.GetProxyFanQie,
		service.GetProxyZdayeIndex,
		service.GetProxySeofangfa,
		service.GetProxyXsdaili,
	}
	for _, getProxyService := range serviceList {
		go service.ProxyService.DoGetProxy(getProxyService, global.Pool)
	}

}
