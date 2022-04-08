package tests

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/tongsq/go-lib/logger"
	"proxy-collect/service"
)

func TestGetIp(t *testing.T) {
	s := service.Geonode
	wg := sync.WaitGroup{}
	for _, requestUrl := range s.GetUrlList() {
		t.Log(requestUrl)
		contentBody := s.GetContentHtml(requestUrl)
		if contentBody == "" {
			time.Sleep(time.Second * 5)
			continue
		}
		proxyList := s.ParseHtml(contentBody)
		logger.Info("get ip list:", logger.Fields{"list": proxyList})
		for _, item := range proxyList {
			p := service.ProxyService.ParseProxyArr(item)
			wg.Add(1)
			go func() {
				defer wg.Done()
				r := service.ProxyService.CheckIpStatus(service.ProxyService.GetProxyUrl(p))
				t.Log(p, r)
			}()
		}
	}
	wg.Wait()
}

func TestRand(t *testing.T) {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(10)
	t.Log(i)
}

//https://pzzqz.com/
func TestCheckIp(t *testing.T) {
	items := []string{
		"socks5://root:123@localhost:9988",
		"socks4://root:123@localhost:8899",
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

func TestAto(t *testing.T) {
	var num int64 = 0
	var num2 int64 = 0
	for i := 0; i <= 10000; i++ {
		//atomic.AddInt64(&num, 1)
		go func() {
			num2++
			atomic.AddInt64(&num, 1)
		}()
	}
	time.Sleep(10)
	t.Log(num, num2)
}
