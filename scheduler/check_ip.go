package scheduler

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"proxy-collect/component"
	"proxy-collect/model"
	"proxy-collect/service"
)

type CheckIp struct {
}

func (s CheckIp) Run() {
	var proxys []model.Proxy
	model.DB.Where("status=?", 1).Find(&proxys)
	fmt.Printf("count:%d, cap: %d\n", len(proxys), cap(proxys))
	pool := component.NewTaskPool(20)
	for _, proxy := range proxys {
		var proxyTmp model.Proxy = proxy
		pool.RunTask(func() { s.CheckProxyStatus(proxyTmp, model.DB) })
	}
}

func (s CheckIp) CheckProxyStatus(proxyModel model.Proxy, db *gorm.DB) {
	fmt.Printf("start check :host:%s, port:%s\n", proxyModel.Host, proxyModel.Port)
	result := service.ProxyService.CheckIpStatus(proxyModel.Host, proxyModel.Port)
	fmt.Printf("%s, %s, the result is %v\n", proxyModel.Host, proxyModel.Port, result)
	if result {
		proxyModel.Status = 1
		db.Save(&proxyModel)
	} else {
		proxyModel.Status = 0
		db.Save(&proxyModel)
	}
}
