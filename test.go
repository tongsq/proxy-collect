package main

import (
	"fmt"
	"proxy-collect/service"
)

var num int64 = 0

func add() {
	num = num + 1
}

func main() {
	//pool := component.NewTaskPool(20)
	//service.ProxyService.DoGetProxy(service.GetProxyProxyList, pool)
	//time.Sleep(50 * time.Second)
	info := service.ProxyService.GetIpInfo("110.243.0.176", "9999")
	fmt.Println(info)
	//scheduler.UpdateIpInfo{}.Run()
}
