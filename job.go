package main

import (
	"github.com/robfig/cron/v3"
	"github.com/tongsq/go-lib/logger"
	"proxy-collect/scheduler"
	"runtime"
	"time"
)

func main() {
	//l := &logger.CronLogger{}
	//c := cron.New(cron.WithChain(
	//	cron.SkipIfStillRunning(l),
	//))
	c := cron.New()
	c.AddJob("*/3 * * * *", scheduler.GetIp{})
	c.AddJob("*/2 * * * *", scheduler.CheckActiveIp{})
	c.AddJob("@every 5h", scheduler.CheckFailIp{})
	c.AddJob("@every 5m", scheduler.UpdateIpInfo{})
	c.Start()
	defer c.Stop()
	for true {
		logger.FWarning("routine num : %d", runtime.NumGoroutine())
		time.Sleep(time.Second)
	}
}
