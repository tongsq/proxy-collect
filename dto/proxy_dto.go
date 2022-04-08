package dto

import (
	"proxy-collect/model"
)

type ProxyInfoDto struct {
	ProxyDto
	Status     int8   `json:"status"`
	CreateTime int64  `json:"create_time"`
	ActiveTime int64  `json:"active_time"`
	Country    string `json:"country"`
	Region     string `json:"region"`
	City       string `json:"city"`
	Isp        string `json:"isp"`
}

type ProxyDto struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Source   string `json:"source"`
	Proto    string `json:"proto"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func NewProxyDto(m model.ProxyModel) ProxyInfoDto {
	return ProxyInfoDto{
		ProxyDto: ProxyDto{
			Host:     m.Host,
			Port:     m.Port,
			Proto:    m.Proto,
			User:     m.User,
			Password: m.Password,
			Source:   m.Source,
		},
		Status:     m.Status,
		CreateTime: m.CreateTime,
		ActiveTime: m.ActiveTime,
		Country:    m.Country,
		Region:     m.Region,
		City:       m.City,
		Isp:        m.Isp,
	}
}
