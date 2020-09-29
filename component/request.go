package component

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"proxy-collect/component/logger"
	"time"
)

func WebRequestProxy(req *http.Request, host string, port string) string {
	proxyServer := fmt.Sprintf("http://%s:%s", host, port)
	proxyUrl, _ := url.Parse(proxyServer)
	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
		Timeout:   time.Second * 5,
	}
	return request(client, req)
}

func WebRequest(req *http.Request) string {
	client := &http.Client{
		Timeout: time.Second * 10,
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return ""
	}
	return string(body)
}
