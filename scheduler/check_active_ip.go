package scheduler

import (
	"fmt"
	"proxy-collect/component"
	"proxy-collect/component/logger"
	"proxy-collect/model"
	"proxy-collect/service"
)

type CheckActiveIp struct {
}

func (s CheckActiveIp) Run() {
	logger.Success("check active ip start run")
	var proxies []model.Proxy
	model.DB.Where("status=?", 1).Find(&proxies)
	fmt.Printf("count:%d, cap: %d\n", len(proxies), cap(proxies))
	pool := component.NewTaskPool(20)
	for _, proxy := range proxies {
		var proxyTmp model.Proxy = proxy
		pool.RunTask(func() { service.ProxyService.CheckProxyAndSave(proxyTmp.Host, proxyTmp.Port, "") })
	}
}
