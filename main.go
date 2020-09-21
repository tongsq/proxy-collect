package main

import (
	"github.com/robfig/cron/v3"
	"proxy-collect/component/logger"
	"proxy-collect/scheduler"
)

func main() {
	l := &logger.CronLogger{}
	c := cron.New(cron.WithChain(
		cron.SkipIfStillRunning(l),
	))
	c.AddJob("@every 3m", scheduler.GetIp{})
	c.AddJob("@every 3h", scheduler.CheckIp{})
	c.Start()
	defer c.Stop()
	select {}
}
