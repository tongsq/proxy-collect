package service

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"proxy-collect/config"
	"strings"
)

func NewGetProxyCoderBusy() *getProxyCoderBusy {
	return &getProxyCoderBusy{}
}

type getProxyCoderBusy struct {
}

func (s *getProxyCoderBusy) GetUrlList() []string {
	list := []string{
		"https://proxy.coderbusy.com/",
	}
	return list
}
func (s *getProxyCoderBusy) GetContentHtml(requestUrl string) string {
	h := &request.RequestHeaderDto{
		UserAgent:               config.USER_AGENT,
		UpgradeInsecureRequests: "1",
	}
	logger.Info("get proxy from coderbusy", requestUrl)
	data, err := request.WebGet(requestUrl, h, nil)
	if err != nil || data == nil {
		logger.Error("get proxy from coderbusy fail", err, data)
		return ""
	}
	return data.Body
}

func (s *getProxyCoderBusy) ParseHtml(body string) [][]string {

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
		port := strings.TrimSpace(td2.Find("a").First().Text())

		if !ProxyService.CheckProxyFormat(host, port) {
			logger.Error("格式错误:", host, port)
			return
		}
		proxyArr := []string{host, port}
		proxyList = append(proxyList, proxyArr)
	})
	return proxyList
}
