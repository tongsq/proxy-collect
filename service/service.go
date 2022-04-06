package service

import (
	"proxy-collect/service/get_proxy"
	"proxy-collect/service/ip"
)

var ProxyService = NewProxyService()
var GetProxy66ip = get_proxy.NewGetProxy66ip()
var GetProxyData5u = get_proxy.NewGetProxyData5u()
var GetProxyKuai = get_proxy.NewGetProxyKuai()
var GetProxyXila = get_proxy.NewGetProxyXila()
var GetProxyNima = get_proxy.NewGetProxyNima()
var GetProxyGuoBanjia = get_proxy.NewGetProxyGuoBanJia()
var GetProxyCoderBusy = get_proxy.NewGetProxyCoderBusy()
var GetProxyIp3366 = get_proxy.NewGetProxyIp3366()
var GetProxyIpJiangXianLi = get_proxy.NewGetProxyIpJiangXianLi()
var GetProxy89Ip = get_proxy.NewGetProxy89Ip()
var GetProxy7Yip = get_proxy.NewGetProxy7Yip()
var GetProxyProxyList = get_proxy.NewGetProxyProxyList()
var GetProxyZdaye = get_proxy.NewGetProxyZdaye()
var GetProxyZdayeIndex = get_proxy.NewGetProxyZdayeIndex()
var GetProxyFanQie = get_proxy.NewGetProxyFanQie()
var GetProxySeofangfa = get_proxy.NewGetProxySeofangfa()
var GetProxyXsdaili = get_proxy.NewGetProxyXsdaili()
var GetProxyYqie = get_proxy.NewGetProxyYqie()
var GetProxyPaChong = get_proxy.NewGetProxyPachong()
var KxDaili = get_proxy.NewGetProxyKxDaili()

func LoadService() {
	ip.LoadLocalIpData()
}
