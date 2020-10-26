package main

import (
	"github.com/robfig/cron/v3"
	"proxy-collect/component/logger"
	"time"
)

type Test1 struct {
}

func (s Test1) Run() {
	for true {
		logger.Info("hello")
		time.Sleep(time.Second)
	}
}

type Test2 struct {
}

func (s Test2) Run() {

	logger.Info("hello22222")

}
func main() {
	l := &logger.CronLogger{}
	c := cron.New(cron.WithChain(
		cron.SkipIfStillRunning(l),
	))
	//c := cron.New()
	c.AddJob("@every 1m", Test1{})
	c.AddJob("@every 1m", Test2{})
	c.Start()
	defer c.Stop()
	select {}
}
