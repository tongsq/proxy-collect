package proxy_getter

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"proxy-collect/consts"
)

func NewGetProxyYqie() *Yqie {
	return &Yqie{}
}

type Yqie struct {
}

func (s *Yqie) GetUrlList() []string {
	return []string{"http://ip.yqie.com/ipproxy.htm"}
}

func (s *Yqie) GetContentHtml(requestUrl string) string {

	h := &request.RequestHeaderDto{
		UserAgent:               consts.USER_AGENT,
		UpgradeInsecureRequests: "1",
		Host:                    "ip.yqie.com",
		Accept:                  "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		AcceptEncoding:          "gzip, deflate, br",
		AcceptLanguage:          "zh-CN,zh;q=0.9",
		SecFetchDest:            "document",
		SecFetchMode:            "navigate",
	}

	logger.Info("get proxy from ip.yqie.com", logger.Fields{"url": requestUrl})

	data, err := request.WebGet(requestUrl, h, nil)

	if err != nil || data == nil {
		logger.Error("get proxy from ip.yqie.com fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *Yqie) ParseHtml(body string) [][]string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(consts.GO_QUERY_READ_ERROR, logger.Fields{"err": err})
		return nil
	}
	var proxyList [][]string
	doc.Find("tbody > tr").Each(func(i int, selection *goquery.Selection) {
		td := selection.ChildrenFiltered("td").First()
		proxyHost := td.Text()
		td2 := selection.ChildrenFiltered("td").Eq(1)
		proxyPort := td2.Text()
		if proxyHost == "" || proxyPort == "" {
			logger.FError("parse html node fail")
			return
		}
		proxyArr := []string{strings.TrimSpace(proxyHost), strings.TrimSpace(proxyPort)}
		proxyList = append(proxyList, proxyArr)
	})
	return proxyList
}
