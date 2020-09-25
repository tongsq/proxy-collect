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

func NewGetProxyProxyList() *getProxyProxyList {
	return &getProxyProxyList{}
}

type getProxyProxyList struct {
}

func (s *getProxyProxyList) GetUrlList() []string {
	list := []string{
		"https://list.proxylistplus.com/Fresh-HTTP-Proxy-List-1",
	}
	for i := 2; i < 6; i++ {
		list = append(list, fmt.Sprintf("https://list.proxylistplus.com/Fresh-HTTP-Proxy-List-%d", i))
	}
	return list
}
func (s *getProxyProxyList) GetContentHtml(requestUrl string) string {
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("user-agent", config.USER_AGENT)
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("referer", "https://list.proxylistplus.com/update-2")
	logger.Info("get proxy from list.proxylistplus.com", requestUrl)
	return component.WebRequest(req)
}

func (s *getProxyProxyList) ParseHtml(body string) [][]string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(err)
		return nil
	}

	var proxyList [][]string
	doc.Find("tr").Each(func(i int, selection *goquery.Selection) {
		td := selection.ChildrenFiltered("td").Eq(1)
		host := strings.TrimSpace(td.Text())
		td2 := selection.ChildrenFiltered("td").Eq(2)
		port := strings.TrimSpace(td2.Text())

		if !ProxyService.CheckProxyFormat(host, port) {
			return
		}
		proxyArr := []string{host, port}
		proxyList = append(proxyList, proxyArr)
	})
	return proxyList
}
