package scheduler

import (
	"github.com/tongsq/go-lib/component"
	"github.com/tongsq/go-lib/logger"
	"proxy-collect/dao"
	"proxy-collect/model"
	"proxy-collect/service"
)

type CheckFailIp struct {
}

func (s CheckFailIp) Run() {
	logger.FSuccess("check fail ip start run")
	var proxies []model.ProxyModel = dao.ProxyDao.GetFailList()
	logger.FInfo("check fail ip, len:%d, cap:%d", len(proxies), cap(proxies))
	for _, proxy := range proxies {
		var proxyTmp model.ProxyModel = proxy
		component.TaskPool.RunTask(func() { service.ProxyService.CheckProxyAndSave(proxyTmp.Host, proxyTmp.Port, "") })
	}
}
