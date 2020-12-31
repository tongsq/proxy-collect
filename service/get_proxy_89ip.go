package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"proxy-collect/config"
	"strings"
)

func NewGetProxy89Ip() *getProxy89Ip {
	return &getProxy89Ip{}
}

type getProxy89Ip struct {
}

func (s *getProxy89Ip) GetUrlList() []string {
	list := []string{
		"https://www.89ip.cn/",
	}
	for i := 2; i < 6; i++ {
		list = append(list, fmt.Sprintf("https://www.89ip.cn/index_%d.html", i))
	}
	return list
}
func (s *getProxy89Ip) GetContentHtml(requestUrl string) string {
	h := &request.RequestHeaderDto{
		UserAgent:               config.USER_AGENT,
		UpgradeInsecureRequests: "1",
		Host:                    "www.89ip.cn",
		Referer:                 "https://www.89ip.cn/",
	}

	logger.Info("get proxy from 89ip", requestUrl)
	data, err := request.WebGet(requestUrl, h, nil)
	if err != nil || data == nil {
		logger.Error("get proxy from 89ip fail", err, data)
		return ""
	}
	return data.Body
}

func (s *getProxy89Ip) ParseHtml(body string) [][]string {

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
