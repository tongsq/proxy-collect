package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"proxy-collect/component/logger"
	"proxy-collect/config"
	"strings"
	"time"
)

type GetProxyNima struct {
}

func (s *GetProxyNima) GetUrlList() []string {
	list := []string{
		"http://www.nimadaili.com/https/",
	}
	for i := 1; i < 6; i++ {
		list = append(list, fmt.Sprintf("http://www.nimadaili.com/https/%d/", i))
	}
	return list
}
func (s *GetProxyNima) GetContentHtml(requestUrl string) io.ReadCloser {
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("User-Agent", config.USER_AGENT)
	req.Header.Set("Host", "www.nimadaili.com")
	req.Header.Set("Referer", "http://www.nimadaili.com/https/3/")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	client := http.Client{
		Timeout: time.Second * 5,
	}
	logger.Info("get proxy from nimadaili", requestUrl)
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("http get error", err)
		return nil
	}
	if resp.StatusCode != 200 {
		resp.Body.Close()
		logger.Error("http status error ", resp.StatusCode)
		return nil
	}
	return resp.Body
}

func (s *GetProxyNima) ParseHtml(body io.ReadCloser) [][]string {
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		logger.Error(err)
		return nil
	}
	var proxy_list [][]string
	doc.Find("tbody > tr").Each(func(i int, selection *goquery.Selection) {
		td := selection.ChildrenFiltered("td").First()
		proxy_str := td.Text()
		proxy_str = strings.Trim(proxy_str, " ")
		proxy_arr := strings.Split(proxy_str, ":")
		if len(proxy_arr) != 2 {
			logger.Error("格式错误:", proxy_str)
			return
		}
		proxy_list = append(proxy_list, proxy_arr)
	})
	return proxy_list
}
