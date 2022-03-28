package main

import (
	"github.com/robfig/cron/v3"
	"proxy-collect/scheduler"
	"proxy-collect/service/ip"
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
	c.AddJob("@every 1m", scheduler.RecheckIp{})
	//update local ip database
	c.AddFunc("2 2 * * *", func() {
		ip.UpdateLocalIpData()
	})
	c.Start()
	defer c.Stop()
	scheduler.GetIp{}.Run()
	select {}
}
