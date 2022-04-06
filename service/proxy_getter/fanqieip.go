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

func NewGetProxyFanQie() *getProxyFanQie {
	return &getProxyFanQie{}
}

type getProxyFanQie struct {
}

func (s *getProxyFanQie) GetUrlList() []string {
	list := []string{
		"https://www.fanqieip.com/free",
	}
	for i := 2; i < 6; i++ {
		list = append(list, fmt.Sprintf("https://www.fanqieip.com/free/%d", i))
	}
	return list
}

func (s *getProxyFanQie) GetContentHtml(requestUrl string) string {
	h := &request.RequestHeaderDto{
		UserAgent:               consts.USER_AGENT,
		Host:                    "www.fanqieip.com",
		UpgradeInsecureRequests: "1",
		Referer:                 "https://www.fanqieip.com/free",
	}

	logger.Info("get proxy from fanqieip", logger.Fields{"url": requestUrl})
	data, err := request.WebGet(requestUrl, h, nil)
	if err != nil || data == nil {
		logger.Error("get proxy from fanqieip fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *getProxyFanQie) ParseHtml(body string) [][]string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(consts.GO_QUERY_READ_ERROR, logger.Fields{"err": err})
		return nil
	}

	var proxyList [][]string
	doc.Find("tbody > tr").Each(func(i int, selection *goquery.Selection) {
		td := selection.ChildrenFiltered("td").First()
		host := strings.TrimSpace(td.ChildrenFiltered("div").First().Text())
		td2 := selection.ChildrenFiltered("td").Eq(1)
		port := strings.TrimSpace(td2.ChildrenFiltered("div").First().Text())

		if !common.CheckProxyFormat(host, port) {
			logger.Error(consts.PROXY_FORMAT_ERROR, logger.Fields{"host": host, "port": port})
			return
		}
		proxyArr := []string{host, port}
		proxyList = append(proxyList, proxyArr)
	})
	return proxyList
}
