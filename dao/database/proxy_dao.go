package database

import (
	"time"

	"github.com/jinzhu/gorm"
	"proxy-collect/consts"
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

func (d *proxyDao) GetOne(host string, port string) (*model.ProxyModel, error) {
	var m model.ProxyModel
	db := DB()
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
		ActiveTime: time.Now().Unix(),
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		CheckCount: 1,
		Source:     source,
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

func (d *proxyDao) Delete(host string, port string) error {
	DB().Where("host=? and port=?", host, port).Delete(model.ProxyModel{})
	return nil
}
