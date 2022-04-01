package global

import (
	"github.com/tongsq/go-lib/component"
	"proxy-collect/config"
)

var Pool *component.Pool

func LoadGlobal() {
	Pool = component.NewTaskPool(config.Get().PoolSize)
}
