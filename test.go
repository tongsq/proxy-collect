package main

import (
	"proxy-collect/scheduler"
)

var num int64 = 0

func add() {
	num = num + 1
}

func main() {
	//pool := component.NewTaskPool(20)
	//service.ProxyService.DoGetProxy(service.GetProxyProxyList, pool)
	//time.Sleep(50 * time.Second)
	scheduler.UpdateIpInfo{}.Run()
}
