package service

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net/http"
	"net/url"
	"proxy-collect/component"
	"proxy-collect/component/logger"
	"proxy-collect/config"
	"proxy-collect/dto"
	"proxy-collect/model"
	"reflect"
	"regexp"
	"sync"
	"time"
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
		logger.Error("http get error", err)
		return false
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("http read error", err)
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
		logger.Warning("read error", err)
		return false
	}
	return true
}

func (s *proxyService) CheckProxyAndSave(host string, port string, source string) {
	result := s.CheckIpStatus(host, port)
	if result {
		logger.Success(result, host, port)
	} else {
		logger.Warning(result, host, port)
	}
	var status int8 = 1
	if !result {
		status = 0
	}
	var proxyModel model.Proxy
	db := model.DB
	err := db.Where("host = ? AND port = ?", host, port).First(&proxyModel).Error

	if err != nil && gorm.IsRecordNotFoundError(err) {
		if status == 0 {
			return
		}
		proxyModel = model.Proxy{
			Host:       host,
			Port:       port,
			Status:     status,
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
			CheckCount: 1,
			Source:     source,
		}
		db.Create(&proxyModel)
		return
	}
	if status == 1 {
		if proxyModel.CheckCount <= 20 {
			proxyModel.CheckCount = proxyModel.CheckCount + 1
		}
		proxyModel.ActiveTime = time.Now().Unix()
	} else if status == 0 && proxyModel.CheckCount > -10 {
		proxyModel.CheckCount = proxyModel.CheckCount - 1
	}
	if source != "" {
		proxyModel.Source = source
	}
	proxyModel.Status = status
	proxyModel.UpdateTime = time.Now().Unix()
	db.Save(&proxyModel)
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
		logger.Info("获取到ip:", proxyList)
		var wg sync.WaitGroup = sync.WaitGroup{}
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			logger.Info("wait 10s ...")
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

func (s *proxyService) UpdateIpDetail(m *model.Proxy) {

}

func (s *proxyService) GetIpInfo(host string, port string) *dto.IpInfoDto {
	requestUrl := fmt.Sprintf("https://www.ip138.com/iplookup.asp?ip=%s&action=2", host)
	h := dto.RequestHeaderDto{
		UserAgent:               config.USER_AGENT,
		UpgradeInsecureRequests: "1",
		Host:                    "www.ip138.com",
		Referer:                 "https://www.ip138.com/",
		AcceptEncoding:          "gzip, deflate, br",
		Accept:                  "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	}

	logger.Info("get ip info from ip138", requestUrl)
	body := component.WebGetProxy(requestUrl, h, host, port)
	if body == "" {
		logger.Error("get from ip138 use proxy error")
		body = component.WebGet(requestUrl, h)
		if body == "" {
			logger.Error("get from ip138 no proxy error")
			return nil
		}
	}
	fmt.Println(body)
	re := regexp.MustCompile(`var ip_result = (.+);`)
	matched := re.FindAllStringSubmatch(body, -1)
	if len(matched) < 1 {
		return nil
	}
	jsonStr := matched[0][1]
	jsonStr, err := simplifiedchinese.GBK.NewDecoder().String(jsonStr)
	if err != nil {
		logger.Error("gb2313 decode error")
	}
	var data map[string][]map[string]string
	err = json.Unmarshal([]byte(jsonStr), &data)

	info := data["ip_c_list"][0]
	ipInfoDto := &dto.IpInfoDto{}
	ipInfoDto.Country = info["ct"]
	ipInfoDto.Region = info["prov"]
	ipInfoDto.City = info["city"]
	ipInfoDto.Isp = info["yunyin"]
	if ipInfoDto.Country == "" {
		return nil
	}
	return ipInfoDto
}
