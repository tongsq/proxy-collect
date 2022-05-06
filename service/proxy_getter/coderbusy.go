package proxy_getter

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/tongsq/go-lib/logger"
	"proxy-collect/consts"
	"proxy-collect/global"
	"proxy-collect/service/common"

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
	logger.Info("get proxy from coderbusy", logger.Fields{"url": requestUrl})
	data, err := global.SimpleGet(requestUrl)
	if err != nil || data == nil {
		logger.Error("get proxy from coderbusy fail", logger.Fields{"err": err, "data": data})
		logger.Error("get proxy from coderbusy fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *getProxyCoderBusy) ParseHtml(body string) [][]string {

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
		port := strings.TrimSpace(td2.Find("a").First().Text())

		if !common.CheckProxyFormat(host, port) {
			logger.Error(consts.PROXY_FORMAT_ERROR, logger.Fields{"host": host, "port": port})
			return
		}
		proxyArr := []string{host, port}
		proxyList = append(proxyList, proxyArr)
	})
	return proxyList
}
