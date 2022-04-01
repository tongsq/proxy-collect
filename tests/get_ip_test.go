package tests

import (
	"math/rand"
	"sync/atomic"
	"testing"
	"time"

	"github.com/tongsq/go-lib/logger"
	"proxy-collect/service"
)

func TestGetIp(t *testing.T) {
	s := service.GetProxyPaChong
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
			r := service.ProxyService.CheckIpStatus(item[0], item[1])
			t.Log(r)
		}
	}
}

func TestRand(t *testing.T) {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(10)
	t.Log(i)
}

func TestCheckIp(t *testing.T) {
	items := [][]string{
		{"root:123@127.0.0.1", "8090"},
	}
	for i := 0; i < 10; i++ {
		go func() {
			for _, item := range items {
				r := service.ProxyService.CheckIpStatus(item[0], item[1])
				t.Log(r)
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
