package proxy_getter

import (
	"github.com/tongsq/go-lib/request"
	"math/rand"
	"proxy-collect/config"
	"proxy-collect/dao"
	"regexp"
	"time"

	"github.com/tongsq/go-lib/logger"
)

func NewGetter(conf *config.Getter) *getter {
	return &getter{config: conf}
}

type getter struct {
	config *config.Getter
}

func (s *getter) GetUrlList() []string {
	return s.config.Urls
}
func (s *getter) GetContentHtml(requestUrl string) string {
	logger.Info("get proxy from config getter", logger.Fields{"url": requestUrl})
	h := &request.HeaderDto{
		UserAgent: s.config.Agent,
	}
	var data *request.HttpResultDto
	var err error
	options := request.NewOptions().WithHeader(h)
	if s.config.Proxy {
		proxies, err := dao.ProxyDao.GetActiveList()
		if err == nil && len(proxies) > 0 {
			rand.Seed(time.Now().Unix())
			i := rand.Intn(len(proxies))
			proxy := proxies[i]
			options = options.WithProxy(&request.ProxyDto{Host: proxy.Host, Port: proxy.Port})
		}
	}
	if s.config.Method == "GET" {
		data, err = request.Get(requestUrl, options)
	} else {
		data, err = request.Post(requestUrl, options)
	}
	if err != nil || data == nil {
		logger.Error("get proxy from config getter fail", logger.Fields{"err": err, "data": data, "url": requestUrl})
		return ""
	}
	return data.Body
}

func (s *getter) ParseHtml(body string) [][]string {
	var proxyList [][]string
	re := regexp.MustCompile(s.config.Regexp)
	matched := re.FindAllStringSubmatch(body, -1)
	for _, match := range matched {
		proxyArr := []string{match[1], match[2], s.config.Proto}
		proxyList = append(proxyList, proxyArr)
	}
	return proxyList
}
