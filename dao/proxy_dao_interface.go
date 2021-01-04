package dao

import "proxy-collect/model"

type proxyDaoInterface interface {
	GetFailList() []model.ProxyModel
	GetOne(host string, port string) (*model.ProxyModel, error)
	Create(host string, port string, status int8, source string) (*model.ProxyModel, error)
	Save(m *model.ProxyModel) error
	GetActiveList() []model.ProxyModel
	GetNeedUpdateInfoList() []model.ProxyModel
}
