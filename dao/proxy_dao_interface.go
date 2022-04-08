package dao

import (
	"proxy-collect/dto"
	"proxy-collect/model"
)

type proxyDaoInterface interface {
	GetFailList() ([]model.ProxyModel, error)
	GetOne(host string, port string, proto string) (*model.ProxyModel, error)
	Create(proxy dto.ProxyDto, status int8) (*model.ProxyModel, error)
	Save(m *model.ProxyModel) error
	GetActiveList() ([]model.ProxyModel, error)
	GetRecheckList() ([]model.ProxyModel, error)
	GetNeedUpdateInfoList() []model.ProxyModel
	Delete(host string, port string, proto string) error
}
