package main

import (
	"proxy-collect/component"
	"proxy-collect/service"
)

var num int64 = 0

func add() {
	num = num + 1
}

func main() {
	pool := component.NewTaskPool(20)
	service.ProxyService.DoGetProxy(service.GetProxyZdaye, pool)
	//time.Sleep(50 * time.Second)
	//scheduler.UpdateIpInfo{}.Run()
	//res := service.ProxyService.CheckIpStatus("129.28.173.182", "8388")
	//logger.Info(res)
}
