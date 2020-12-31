package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"proxy-collect/config"
	"strings"
)

func NewGetProxyNima() *getProxyNima {
	return &getProxyNima{}
}

type getProxyNima struct {
}

func (s *getProxyNima) GetUrlList() []string {
	list := []string{
		"http://www.nimadaili.com/https/",
	}
	for i := 2; i < 6; i++ {
		list = append(list, fmt.Sprintf("http://www.nimadaili.com/https/%d/", i))
	}
	return list
}
func (s *getProxyNima) GetContentHtml(requestUrl string) string {
	h := &request.RequestHeaderDto{
		UserAgent:               config.USER_AGENT,
		Host:                    "www.nimadaili.com",
		Referer:                 "http://www.nimadaili.com/https/3/",
		UpgradeInsecureRequests: "1",
	}

	logger.Info("get proxy from nimadaili", requestUrl)
	data, err := request.WebGet(requestUrl, h, nil)
	if err != nil || data == nil {
		logger.Error("get proxy from nimadaili fail", err, data)
		return ""
	}
	return data.Body
}

func (s *getProxyNima) ParseHtml(body string) [][]string {

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
