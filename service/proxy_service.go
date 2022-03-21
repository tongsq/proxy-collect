package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"sync"
	"time"

	"github.com/tongsq/go-lib/component"
	"github.com/tongsq/go-lib/logger"
	"proxy-collect/dao"
	"proxy-collect/model"
	"proxy-collect/service/ip"
)

func NewProxyService() *proxyService {
	return &proxyService{}
}

type proxyService struct {
}

func (s *proxyService) CheckIpStatusActive(host, port string) bool {
	request_url := "https://www.baidu.com"
	req, _ := http.NewRequest("GET", request_url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36")
	proxyServer := fmt.Sprintf("http://%s:%s", host, port)
	proxyUrl, _ := url.Parse(proxyServer)
	client := http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
		Timeout:   time.Second * 5,
	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("http get error", logger.Fields{"err": err})
		return false
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("http read error", logger.Fields{"err": err})
		return false
	}
	return true
}

func (s *proxyService) CheckIpStatus(host, port string) bool {
	request_url := "https://www.baidu.com"
	req, _ := http.NewRequest("GET", request_url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36")
	proxyServer := fmt.Sprintf("http://%s:%s", host, port)
	proxyUrl, _ := url.Parse(proxyServer)
	client := http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
		Timeout:   time.Second * 5,
	}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Warning("read error", logger.Fields{"err": err})
		return false
	}
	return true
}

func (s *proxyService) CheckProxyAndSave(host string, port string, source string) {
	result := s.CheckIpStatus(host, port)
	if result {
		logger.Success("ip is success", logger.Fields{"host": host, "port": port})
	} else {
		logger.Warning("ip is fail", logger.Fields{"host": host, "port": port})
	}
	var status int8 = 1
	if !result {
		status = 0
	}

	proxyModel, err := dao.ProxyDao.GetOne(host, port)
	if err != nil {
		logger.Error("get model fail:%s", logger.Fields{"err": err})
		return
	}
	if proxyModel == nil {
		if status == 0 {
			return
		}
		_, err = dao.ProxyDao.Create(host, port, status, source)
		return
	}
	if status == 1 {
		if proxyModel.CheckCount <= 20 {
			proxyModel.CheckCount = proxyModel.CheckCount + 1
		}
		proxyModel.ActiveTime = time.Now().Unix()
	} else if status == 0 && proxyModel.CheckCount > -10 {
		proxyModel.CheckCount = proxyModel.CheckCount - 1
		if proxyModel.CheckCount <= -10 {
			if err := dao.ProxyDao.Delete(host, port); err != nil {
				logger.Error("delete proxy fail", logger.Fields{"host": host, "port": port})
			}
			return
		}
	}
	if source != "" {
		proxyModel.Source = source
	}
	//if ip is ok, update ip info
	if result && proxyModel.City == "" {
		ipInfo := ip.GetIpInfo(host, port)
		logger.Success("get ip info", map[string]interface{}{"result": ipInfo})
		proxyModel.City = ipInfo.City
		proxyModel.Country = ipInfo.Country
		proxyModel.Isp = ipInfo.Isp
		proxyModel.Region = ipInfo.Region
	}
	proxyModel.Status = status
	proxyModel.UpdateTime = time.Now().Unix()
	err = dao.ProxyDao.Save(proxyModel)
	return
}

func (s *proxyService) DoGetProxy(getProxyService GetProxyInterface, pool *component.Pool) {
	for _, requestUrl := range getProxyService.GetUrlList() {
		contentBody := getProxyService.GetContentHtml(requestUrl)
		if contentBody == "" {
			time.Sleep(time.Second * 5)
			continue
		}
		proxyList := getProxyService.ParseHtml(contentBody)
		logger.Info("get ip list:", logger.Fields{"list": proxyList})
		var wg sync.WaitGroup = sync.WaitGroup{}
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			logger.FInfo("wait 10s ...")
			time.Sleep(time.Second * 10)
		}(&wg)
		for _, proxyArr := range proxyList {
			ip, port := proxyArr[0], proxyArr[1]
			source := reflect.TypeOf(getProxyService).String()[9:]
			pool.RunTask(func() { s.CheckProxyAndSave(ip, port, source) })
		}

		wg.Wait()
	}
}

func (s *proxyService) CheckProxyFormat(host string, port string) bool {
	ok, _ := regexp.Match(`^[\d\.]+$`, []byte(host))
	if !ok {
		return false
	}
	ok, _ = regexp.Match(`^\d+$`, []byte(port))
	if !ok {
		return false
	}
	return true
}

func (s *proxyService) UpdateIpDetail(m *model.ProxyModel) {

}
