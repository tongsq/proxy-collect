package get_proxy

import (
	"regexp"

	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"proxy-collect/consts"
)

func NewGetProxyZdayeIndex() *ZdayeIndex {
	return &ZdayeIndex{}
}

type ZdayeIndex struct {
}

func (s *ZdayeIndex) GetUrlList() []string {
	return []string{"https://www.zdaye.com/dayProxy/1.html"}
}

func (s *ZdayeIndex) GetContentHtml(requestUrl string) string {

	h := &request.RequestHeaderDto{
		UserAgent:               consts.USER_AGENT,
		UpgradeInsecureRequests: "1",
		Host:                    "www.zdaye.com",
		Accept:                  "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		AcceptEncoding:          "gzip, deflate, br",
		AcceptLanguage:          "zh-CN,zh;q=0.9",
		SecFetchDest:            "document",
		SecFetchMode:            "navigate",
	}

	logger.Info("get proxy from zdaye.com index", logger.Fields{"url": requestUrl})

	data, err := request.WebGet(requestUrl, h, nil)

	if err != nil || data == nil {
		logger.Error("get proxy from zdaye.com index fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *ZdayeIndex) ParseHtml(body string) [][]string {

	var proxyList [][]string
	re := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+):(\d+)`)
	matched := re.FindAllStringSubmatch(body, -1)
	for _, match := range matched {
		proxyArr := []string{match[1], match[2]}
		proxyList = append(proxyList, proxyArr)
	}
	return proxyList
}
