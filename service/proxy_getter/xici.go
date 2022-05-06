package proxy_getter

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"proxy-collect/consts"
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

	h := &request.HeaderDto{
		UserAgent:               consts.USER_AGENT,
		Host:                    "www.xicidaili.com",
		Referer:                 "http://www.xicidaili.com",
		UpgradeInsecureRequests: "1",
	}

	logger.Info("get proxy from xicidaili", logger.Fields{"url": requestUrl})
	data, err := request.Get(requestUrl, request.NewOptions().WithHeader(h))
	if err != nil || data == nil {
		logger.Error("get proxy from xicidaili fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *getProxyXici) ParseHtml(body string) [][]string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(consts.GO_QUERY_READ_ERROR, logger.Fields{"err": err})
		return nil
	}
	var proxyList [][]string
	doc.Find("tbody > tr").Each(func(i int, selection *goquery.Selection) {
		td := selection.ChildrenFiltered("td").First()
		proxyStr := td.Text()
		proxyStr = strings.TrimSpace(proxyStr)
		proxyArr := strings.Split(proxyStr, ":")
		if len(proxyArr) != 2 {
			logger.Error(consts.PROXY_FORMAT_ERROR, logger.Fields{"proxyStr": proxyStr})
			return
		}
		proxyList = append(proxyList, proxyArr)
	})
	return proxyList
}
