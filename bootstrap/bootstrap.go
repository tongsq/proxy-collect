package bootstrap

import (
	"proxy-collect/config"
	"proxy-collect/dao"
	"proxy-collect/global"
	"proxy-collect/service"
)

func Bootstrap() {
	config.LoadConfig()
	global.LoadGlobal()
	dao.LoadDao()
	service.LoadService()
}
