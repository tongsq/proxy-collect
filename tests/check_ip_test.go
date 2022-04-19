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
		{"localhost", "9988", consts.PROTO_SOCKS5, "root", "123"},
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

func TestCheckIpBatch(t *testing.T) {
	item := []string{"0.0.0.0", "8888", consts.PROTO_HTTP, "root", "123"}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			proxy := service.ProxyService.ParseProxyArr(item)
			r := service.ProxyService.CheckIpStatus(&proxy)
			t.Log(item, r)
			defer wg.Done()
		}()
	}
	wg.Wait()
}
