package tests

import (
	"testing"

	"proxy-collect/service"
)

func TestGetIpInfoLocal(t *testing.T) {
	//result := ip.GetIpInfo("113.214.48.5", "")
	//t.Log(result)
	service.ProxyService.CheckProxyAndSave("20.47.108.204", "8888", "a")
}
