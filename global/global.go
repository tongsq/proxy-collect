package global

import (
	"time"

	"github.com/tongsq/go-lib/component"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"proxy-collect/config"
	"proxy-collect/consts"
)

var Pool *component.Pool
var MaxPing time.Duration = request.DefaultTimeout

var CommonHeader = &request.HeaderDto{
	UserAgent: consts.USER_AGENT,
}

func LoadGlobal() {
	Pool = component.NewTaskPool(config.Get().PoolSize)
	logger.SetLogLevel(config.Get().Log.LogLevel)
	if config.Get().Log.ErrorLogFile != "" {
		logger.SetErrorFile(config.Get().Log.ErrorLogFile)
	}
	//set max proxy timeout
	MaxPing = time.Millisecond * time.Duration(config.Get().MaxPing)

	config.RegisterConfigRefreshHandler(func(old, new *config.ConfDto) {
		if old.Log.LogLevel != new.Log.LogLevel {
			logger.SetLogLevel(new.Log.LogLevel)
		}
		if old.Log.ErrorLogFile != new.Log.ErrorLogFile && new.Log.ErrorLogFile != "" {
			logger.SetErrorFile(new.Log.ErrorLogFile)
		}
		MaxPing = time.Millisecond * time.Duration(new.MaxPing)
	})
}

func SimpleGet(requestUrl string) (*request.HttpResultDto, error) {
	return request.Get(requestUrl, request.NewOptions().WithHeader(CommonHeader))
}
