package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"golang.org/x/text/encoding/simplifiedchinese"
	"proxy-collect/consts"
	"strings"
)

func NewGetProxy66ip() *getProxy66ip {
	return &getProxy66ip{}
}

type getProxy66ip struct {
}

func (s *getProxy66ip) GetUrlList() []string {
	list := []string{
		"http://www.66ip.cn/index.html",
	}
	for i := 2; i < 6; i++ {
		list = append(list, fmt.Sprintf("http://www.66ip.cn/%d.html", i))
	}
	return list
}

func (s *getProxy66ip) GetContentHtml(requestUrl string) string {
	h := &request.RequestHeaderDto{
		UserAgent:               consts.USER_AGENT,
		Host:                    "www.66ip.cn",
		Referer:                 "http://www.66ip.cn/2.html",
		UpgradeInsecureRequests: "1",
	}
	logger.Info("get proxy from 66ip", logger.Fields{"url": requestUrl})
	data, err := request.WebGet(requestUrl, h, nil)
	if err != nil || data == nil {
		logger.Error("get proxy from 66ip fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *getProxy66ip) ParseHtml(body string) [][]string {
	body, err := simplifiedchinese.GBK.NewDecoder().String(body)
	if err != nil {
		logger.FError("chang charset to utf8 fail")
		return nil
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.FError(consts.GO_QUERY_READ_ERROR, logger.Fields{"err": err})
		return nil
	}
	var proxyList [][]string

	doc.Find("tr").Each(func(i int, selection *goquery.Selection) {
		td := selection.ChildrenFiltered("td").First()
		td2 := selection.ChildrenFiltered("td").Eq(1)
		proxyHost := td.Text()
		proxyPort := td2.Text()
		if !ProxyService.CheckProxyFormat(proxyHost, proxyPort) {
			return
		}
		proxyArr := []string{strings.TrimSpace(proxyHost), strings.TrimSpace(proxyPort)}
		proxyList = append(proxyList, proxyArr)
	})
	return proxyList
}
