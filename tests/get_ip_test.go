package tests

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/tongsq/go-lib/logger"
	"proxy-collect/bootstrap"
	"proxy-collect/service"
)

func TestGetIp(t *testing.T) {
	bootstrap.Bootstrap()
	s := service.CommonGetterHttp
	wg := sync.WaitGroup{}
	succ := 0
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
				r, ping := service.ProxyService.CheckIpStatus(&p)
				t.Log(p, r, ping)
				if r {
					succ++
				}
			}()
		}
	}
	wg.Wait()
	t.Log("success count:", succ)
}

func TestRand(t *testing.T) {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(10)
	t.Log(i)
}

func TestAto(t *testing.T) {
	var num int64 = 0
	var num2 int64 = 0
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			num2++
			atomic.AddInt64(&num, 1)
		}()
	}
	wg.Wait()
	t.Log(num, num2)
}
