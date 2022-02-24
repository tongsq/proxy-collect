package get_proxy

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"proxy-collect/consts"
)

func NewGetProxyData5u() *getProxyData5u {
	return &getProxyData5u{}
}

type getProxyData5u struct {
}

func (s *getProxyData5u) GetUrlList() []string {
	list := []string{
		"http://www.data5u.com/",
		"http://www.data5u.com/free/gngn/index.shtml",
		"http://www.data5u.com/free/gnpt/index.shtml",
	}
	return list
}

func (s *getProxyData5u) GetContentHtml(requestUrl string) string {
	h := &request.RequestHeaderDto{
		UserAgent:               consts.USER_AGENT,
		Host:                    "www.data5u.com",
		UpgradeInsecureRequests: "1",
	}

	logger.Info("get proxy from data5u", logger.Fields{"url": requestUrl})
	data, err := request.WebGet(requestUrl, h, nil)
	if err != nil || data == nil {
		logger.Error("get proxy from data5u fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *getProxyData5u) ParseHtml(body string) [][]string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(consts.GO_QUERY_READ_ERROR, logger.Fields{"err": err})
		return nil
	}

	var proxyList [][]string
	doc.Find("ul.l2").Each(func(i int, selection *goquery.Selection) {
		td := selection.Find("span>li").First()
		proxyHost := td.Text()
		td2 := selection.Find("span>li").Eq(1)
		proxyPort := td2.Text()
		if proxyHost == "" || proxyPort == "" {
			logger.FError("parse html node fail")
		}
		proxyArr := []string{strings.TrimSpace(proxyHost), strings.TrimSpace(proxyPort)}
		proxyList = append(proxyList, proxyArr)
	})

	return proxyList
}
