package main

import (
	"fmt"
	"github.com/tongsq/go-lib/logger"
	"proxy-collect/dao"
)

func main() {
	logger.FSuccess("check active ip start run")
	proxies, err := dao.ProxyDao.GetActiveList()
	if err != nil {
		logger.Error("get active ip fail", logger.Fields{"err": err})
	}
	logger.FInfo("check active ip, len:%d, cap:%d", len(proxies), cap(proxies))
	//pool := component.NewTaskPool(20)
	//defer pool.Close()
	for _, proxy := range proxies {
		//var proxyTmp model.ProxyModel = proxy
		proxy := proxy
		go func(h, p string) {
			fmt.Println(h, p)
			fmt.Println(proxy.Host, proxy.Port, "a")
		}(proxy.Host, proxy.Port)
		//component.TaskPool.RunTask(func() { service.ProxyService.CheckProxyAndSave(proxyTmp.Host, proxyTmp.Port, "") })
	}
}
