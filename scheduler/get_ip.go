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
	proxyService := service.ProxyService{}
	service_list := []service.GetProxyInterface{&service.GetProxyXila{}, &service.GetProxyNima{}}
	for _, getProxyService := range service_list {
		go proxyService.DoGetProxy(getProxyService, pool, model.DB)
	}
}
