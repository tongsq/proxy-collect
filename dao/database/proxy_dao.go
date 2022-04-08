package database

import (
	"time"

	"github.com/jinzhu/gorm"
	"proxy-collect/consts"
	"proxy-collect/dto"
	"proxy-collect/model"
)

func NewMysqlProxyDao() *proxyDao {
	return &proxyDao{}
}

type proxyDao struct {
}

func (d *proxyDao) GetFailList() ([]model.ProxyModel, error) {
	var proxies []model.ProxyModel
	DB().Where("status=?", consts.STATUS_NO).Where("check_count>0").Find(&proxies)
	return proxies, nil
}

func (d *proxyDao) GetRecheckList() ([]model.ProxyModel, error) {
	var proxies []model.ProxyModel
	DB().Where("status=?", consts.STATUS_RECHECK).Where("check_count>0").Find(&proxies)
	return proxies, nil
}

func (d *proxyDao) GetOne(host string, port string, proto string) (*model.ProxyModel, error) {
	var m model.ProxyModel
	db := DB()
	err := db.Where("host = ? AND port = ? AND proto = ?", host, port, proto).First(&m).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *proxyDao) Create(proxy dto.ProxyDto, status int8) (*model.ProxyModel, error) {
	m := &model.ProxyModel{
		Host:       proxy.Host,
		Port:       proxy.Port,
		Status:     status,
		ActiveTime: time.Now().Unix(),
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		CheckCount: 1,
		Source:     proxy.Source,
		Proto:      proxy.Proto,
		User:       proxy.User,
		Password:   proxy.Password,
	}
	DB().Create(m)
	return m, nil
}

func (d *proxyDao) Save(m *model.ProxyModel) error {
	DB().Save(m)
	return nil
}

func (d *proxyDao) GetActiveList() ([]model.ProxyModel, error) {
	var proxies []model.ProxyModel
	DB().Where("status=?", consts.STATUS_YES).Find(&proxies)
	return proxies, nil
}

func (d *proxyDao) GetNeedUpdateInfoList() []model.ProxyModel {
	var proxies []model.ProxyModel
	DB().Where("status=? and country=?", consts.STATUS_YES, "").Find(&proxies)
	return proxies
}

func (d *proxyDao) Delete(host string, port string, proto string) error {
	DB().Where("host=? and port=? and proto=?", host, port, proto).Delete(model.ProxyModel{})
	return nil
}
