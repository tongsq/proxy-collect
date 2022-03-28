package scheduler

import (
	"github.com/tongsq/go-lib/logger"
	"proxy-collect/dao"
	"proxy-collect/global"
	"proxy-collect/model"
	"proxy-collect/service"
)

type RecheckIp struct {
}

func (s RecheckIp) Run() {
	logger.FSuccess("recheck ip start run")
	proxies, err := dao.ProxyDao.GetRecheckList()
	if err != nil {
		logger.Error("get recheck ip list fail", logger.Fields{"err": err})
	}
	logger.FInfo("recheck ip, len:%d, cap:%d", len(proxies), cap(proxies))
	for _, proxy := range proxies {
		var proxyTmp model.ProxyModel = proxy
		global.Pool.RunTask(func() { service.ProxyService.CheckProxyAndSave(proxyTmp.Host, proxyTmp.Port, "") })
	}
}
