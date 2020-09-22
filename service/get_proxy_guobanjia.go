package service

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"proxy-collect/component"
	"proxy-collect/component/logger"
	"proxy-collect/config"
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

	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("User-Agent", config.USER_AGENT)
	req.Header.Set("Host", "www.goubanjia.com")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	logger.Info("get proxy from guobanjia", requestUrl)
	return component.WebRequest(req)
}

func (s *getProxyGuoBanJia) ParseHtml(body string) [][]string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(err)
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
		hostStr = strings.Trim(hostStr, " ")
		if !ProxyService.CheckProxyFormat(hostStr, port) {
			logger.Error("格式错误:", hostStr, port)
			return
		}
		proxyArr := []string{hostStr, port}
		proxyList = append(proxyList, proxyArr)
	})

	return proxyList
}
