package service

import (
	"github.com/PuerkitoBio/goquery"
	"proxy-collect/component"
	"proxy-collect/component/logger"
	"proxy-collect/config"
	"proxy-collect/dto"
	"strings"
)

func NewGetProxyData5u() *getProxyData5u {
	return &getProxyData5u{}
}

type getProxyData5u struct {
}

func (s *getProxyData5u) GetUrlList() []string {
	list := []string{
		"http://www.data5u.com/",
		"http://www.data5u.com/free/gngn/index.shtml",
		"http://www.data5u.com/free/gnpt/index.shtml",
	}
	return list
}

func (s *getProxyData5u) GetContentHtml(requestUrl string) string {
	h := dto.RequestHeaderDto{
		UserAgent:               config.USER_AGENT,
		Host:                    "www.data5u.com",
		UpgradeInsecureRequests: "1",
	}

	logger.Info("get proxy from data5u", requestUrl)
	return component.WebGet(requestUrl, h)
}

func (s *getProxyData5u) ParseHtml(body string) [][]string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(err)
		return nil
	}

	var proxyList [][]string
	doc.Find("ul.l2").Each(func(i int, selection *goquery.Selection) {
		td := selection.Find("span>li").First()
		proxyHost := td.Text()
		td2 := selection.Find("span>li").Eq(1)
		proxyPort := td2.Text()
		if proxyHost == "" || proxyPort == "" {
			logger.Error("解析html node 失败")
		}
		proxyArr := []string{strings.TrimSpace(proxyHost), strings.TrimSpace(proxyPort)}
		proxyList = append(proxyList, proxyArr)
	})

	return proxyList
}
