package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"net/http"
	"proxy-collect/component"
	"proxy-collect/component/logger"
	"proxy-collect/config"
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

	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("User-Agent", config.USER_AGENT)
	req.Header.Set("Host", "www.66ip.cn")
	req.Header.Set("Referer", "http://www.66ip.cn/2.html")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	logger.Info("get proxy from 66ip", requestUrl)
	return component.WebRequest(req)
}

func (s *getProxy66ip) ParseHtml(body string) [][]string {
	body, err := simplifiedchinese.GBK.NewDecoder().String(body)
	if err != nil {
		logger.Error("chang charset to utf8 fail")
		return nil
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(err)
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
