package tests

import (
	"testing"

	"proxy-collect/service/ip"
)

func TestGetIpInfoLocal(t *testing.T) {
	result := ip.GetIpInfo("183.129.167.42", "")
	t.Log(result)
}
