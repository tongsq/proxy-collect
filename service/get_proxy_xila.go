package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"proxy-collect/component"
	"proxy-collect/component/logger"
	"proxy-collect/config"
	"strings"
)

func NewGetProxyXila() *getProxyXila {
	return &getProxyXila{}
}

type getProxyXila struct {
}

func (s *getProxyXila) GetUrlList() []string {
	list := []string{
		"http://www.xiladaili.com/https/",
	}
	for i := 2; i < 6; i++ {
		list = append(list, fmt.Sprintf("http://www.xiladaili.com/https/%d/", i))
	}
	return list
}

func (s *getProxyXila) GetContentHtml(requestUrl string) string {

	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("User-Agent", config.USER_AGENT)
	req.Header.Set("Host", "www.xiladaili.com")
	req.Header.Set("Referer", "http://www.xiladaili.com/https/")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	logger.Info("get proxy from xiladaili", requestUrl)
	return component.WebRequest(req)
}

func (s *getProxyXila) ParseHtml(body string) [][]string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(err)
		return nil
	}
	var proxy_list [][]string
	doc.Find("tbody > tr").Each(func(i int, selection *goquery.Selection) {
		td := selection.ChildrenFiltered("td").First()
		proxyStr := td.Text()
		proxyStr = strings.Trim(proxyStr, " ")
		proxyArr := strings.Split(proxyStr, ":")
		if len(proxyArr) != 2 {
			logger.Error("格式错误:", proxyStr)
			return
		}
		proxy_list = append(proxy_list, proxyArr)
	})
	return proxy_list
}
