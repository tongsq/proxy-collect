package scheduler

import (
	"github.com/tongsq/go-lib/logger"
	"proxy-collect/dao"
	"proxy-collect/global"
	"proxy-collect/model"
	"proxy-collect/service"
)

type CheckFailIp struct {
}

func (s CheckFailIp) Run() {
	logger.FSuccess("check fail ip start run")
	proxies, err := dao.ProxyDao.GetFailList()
	if err != nil {
		logger.Error("get need check ip list fail", logger.Fields{"err": err})
	}
	logger.FInfo("check fail ip, len:%d, cap:%d", len(proxies), cap(proxies))
	for _, proxy := range proxies {
		var proxyTmp model.ProxyModel = proxy
		global.Pool.RunTask(func() { service.ProxyService.CheckProxyAndSave(proxyTmp.Host, proxyTmp.Port, "") })
	}
}
