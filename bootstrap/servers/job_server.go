package servers

import (
	"github.com/robfig/cron/v3"
	"github.com/tongsq/go-lib/logger"
	"proxy-collect/scheduler"
	"proxy-collect/service/ip"
)

type JobItem struct {
	Spec string
	Job  cron.Job
}

func StartJobServer() {
	c := cron.New()
	jobs := []JobItem{
		{"*/3 * * * *", scheduler.GetIp{}},
		{"*/2 * * * *", scheduler.CheckActiveIp{}},
		{"@every 5h", scheduler.CheckFailIp{}},
		{"@every 5m", scheduler.UpdateIpInfo{}},
		{"@every 1m", scheduler.RecheckIp{}},
	}
	for _, job := range jobs {
		_, err := c.AddJob(job.Spec, job.Job)
		if err != nil {
			logger.Error("start job fail", map[string]interface{}{"job": job})
		}
	}

	//update local ip database
	_, err := c.AddFunc("2 2 * * *", func() {
		ip.UpdateLocalIpData()
	})
	if err != nil {
		logger.Error("start job fail", map[string]interface{}{})
	}
	c.Start()
	defer c.Stop()
	scheduler.GetIp{}.Run()
	select {}
}
