package service

import (
	"github.com/PuerkitoBio/goquery"
	"proxy-collect/component"
	"proxy-collect/component/logger"
	"proxy-collect/config"
	"proxy-collect/dto"
	"strings"
)

func NewGetProxyGuoBanJia() *getProxyGuoBanJia {
	return &getProxyGuoBanJia{}
}

type getProxyGuoBanJia struct {
}

func (s *getProxyGuoBanJia) GetUrlList() []string {
	list := []string{
		"http://www.goubanjia.com/",
	}
	return list
}

func (s *getProxyGuoBanJia) GetContentHtml(requestUrl string) string {
	h := dto.RequestHeaderDto{
		UserAgent:               config.USER_AGENT,
		Host:                    "www.goubanjia.com",
		UpgradeInsecureRequests: "1",
	}
	logger.Info("get proxy from guobanjia", requestUrl)
	return component.WebGet(requestUrl, h)
}

func (s *getProxyGuoBanJia) ParseHtml(body string) [][]string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(err)
		return nil
	}
	var proxyList [][]string
	doc.Find("tbody > tr").Each(func(i int, selection *goquery.Selection) {
		td := selection.ChildrenFiltered("td").First()
		hostStr := ""
		len := td.Children().Size()
		td.Children().Each(func(i int, item *goquery.Selection) {
			style, _ := item.Attr("style")
			if !strings.Contains(style, "none") && i != (len-1) {
				hostStr = hostStr + item.Text()
			}
		})
		port := td.Children().Last().Text()
		hostStr = strings.TrimSpace(hostStr)
		if !ProxyService.CheckProxyFormat(hostStr, port) {
			logger.Error("格式错误:", hostStr, port)
			return
		}
		proxyArr := []string{hostStr, port}
		proxyList = append(proxyList, proxyArr)
	})

	return proxyList
}
