package service

import (
	"proxy-collect/service/ip"
	"proxy-collect/service/proxy_getter"
)

var ProxyService = NewProxyService()
var GetProxy66ip = proxy_getter.NewGetProxy66ip()
var GetProxyData5u = proxy_getter.NewGetProxyData5u()
var GetProxyKuai = proxy_getter.NewGetProxyKuai()
var GetProxyXila = proxy_getter.NewGetProxyXila()
var GetProxyNima = proxy_getter.NewGetProxyNima()
var GetProxyGuoBanjia = proxy_getter.NewGetProxyGuoBanJia()
var GetProxyCoderBusy = proxy_getter.NewGetProxyCoderBusy()
var GetProxyIp3366 = proxy_getter.NewGetProxyIp3366()
var GetProxyIpJiangXianLi = proxy_getter.NewGetProxyIpJiangXianLi()
var GetProxy89Ip = proxy_getter.NewGetProxy89Ip()
var GetProxy7Yip = proxy_getter.NewGetProxy7Yip()
var GetProxyProxyList = proxy_getter.NewGetProxyProxyList()
var GetProxyZdaye = proxy_getter.NewGetProxyZdaye()
var GetProxyZdayeIndex = proxy_getter.NewGetProxyZdayeIndex()
var GetProxyFanQie = proxy_getter.NewGetProxyFanQie()
var GetProxySeofangfa = proxy_getter.NewGetProxySeofangfa()
var GetProxyXsdaili = proxy_getter.NewGetProxyXsdaili()
var GetProxyYqie = proxy_getter.NewGetProxyYqie()
var GetProxyPaChong = proxy_getter.NewGetProxyPachong()
var KxDaili = proxy_getter.NewGetProxyKxDaili()
var Geonode = proxy_getter.NewGetProxyGeonode()

func LoadService() {
	ip.LoadLocalIpData()
}
