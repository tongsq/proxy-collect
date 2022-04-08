package scheduler

import (
	"github.com/tongsq/go-lib/logger"
	"proxy-collect/dao"
	"proxy-collect/dto"
	"proxy-collect/global"
	"proxy-collect/service"
)

type CheckActiveIp struct {
}

func (s CheckActiveIp) Run() {
	logger.FSuccess("check active ip start run")
	proxies, err := dao.ProxyDao.GetActiveList()
	if err != nil {
		logger.Error("get active ip fail", logger.Fields{"err": err})
	}
	logger.FInfo("check active ip, len:%d, cap:%d", len(proxies), cap(proxies))
	//pool := component.NewTaskPool(20)
	//defer pool.Close()
	for _, proxy := range proxies {
		var p = dto.NewProxyDto(proxy)
		global.Pool.RunTask(func() { service.ProxyService.CheckProxyAndSave(p.ProxyDto) })
	}
}
