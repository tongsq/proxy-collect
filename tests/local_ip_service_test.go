package tests

import (
	"testing"

	"proxy-collect/service/ip"
)

func TestGetIpInfoLocal(t *testing.T) {
	result := ip.GetIpInfo("59.66.190.25", "")
	t.Log(result)
}
