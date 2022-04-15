package tests

import (
	"sync"
	"testing"

	"proxy-collect/consts"
	"proxy-collect/service"
)

//https://pzzqz.com/
func TestCheckIp(t *testing.T) {
	items := [][]string{
		{"localhost", "9999", consts.PROTO_HTTPS},
		{"localhost", "8888", consts.PROTO_HTTP},
		{"localhost", "8899", consts.PROTO_SOCKS4},
		{"122.193.10.184", "7300", consts.PROTO_SOCKS5},
	}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for _, item := range items {
			proxy := service.ProxyService.ParseProxyArr(item)
			r := service.ProxyService.CheckIpStatus(&proxy)
			t.Log(item, r)
		}
		defer wg.Done()
	}()
	wg.Wait()
}
