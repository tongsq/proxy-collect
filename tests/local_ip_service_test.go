package tests

import (
	"testing"

	"proxy-collect/service/ip"
)

func TestGetIpInfoLocal(t *testing.T) {
	result := ip.GetIpInfo("95.113.214.48", "")
	t.Log(result)
}
