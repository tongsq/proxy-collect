package main

import (
	"proxy-collect/component"
	"proxy-collect/service"
	"time"
)

func main() {
	pool := component.NewTaskPool(20)
	service.ProxyService.DoGetProxy(service.GetProxyProxyList, pool)
	time.Sleep(50 * time.Second)

}
