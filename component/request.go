package component

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"proxy-collect/component/logger"
	"proxy-collect/dto"
	"time"
)

func WebGetProxy(requestUrl string, header dto.RequestHeaderDto, host string, port string) string {
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req = addHeader(req, &header)
	proxyServer := fmt.Sprintf("http://%s:%s", host, port)
	proxyUrl, _ := url.Parse(proxyServer)
	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
		Timeout:   time.Second * 5,
	}
	return request(client, req)
}

func WebGet(requestUrl string, header dto.RequestHeaderDto) string {
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req = addHeader(req, &header)

	client := &http.Client{
		Timeout: time.Second * 5,
	}
	return request(client, req)
}

func WebRequest(req *http.Request) string {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	return request(client, req)
}

func WebRequestProxy(req *http.Request, host string, port string) string {
	proxyServer := fmt.Sprintf("http://%s:%s", host, port)
	proxyUrl, _ := url.Parse(proxyServer)
	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
		Timeout:   time.Second * 5,
	}
	return request(client, req)
}

func request(client *http.Client, req *http.Request) string {
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("http get error", err)
		return ""
	}
	if resp.StatusCode != 200 {
		resp.Body.Close()
		logger.Error("http status error ", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()
	data := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		data, err = gzip.NewReader(resp.Body)
		if err != nil {
			logger.Error("read gzip response error", err)
			return ""
		}
		defer data.Close()
	}
	body, err := ioutil.ReadAll(data)
	if err != nil {
		logger.Error("read error", err)
		return ""
	}
	return string(body)
}

func addHeader(req *http.Request, h *dto.RequestHeaderDto) *http.Request {
	if h.Host != "" {
		req.Header.Set("Host", h.Host)
	}
	if h.Accept != "" {
		req.Header.Set("Accept", h.Accept)
	}
	if h.AcceptEncoding != "" {
		req.Header.Set("Accept-Encoding", h.AcceptEncoding)
	}
	if h.Referer != "" {
		req.Header.Set("Referer", h.Referer)
	}
	if h.UpgradeInsecureRequests != "" {
		req.Header.Set("Upgrade-Insecure-Requests", h.UpgradeInsecureRequests)
	}
	if h.UserAgent != "" {
		req.Header.Set("User-Agent", h.UserAgent)
	}
	if h.AcceptLanguage != "" {
		req.Header.Set("Accept-Language", h.AcceptLanguage)
	}
	if h.SecFetchDest != "" {
		req.Header.Set("Sec-Fetch-Dest", h.SecFetchDest)
	}
	if h.SecFetchMode != "" {
		req.Header.Set("Sec-Fetch-Mode", h.SecFetchMode)
	}
	return req
}
