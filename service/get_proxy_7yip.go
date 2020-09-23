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

func NewGetProxy7Yip() *getProxy7Yip {
	return &getProxy7Yip{}
}

type getProxy7Yip struct {
}

func (s *getProxy7Yip) GetUrlList() []string {
	list := []string{
		"https://www.7yip.cn/free/",
	}
	for i := 2; i < 6; i++ {
		list = append(list, fmt.Sprintf("https://www.7yip.cn/free/?action=china&page=%d", i))
	}
	return list
}
func (s *getProxy7Yip) GetContentHtml(requestUrl string) string {
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("User-Agent", config.USER_AGENT)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Host", "www.7yip.cn")
	req.Header.Set("Referer", "https://www.7yip.cn/")
	logger.Info("get proxy from 7yip", requestUrl)
	return component.WebRequest(req)
}

func (s *getProxy7Yip) ParseHtml(body string) [][]string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(err)
		return nil
	}
	var proxyList [][]string
	doc.Find("tbody > tr").Each(func(i int, selection *goquery.Selection) {
		td := selection.ChildrenFiltered("td").First()
		host := strings.TrimSpace(td.Text())
		td2 := selection.ChildrenFiltered("td").Eq(1)
		port := strings.TrimSpace(td2.Text())

		if !ProxyService.CheckProxyFormat(host, port) {
			logger.Error("格式错误:", host+",", port)
			return
		}
		proxyArr := []string{host, port}
		proxyList = append(proxyList, proxyArr)
	})

	return proxyList
}
