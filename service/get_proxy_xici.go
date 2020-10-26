package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"proxy-collect/component"
	"proxy-collect/component/logger"
	"proxy-collect/config"
	"proxy-collect/dto"
	"strings"
)

func NewGetProxyXici() *getProxyXici {
	return &getProxyXici{}
}

type getProxyXici struct {
}

func (s *getProxyXici) GetUrlList() []string {
	list := []string{
		"http://www.xicidaili.com/wn/",
	}
	for i := 2; i < 6; i++ {
		list = append(list, fmt.Sprintf("http://www.xicidaili.com/wn/%d", i))
	}
	return list
}

func (s *getProxyXici) GetContentHtml(requestUrl string) string {

	h := dto.RequestHeaderDto{
		UserAgent:               config.USER_AGENT,
		Host:                    "www.xicidaili.com",
		Referer:                 "http://www.xicidaili.com",
		UpgradeInsecureRequests: "1",
	}

	logger.Info("get proxy from xicidaili", requestUrl)
	return component.WebGet(requestUrl, h)
}

func (s *getProxyXici) ParseHtml(body string) [][]string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(err)
		return nil
	}
	var proxyList [][]string
	doc.Find("tbody > tr").Each(func(i int, selection *goquery.Selection) {
		td := selection.ChildrenFiltered("td").First()
		proxyStr := td.Text()
		proxyStr = strings.TrimSpace(proxyStr)
		proxyArr := strings.Split(proxyStr, ":")
		if len(proxyArr) != 2 {
			logger.Error("格式错误:", proxyStr)
			return
		}
		proxyList = append(proxyList, proxyArr)
	})
	return proxyList
}
