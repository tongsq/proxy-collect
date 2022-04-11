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

func NewGetProxyIpJiangXianLi() *getProxyIpJiangXianLi {
	return &getProxyIpJiangXianLi{}
}

type getProxyIpJiangXianLi struct {
}

func (s *getProxyIpJiangXianLi) GetUrlList() []string {
	list := []string{
		"https://ip.jiangxianli.com/",
	}
	for i := 2; i < 6; i++ {
		list = append(list, fmt.Sprintf("https://ip.jiangxianli.com/?page=%d", i))
	}
	return list
}
func (s *getProxyIpJiangXianLi) GetContentHtml(requestUrl string) string {
	h := &request.HeaderDto{
		UserAgent:               consts.USER_AGENT,
		UpgradeInsecureRequests: "1",
		Host:                    "ip.jiangxianli.com",
		Referer:                 "https://ip.jiangxianli.com/",
	}
	logger.Info("get proxy from jangxianli", logger.Fields{"url": requestUrl})
	data, err := request.WebGet(requestUrl, h, nil)
	if err != nil || data == nil {
		logger.Error("get proxy from jangxianli fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *getProxyIpJiangXianLi) ParseHtml(body string) [][]string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error("read fail", logger.Fields{"err": err})
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
