package scheduler

import (
	"encoding/json"
	"github.com/tongsq/go-lib/logger"
	"proxy-collect/config"
	"proxy-collect/global"
	"proxy-collect/service"
	"proxy-collect/service/proxy_getter"
)

type GetIp struct {
}

func (s GetIp) Run() {
	logger.FSuccess("collect ip start run")
	//pool := component.NewTaskPool(100)
	//defer pool.Close()
	serviceList := []service.ProxyGetterInterface{
		service.GetProxyXila,
		service.GetProxyNima,
		service.GetProxyKuai,
		service.GetProxyData5u,
		service.GetProxy66ip,
		//service.GetProxyGuoBanjia,
		service.GetProxyCoderBusy,
		service.GetProxyIp3366,
		service.GetProxyIpJiangXianLi,
		service.GetProxy89Ip,
		service.GetProxy7Yip,
		service.GetProxyProxyList,
		service.GetProxyZdaye,
		service.GetProxyFanQie,
		//service.GetProxyZdayeIndex,
		service.GetProxySeofangfa,
		service.GetProxyXsdaili,
		service.GetProxyPaChong,
		service.KxDaili,
		service.Geonode,
	}
	var configGetterList []service.ProxyGetterInterface
	for _, conf := range config.Get().Getters {
		configGetterList = append(configGetterList, proxy_getter.NewGetter(&conf))
	}
	config.RegisterConfigRefreshHandler(func(old, new *config.ConfDto) {
		var oldConfig, newConfig []byte
		oldConfig, _ = json.Marshal(old.Getters)
		newConfig, _ = json.Marshal(new.Getters)
		if string(oldConfig) != string(newConfig) {
			logger.Info("reload getter configs", map[string]interface{}{"old": string(oldConfig), "new": string(newConfig)})
			configGetterList = []service.ProxyGetterInterface{}
			for _, conf := range new.Getters {
				configGetterList = append(configGetterList, proxy_getter.NewGetter(&conf))
			}
		}
	})
	for _, getProxyService := range serviceList {
		go service.ProxyService.DoGetProxy(getProxyService, global.Pool)
	}
	for _, configGetter := range configGetterList {
		go service.ProxyService.DoGetProxy(configGetter, global.Pool)
	}
}
