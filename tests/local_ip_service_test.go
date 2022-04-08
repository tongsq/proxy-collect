package tests

import (
	"testing"

	"proxy-collect/service/ip"
)

func TestGetIpInfoLocal(t *testing.T) {
	result := ip.GetIpInfo("113.214.48.5", "")
	t.Log(result)
}
