package tests

import (
	"math/rand"
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
		{"223.96.90.216", "8085"},
		{"39.96.11.1", "8003"},
		{"47.106.105.236", "80"},
		{"1.180.156.226", "65001"},
		{"221.122.91.34", "80"},
		{"153.35.185.69", "80"},
	}
	for _, item := range items {
		r := service.ProxyService.CheckIpStatus(item[0], item[1])
		t.Log(r)
	}
}
