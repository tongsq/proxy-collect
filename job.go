package main

import (
	"github.com/robfig/cron/v3"
	"proxy-collect/scheduler"
)

func main() {
	//l := &logger.CronLogger{}
	//c := cron.New(cron.WithChain(
	//	cron.SkipIfStillRunning(l),
	//))
	c := cron.New()
	c.AddJob("@every 3m", scheduler.GetIp{})
	c.AddJob("@every 2m", scheduler.CheckActiveIp{})
	c.AddJob("@delay 5h", scheduler.CheckFailIp{})
	c.Start()
	defer c.Stop()
	select {}
}
