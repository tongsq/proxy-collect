package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"proxy-collect/config"
	"proxy-collect/consts"
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

	h := &request.RequestHeaderDto{
		UserAgent:               config.USER_AGENT,
		Host:                    "www.xiladaili.com",
		Referer:                 "http://www.xiladaili.com/https/",
		UpgradeInsecureRequests: "1",
	}
	logger.Info("get proxy from xiladaili", logger.Fields{"url": requestUrl})
	data, err := request.WebGet(requestUrl, h, nil)
	if err != nil || data == nil {
		logger.Error("get proxy from xiladaili fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *getProxyXila) ParseHtml(body string) [][]string {
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
