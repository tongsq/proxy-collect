package ip

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"golang.org/x/text/encoding/simplifiedchinese"
	"proxy-collect/consts"
	"proxy-collect/dto"
)

func GetIpInfoByIp138(host string, port string) *dto.IpInfoDto {
	requestUrl := fmt.Sprintf("https://www.ip138.com/iplookup.asp?ip=%s&action=2", host)
	h := &request.HeaderDto{
		UserAgent:               consts.USER_AGENT,
		UpgradeInsecureRequests: "1",
		Host:                    "www.ip138.com",
		Referer:                 "https://www.ip138.com/",
		AcceptEncoding:          "gzip, deflate, br",
		Accept:                  "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	}

	logger.Info("get ip info from ip138", logger.Fields{"url": requestUrl})
	data, err := request.WebGetProxy(requestUrl, h, nil, &request.ProxyDto{Host: host, Port: port})
	if err != nil || data == nil {
		logger.Error("get from ip138 use proxy error", logger.Fields{"err": err, "data": data})
		data, err = request.WebGet(requestUrl, h, nil)
		if err != nil || data == nil {
			logger.Error("get from ip138 no proxy error", logger.Fields{"err": err, "data": data})
			return nil
		}
	}
	re := regexp.MustCompile(`var ip_result = (.+);`)
	matched := re.FindAllStringSubmatch(data.Body, -1)
	if len(matched) < 1 {
		return nil
	}
	jsonStr := matched[0][1]
	jsonStr, err = simplifiedchinese.GBK.NewDecoder().String(jsonStr)
	if err != nil {
		logger.FError("gb2313 decode error")
	}
	var result map[string][]map[string]string
	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		logger.Error("parse ip info fail", map[string]interface{}{"jsonStr": jsonStr})
	}
	logger.Info("get ip info result", logger.Fields{"jsonStr": jsonStr, "result": result})
	info := result["ip_c_list"][0]
	ipInfoDto := &dto.IpInfoDto{
		Country: info["ct"],
		Region:  info["prov"],
		City:    info["city"],
		Isp:     info["yunyin"],
	}
	if ipInfoDto.Country == "" {
		return nil
	}
	return ipInfoDto
}
