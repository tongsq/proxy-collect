package dto

import "proxy-collect/model"

type ProxyDto struct {
	Host       string `json:"host"`
	Port       string `json:"port"`
	Status     int8   `json:"status"`
	CreateTime int64  `json:"create_time"`
	ActiveTime int64  `json:"active_time"`
	Country    string `json:"country"`
	Region     string `json:"region"`
	City       string `json:"city"`
	Isp        string `json:"isp"`
	Source     string `json:"source"`
}

func NewProxyDto(m model.ProxyModel) ProxyDto {
	return ProxyDto{
		Host:       m.Host,
		Port:       m.Port,
		Status:     m.Status,
		CreateTime: m.CreateTime,
		ActiveTime: m.ActiveTime,
		Country:    m.Country,
		Region:     m.Region,
		City:       m.City,
		Isp:        m.Isp,
		Source:     m.Source,
	}
}
