package tests

import (
	"math/rand"
	"testing"
	"time"

	"github.com/tongsq/go-lib/logger"
	"proxy-collect/service"
)

func TestGetIp(t *testing.T) {
	s := service.GetProxySeofangfa
	for _, requestUrl := range s.GetUrlList() {
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
		{"213.101.151.4", "1080"},
		{"177.54.195.48", "4145"},
		{"197.232.47.102", "52567"},
		{"70.35.213.226", "4153"},
		{"117.4.145.16", "41889"},
		{"185.94.218.57", "44421"},
	}
	for _, item := range items {
		r := service.ProxyService.CheckIpStatus(item[0], item[1])
		t.Log(r)
	}
}
