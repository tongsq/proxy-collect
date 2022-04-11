package proxy_getter

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"proxy-collect/consts"
)

func NewGetProxyGeonode() *geonode {
	return &geonode{}
}

type geonode struct {
}

func (s *geonode) GetUrlList() []string {
	list := []string{
		"https://proxylist.geonode.com/api/proxy-list?limit=50&page=1&sort_by=lastChecked&sort_type=desc",
	}
	for i := 2; i < 4; i++ {
		list = append(list, fmt.Sprintf("https://proxylist.geonode.com/api/proxy-list?limit=50&page=%d&sort_by=lastChecked&sort_type=desc", i))
	}
	return list
}
func (s *geonode) GetContentHtml(requestUrl string) string {
	h := &request.HeaderDto{
		UserAgent: consts.USER_AGENT,
	}

	logger.Info("get proxy from geonode.com", logger.Fields{"url": requestUrl})
	data, err := request.WebGet(requestUrl, h, nil)
	if err != nil || data == nil {
		logger.Error("get proxy from geonode.com fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *geonode) ParseHtml(body string) [][]string {

	result := geonodeResult{}
	err := json.Unmarshal([]byte(body), &result)
	if err != nil {
		logger.Error("json parse fail", logger.Fields{"err": err})
		return nil
	}
	var proxyList [][]string
	for _, item := range result.Data {
		proto := ""
		for _, protoConst := range consts.PROTO_LIST {
			for _, hasProto := range item.Protocols {
				if protoConst == hasProto {
					proto = protoConst
					break
				}
			}
			if proto != "" {
				break
			}
		}
		if proto != "" {
			proxyList = append(proxyList, []string{item.IP, item.Port, proto})
		}
	}
	return proxyList
}

type geonodeResult struct {
	Data []struct {
		ID                 string      `json:"_id"`
		IP                 string      `json:"ip"`
		Port               string      `json:"port"`
		AnonymityLevel     string      `json:"anonymityLevel"`
		Asn                string      `json:"asn"`
		City               string      `json:"city"`
		Country            string      `json:"country"`
		CreatedAt          time.Time   `json:"created_at"`
		Google             bool        `json:"google"`
		Isp                string      `json:"isp"`
		LastChecked        int         `json:"lastChecked"`
		Latency            float64     `json:"latency"`
		Org                string      `json:"org"`
		Protocols          []string    `json:"protocols"`
		Region             interface{} `json:"region"`
		ResponseTime       int         `json:"responseTime"`
		Speed              int         `json:"speed"`
		UpdatedAt          time.Time   `json:"updated_at"`
		WorkingPercent     interface{} `json:"workingPercent"`
		UpTime             int         `json:"upTime"`
		UpTimeSuccessCount int         `json:"upTimeSuccessCount"`
		UpTimeTryCount     int         `json:"upTimeTryCount"`
		HostName           interface{} `json:"hostName,omitempty"`
	} `json:"data"`
	Total int    `json:"total"`
	Page  string `json:"page"`
	Limit string `json:"limit"`
}
