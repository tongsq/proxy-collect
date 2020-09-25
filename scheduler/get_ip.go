package scheduler

import (
	"proxy-collect/component"
	"proxy-collect/model"
	"proxy-collect/service"
)

type GetIp struct {
}

func (s GetIp) Run() {
	pool := component.NewTaskPool(20)
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
	}
	for _, getProxyService := range serviceList {
		go service.ProxyService.DoGetProxy(getProxyService, pool, model.DB)
	}
}
