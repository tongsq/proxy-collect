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

func NewGetProxyKxDaili() *getProxyKxDaili {
	return &getProxyKxDaili{}
}

type getProxyKxDaili struct {
}

func (s *getProxyKxDaili) GetUrlList() []string {
	list := []string{
		"http://www.kxdaili.com/dailiip.html",
		"http://www.kxdaili.com/dailiip/2/1.html",
	}
	for i := 2; i < 5; i++ {
		list = append(list, fmt.Sprintf("http://www.kxdaili.com/dailiip/1/%d.html", i))
		list = append(list, fmt.Sprintf("http://www.kxdaili.com/dailiip/2/%d.html", i))
	}
	return list
}
func (s *getProxyKxDaili) GetContentHtml(requestUrl string) string {
	h := &request.RequestHeaderDto{
		UserAgent:               consts.USER_AGENT,
		UpgradeInsecureRequests: "1",
	}

	logger.Info("get proxy from kxdaili.com", logger.Fields{"url": requestUrl})
	data, err := request.WebGet(requestUrl, h, nil)
	if err != nil || data == nil {
		logger.Error("get proxy from kxdaili.com fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *getProxyKxDaili) ParseHtml(body string) [][]string {

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
