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
	c.AddJob("*/3 * * * *", scheduler.GetIp{})
	c.AddJob("*/2 * * * *", scheduler.CheckActiveIp{})
	c.AddJob("@every 5h", scheduler.CheckFailIp{})
	c.AddJob("@every 5m", scheduler.UpdateIpInfo{})
	c.Start()
	defer c.Stop()
	scheduler.GetIp{}.Run()
	select {}
}
