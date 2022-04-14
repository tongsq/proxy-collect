package tests

import (
	"testing"
	"time"

	"proxy-collect/service"
)

//https://pzzqz.com/
func TestCheckIp(t *testing.T) {
	items := []string{
		"https://localhost:9999",
		"http://localhost:8888",
		"socks4://127.0.0.1:8899",
		"socks5://122.193.10.184:7300",
	}
	for i := 0; i < 1; i++ {
		go func() {
			for _, item := range items {
				r := service.ProxyService.CheckIpStatus(item)
				t.Log(item, r)
			}
		}()
	}
	time.Sleep(time.Second * 10)
}
