package model

//import (
//	"github.com/jinzhu/gorm"
//)
type Proxy struct {
	Id         int    `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	Host       string `gorm:"column:host"`
	Port       string
	Status     int8   `gorm:"column:status"`
	CreateTime int64  `gorm:"column:create_time"`
	UpdateTime int64  `gorm:"column:update_time"`
	ActiveTime int64  `gorm:"column:active_time"`
	Country    string `gorm:"column:country"`
	Region     string `gorm:"column:region"`
	City       string `gorm:"column:city"`
	Isp        string `gorm:"column:isp"`
}

func (Proxy) TableName() string {
	return "proxy"
}
