package proxy_getter

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"proxy-collect/consts"
	"proxy-collect/service/common"
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
	h := &request.HeaderDto{
		UserAgent:               consts.USER_AGENT,
		UpgradeInsecureRequests: "1",
		Host:                    "www.89ip.cn",
		Referer:                 "https://www.89ip.cn/",
	}

	logger.Info("get proxy from 89ip", logger.Fields{"url": requestUrl})
	data, err := request.WebGet(requestUrl, h, nil)
	if err != nil || data == nil {
		logger.Error("get proxy from 89ip fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *getProxy89Ip) ParseHtml(body string) [][]string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(consts.GO_QUERY_READ_ERROR, logger.Fields{"err": err})
		return nil
	}
	var proxyList [][]string
	doc.Find("tbody > tr").Each(func(i int, selection *goquery.Selection) {
		td := selection.ChildrenFiltered("td").First()
		host := strings.TrimSpace(td.Text())
		td2 := selection.ChildrenFiltered("td").Eq(1)
		port := strings.TrimSpace(td2.Text())

		if !common.CheckProxyFormat(host, port) {
			logger.Error(consts.PROXY_FORMAT_ERROR, logger.Fields{"host": host, "port": port})
			return
		}
		proxyArr := []string{host, port}
		proxyList = append(proxyList, proxyArr)
	})
	return proxyList
}
