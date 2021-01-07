package service

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"proxy-collect/consts"
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
	h := &request.RequestHeaderDto{
		UserAgent:               consts.USER_AGENT,
		Host:                    "www.goubanjia.com",
		UpgradeInsecureRequests: "1",
	}
	logger.Info("get proxy from guobanjia", logger.Fields{"url": requestUrl})
	data, err := request.WebGet(requestUrl, h, nil)
	if err != nil || data == nil {
		logger.Error("get proxy from guobanjia fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *getProxyGuoBanJia) ParseHtml(body string) [][]string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error("read fail", logger.Fields{"err": err})
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
			logger.Error(consts.PROXY_FORMAT_ERROR, logger.Fields{"host": hostStr, "port": port})
			return
		}
		proxyArr := []string{hostStr, port}
		proxyList = append(proxyList, proxyArr)
	})

	return proxyList
}
