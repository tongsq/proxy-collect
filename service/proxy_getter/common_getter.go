package proxy_getter

import (
	"regexp"

	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"proxy-collect/consts"
)

var urlList = map[string][]string{
	consts.PROTO_SOCKS5: {
		"https://github.com/monosans/proxy-list/blob/main/proxies/socks5.txt",
	},
	consts.PROTO_SOCKS4: {
		"https://github.com/monosans/proxy-list/blob/main/proxies/socks4.txt",
	},
	consts.PROTO_HTTP: {
		"https://github.com/monosans/proxy-list/blob/main/proxies/http.txt",
	},
}

func NewCommonGetter(proto string) *commonGetter {
	return &commonGetter{Proto: proto}
}

type commonGetter struct {
	Proto string
}

func (s *commonGetter) GetUrlList() []string {
	if list, ok := urlList[s.Proto]; ok {
		return list
	}
	return nil
}
func (s *commonGetter) GetContentHtml(requestUrl string) string {
	h := &request.HeaderDto{
		UserAgent: consts.USER_AGENT,
	}
	logger.Info("get proxy from common getter", logger.Fields{"url": requestUrl})
	data, err := request.WebGet(requestUrl, h, nil)
	if err != nil || data == nil {
		logger.Error("get proxy from common getter fail", logger.Fields{"err": err, "data": data, "url": requestUrl})
		return ""
	}
	return data.Body
}

func (s *commonGetter) ParseHtml(body string) [][]string {
	var proxyList [][]string
	re := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+):(\d+)`)
	matched := re.FindAllStringSubmatch(body, -1)
	for _, match := range matched {
		proxyArr := []string{match[1], match[2], s.Proto}
		proxyList = append(proxyList, proxyArr)
	}
	return proxyList
}
