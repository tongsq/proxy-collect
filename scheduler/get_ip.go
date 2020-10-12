package scheduler

import (
	"proxy-collect/component"
	"proxy-collect/component/logger"
	"proxy-collect/service"
)

type GetIp struct {
}

func (s GetIp) Run() {
	logger.Success("collect ip start run")
	pool := component.NewTaskPool(100)
	defer pool.Close()
	serviceList := []service.GetProxyInterface{
		service.GetProxyXila,
		service.GetProxyNima,
		service.GetProxyKuai,
		service.GetProxyData5u,
		service.GetProxy66ip,
		service.GetProxyGuoBanjia,
		service.GetProxyCoderBusy,
		service.GetProxyIp3366,
		service.GetProxyIpJiangXianLi,
		service.GetProxy89Ip,
		service.GetProxy7Yip,
		service.GetProxyProxyList,
		service.GetProxyZdaye,
	}
	for _, getProxyService := range serviceList {
		go service.ProxyService.DoGetProxy(getProxyService, pool)
	}
}
