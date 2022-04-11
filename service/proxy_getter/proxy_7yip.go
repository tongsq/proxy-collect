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

func NewGetProxy7Yip() *getProxy7Yip {
	return &getProxy7Yip{}
}

type getProxy7Yip struct {
}

func (s *getProxy7Yip) GetUrlList() []string {
	list := []string{
		"https://www.7yip.cn/free/",
	}
	for i := 2; i < 6; i++ {
		list = append(list, fmt.Sprintf("https://www.7yip.cn/free/?action=china&page=%d", i))
	}
	return list
}
func (s *getProxy7Yip) GetContentHtml(requestUrl string) string {
	h := &request.HeaderDto{
		UserAgent:               consts.USER_AGENT,
		UpgradeInsecureRequests: "1",
		Host:                    "www.7yip.cn",
		Referer:                 "https://www.7yip.cn/",
	}
	logger.Info("get proxy from 7yip", logger.Fields{"url": requestUrl})
	data, err := request.WebGet(requestUrl, h, nil)
	if err != nil || data == nil {
		logger.Error("get proxy from 7yip fail", logger.Fields{"err": err, "data": data})
	}
	return data.Body
}

func (s *getProxy7Yip) ParseHtml(body string) [][]string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(consts.GO_QUERY_READ_ERROR, logger.Fields{"err": err})
		return nil
	}
	var proxyList [][]string
	doc.Find("tbody > tr").Each(func(i int, selection *goquery.Selection) {
		td := selection.ChildrenFiltered("td").First()
		host := strings.TrimSpace(td.Text())
		td2 := selection.ChildrenFiltered("td").Eq(1)
		port := strings.TrimSpace(td2.Text())

		if !common.CheckProxyFormat(host, port) {
			logger.Error(consts.PROXY_FORMAT_ERROR, logger.Fields{"host": host, "port": port})
			return
		}
		proxyArr := []string{host, port}
		proxyList = append(proxyList, proxyArr)
	})

	return proxyList
}

//func (s *getProxy7Yip) GetSource() string {
//	_, file, _, _ := runtime.Caller(0)
//	arr := strings.Split(file, "/")
//	name := arr[len(arr)-1]
//	return name[0:len(name)-3]
//}
