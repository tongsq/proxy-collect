package main

import (
	"fmt"
	"github.com/tongsq/go-lib/component"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/util"
	"proxy-collect/service"
	"time"
)

var num int64 = 0

func add() {
	num = num + 1
}

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(util.Add(1, 2, 3))
	return
	//time.Sleep(50 * time.Second)
	//scheduler.UpdateIpInfo{}.Run()
	//res := service.ProxyService.CheckIpStatus("129.28.173.182", "8388")
	//logger.Info(res)
	//go func() {
	//	fmt.Println("pprof start...")
	//	fmt.Println(http.ListenAndServe(":9876", nil))
	//}()
	//go test()
	//for true {
	//	logger.Info(runtime.NumGoroutine())
	//	time.Sleep(time.Second)
	//}
	//scheduler.UpdateIpInfo{}.Run()
	ipInfoDto := service.ProxyService.GetIpInfo("35.196.118.22", "80")
	logger.Info("get ip info:", ipInfoDto)
}

func test() {
	pool := component.NewTaskPool(10000)
	defer pool.Close()
	for i := 0; i < 100000; i++ {
		pool.RunTask(func() {
			logger.Info("hello")
		})
	}
}
