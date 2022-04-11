package proxy_getter

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"proxy-collect/consts"
)

func NewGetProxyXsdaili() *Xsdaili {
	return &Xsdaili{}
}

type Xsdaili struct {
}

func (s *Xsdaili) GetUrlList() []string {
	u := "http://www.xsdaili.cn/"
	body := s.GetContentHtml(u)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(consts.GO_QUERY_READ_ERROR, logger.Fields{"err": err})
		return nil
	}
	list := []string{}
	doc.Find("div.title").Each(func(i int, selection *goquery.Selection) {
		a := selection.ChildrenFiltered("A").First()
		href, _ := a.Attr("href")
		list = append(list, fmt.Sprintf("http://www.xsdaili.cn%s", href))
	})
	if len(list) > 5 {
		return list[0:5]
	}
	return list
}

func (s *Xsdaili) GetContentHtml(requestUrl string) string {

	h := &request.HeaderDto{
		UserAgent: consts.USER_AGENT,
		Host:      "www.xsdaili.cn",
	}

	logger.Info("get proxy from xsdaili.com", logger.Fields{"url": requestUrl})
	data, err := request.WebGet(requestUrl, h, nil)
	if err != nil || data == nil {
		logger.Error("get proxy from zdaye.com fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *Xsdaili) ParseHtml(body string) [][]string {

	var proxyList [][]string
	re := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+):(\d+)`)
	matched := re.FindAllStringSubmatch(body, -1)
	for _, match := range matched {
		proxyArr := []string{match[1], match[2]}
		proxyList = append(proxyList, proxyArr)
	}
	return proxyList
}
