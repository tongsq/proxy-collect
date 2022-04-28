package global

import (
	"time"

	"github.com/tongsq/go-lib/component"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"proxy-collect/config"
)

var Pool *component.Pool

func LoadGlobal() {
	Pool = component.NewTaskPool(config.Get().PoolSize)
	logger.SetLogLevel(config.Get().Log.LogLevel)
	if config.Get().Log.ErrorLogFile != "" {
		logger.SetErrorFile(config.Get().Log.ErrorLogFile)
	}
	//set max proxy timeout
	request.SetTimeout(time.Millisecond * time.Duration(config.Get().MaxPing))
}
