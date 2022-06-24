package tests

import (
	"testing"

	"proxy-collect/service/ip"
)

func TestGetIpInfoLocal(t *testing.T) {
	result := ip.GetIpInfo("139.9.64.238", "")
	t.Log(result)
}
