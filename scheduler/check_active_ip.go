package scheduler

import (
	"github.com/tongsq/go-lib/component"
	"github.com/tongsq/go-lib/logger"
	"proxy-collect/model"
	"proxy-collect/service"
)

type CheckActiveIp struct {
}

func (s CheckActiveIp) Run() {
	logger.Success("check active ip start run")
	var proxies []model.Proxy
	model.DB.Where("status=?", 1).Find(&proxies)
	logger.Info("count:%d, cap: %d\n", len(proxies), cap(proxies))
	//pool := component.NewTaskPool(20)
	//defer pool.Close()
	for _, proxy := range proxies {
		var proxyTmp model.Proxy = proxy
		component.TaskPool.RunTask(func() { service.ProxyService.CheckProxyAndSave(proxyTmp.Host, proxyTmp.Port, "") })
	}
}
