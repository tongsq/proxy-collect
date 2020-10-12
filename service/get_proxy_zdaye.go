package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"proxy-collect/component"
	"proxy-collect/component/logger"
	"proxy-collect/config"
	"regexp"
	"strings"
)

func NewGetProxyZdaye() *getProxyZdaye {
	return &getProxyZdaye{}
}

type getProxyZdaye struct {
}

func (s *getProxyZdaye) GetUrlList() []string {
	u := "https://www.zdaye.com/dayProxy/1.html"
	body := s.GetContentHtml(u)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(err)
		return nil
	}
	list := []string{}
	doc.Find("p.thread_tags").Each(func(i int, selection *goquery.Selection) {
		a := selection.ChildrenFiltered("A").First()
		href, _ := a.Attr("href")
		list = append(list, fmt.Sprintf("https://www.zdaye.com%s", href))
	})
	if len(list) > 5 {
		return list[0:5]
	}
	return list
}
func (s *getProxyZdaye) GetContentHtml(requestUrl string) string {
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("user-agent", config.USER_AGENT)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Host", "www.zdaye.com")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")

	logger.Info("get proxy from zdaye.com", requestUrl)
	return component.WebRequest(req)
}

func (s *getProxyZdaye) ParseHtml(body string) [][]string {

	var proxyList [][]string
	re := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+):(\d+)`)
	matched := re.FindAllStringSubmatch(body, -1)
	for _, match := range matched {
		proxyArr := []string{match[1], match[2]}
		proxyList = append(proxyList, proxyArr)
	}
	return proxyList
}
