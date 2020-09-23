package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"proxy-collect/component"
	"proxy-collect/component/logger"
	"proxy-collect/config"
	"strings"
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
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("User-Agent", config.USER_AGENT)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Host", "ip.jiangxianli.com")
	req.Header.Set("Referer", "https://ip.jiangxianli.com/")
	logger.Info("get proxy from jangxianli", requestUrl)
	return component.WebRequest(req)
}

func (s *getProxyIpJiangXianLi) ParseHtml(body string) [][]string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(err)
		return nil
	}
	var proxyList [][]string
	doc.Find("tbody > tr").Each(func(i int, selection *goquery.Selection) {
		td := selection.ChildrenFiltered("td").First()
		host := strings.Trim(td.Text(), " ")
		td2 := selection.ChildrenFiltered("td").Eq(1)
		port := strings.Trim(td2.Text(), " ")

		if !ProxyService.CheckProxyFormat(host, port) {
			logger.Error("格式错误:", host, port)
			return
		}
		proxyArr := []string{host, port}
		proxyList = append(proxyList, proxyArr)
	})
	return proxyList
}
