package mysql

import (
	"github.com/jinzhu/gorm"
	"proxy-collect/model"
	"time"
)

func NewMysqlProxyDao() *proxyDao {
	return &proxyDao{}
}

type proxyDao struct {
}

func (d *proxyDao) GetFailList() []model.ProxyModel {
	var proxies []model.ProxyModel
	model.DB.Where("status<>?", 1).Where("check_count>0").Find(&proxies)
	return proxies
}

func (d *proxyDao) GetOne(host string, port string) (*model.ProxyModel, error) {
	var m model.ProxyModel
	db := model.DB
	err := db.Where("host = ? AND port = ?", host, port).First(&m).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *proxyDao) Create(host string, port string, status int8, source string) (*model.ProxyModel, error) {
	m := &model.ProxyModel{
		Host:       host,
		Port:       port,
		Status:     status,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		CheckCount: 1,
		Source:     source,
	}
	model.DB.Create(m)
	return m, nil
}

func (d *proxyDao) Save(m *model.ProxyModel) error {
	model.DB.Save(m)
	return nil
}

func (d *proxyDao) GetActiveList() []model.ProxyModel {
	var proxies []model.ProxyModel
	model.DB.Where("status=?", 1).Find(&proxies)
	return proxies
}

func (d *proxyDao) GetNeedUpdateInfoList() []model.ProxyModel {
	var proxies []model.ProxyModel
	model.DB.Where("status=? and country=?", 1, "").Find(&proxies)
	return proxies
}
