package proxy_getter

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"proxy-collect/consts"
	"proxy-collect/dao"
)

func NewGetProxyZdaye() *Zdaye {
	return &Zdaye{}
}

type Zdaye struct {
}

func (s *Zdaye) GetUrlList() []string {
	u := "https://www.zdaye.com/dayProxy/1.html"
	body := s.GetContentHtml(u)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(consts.GO_QUERY_READ_ERROR, logger.Fields{"err": err})
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

func (s *Zdaye) GetContentHtml(requestUrl string) string {

	h := &request.HeaderDto{
		UserAgent:               consts.USER_AGENT,
		UpgradeInsecureRequests: "1",
		Host:                    "www.zdaye.com",
		Accept:                  "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		AcceptEncoding:          "gzip, deflate, br",
		AcceptLanguage:          "zh-CN,zh;q=0.9",
		SecFetchDest:            "document",
		SecFetchMode:            "navigate",
	}

	logger.Info("get proxy from zdaye.com", logger.Fields{"url": requestUrl})
	proxies, err := dao.ProxyDao.GetActiveList()
	var data *request.HttpResultDto
	if err == nil && len(proxies) > 0 {
		rand.Seed(time.Now().Unix())
		i := rand.Intn(len(proxies))
		proxy := proxies[i]
		data, err = request.Get(requestUrl, request.NewOptions().WithHeader(h).WithProxy(&request.ProxyDto{Host: proxy.Host, Port: proxy.Port}))
	} else {
		data, err = request.Get(requestUrl, request.NewOptions().WithHeader(h))
	}
	if err != nil || data == nil {
		logger.Error("get proxy from zdaye.com fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *Zdaye) ParseHtml(body string) [][]string {

	var proxyList [][]string
	re := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+):(\d+)`)
	matched := re.FindAllStringSubmatch(body, -1)
	for _, match := range matched {
		proxyArr := []string{match[1], match[2]}
		proxyList = append(proxyList, proxyArr)
	}
	return proxyList
}
