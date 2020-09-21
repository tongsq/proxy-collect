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
		&service.GetProxyXila{},
		&service.GetProxyNima{},
		&service.GetProxyKuai{},
	}
	for _, getProxyService := range serviceList {
		go service.ProxyService.DoGetProxy(getProxyService, pool, model.DB)
	}
}
