package redis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tongsq/go-lib/logger"
	redis_client "github.com/tongsq/go-lib/redis-client"
	"proxy-collect/consts"
	"proxy-collect/dto"
	"proxy-collect/model"
)

const PROXY_FAIL_SET = "proxy_fail_set"
const PROXY_INFO_MAP = "proxy_info_map"
const PROXY_SUCCESS_SET = "proxy_success_set"
const PROXY_RECHECK_SET = "proxy_recheck_set"

func NewRedisProxyDao() *proxyDao {
	return &proxyDao{}
}

type proxyDao struct {
}

func (d *proxyDao) GetFailList() ([]model.ProxyModel, error) {
	return d.getProxyList(PROXY_FAIL_SET)
}

func (d *proxyDao) GetRecheckList() ([]model.ProxyModel, error) {
	return d.getProxyList(PROXY_RECHECK_SET)
}

func (d *proxyDao) GetOne(host string, port string, proto string) (*model.ProxyModel, error) {
	var m model.ProxyModel
	info, err := Client().HMGetOne(PROXY_INFO_MAP, getProxyKey(host, port, proto))
	if info == "" || err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(info), &m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (d *proxyDao) Create(proxy dto.ProxyDto, status int8) (*model.ProxyModel, error) {
	m := &model.ProxyModel{
		Host:       proxy.Host,
		Port:       proxy.Port,
		Status:     status,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		ActiveTime: time.Now().Unix(),
		CheckCount: 1,
		Source:     proxy.Source,
		Proto:      proxy.Proto,
		User:       proxy.User,
		Password:   proxy.Password,
	}
	key := getProxyKey(proxy.Host, proxy.Port, proxy.Proto)
	logger.FInfo("redis dao create new proxy")
	value, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	result, err := Client().HMSet(PROXY_INFO_MAP, redis_client.HMDto{Field: key, Value: string(value)})
	if !result || err != nil {
		return nil, err
	}
	updateProxySet(m)
	return m, err
}

func (d *proxyDao) Save(m *model.ProxyModel) error {
	key := getProxyKey(m.Host, m.Port, m.Proto)
	value, err := json.Marshal(m)
	if err != nil {
		return err
	}
	result, err := Client().HMSet(PROXY_INFO_MAP, redis_client.HMDto{Field: key, Value: string(value)})
	if !result || err != nil {
		return err
	}
	updateProxySet(m)
	return nil
}

func (d *proxyDao) GetActiveList() ([]model.ProxyModel, error) {
	return d.getProxyList(PROXY_SUCCESS_SET)
}

func (d *proxyDao) GetNeedUpdateInfoList() []model.ProxyModel {
	proxies, err := d.GetActiveList()
	if err != nil || proxies == nil {
		return nil
	}
	var needUpdateList []model.ProxyModel
	for _, proxy := range proxies {
		if proxy.Country == "" {
			needUpdateList = append(needUpdateList, proxy)
		}
	}
	return needUpdateList
}

func (d *proxyDao) Delete(host string, port string, proto string) error {
	key := getProxyKey(host, port, proto)
	_, err := Client().HDel(PROXY_INFO_MAP, key)
	if err != nil {
		return err
	}
	deleteProxySet(host, port, proto)
	return nil
}

func ConvertInterface(args ...string) []interface{} {
	var result []interface{}
	for _, v := range args {
		result = append(result, v)
	}
	return result
}

func (d *proxyDao) getProxyList(key string) ([]model.ProxyModel, error) {
	proxys, err := Client().SMembers(key)
	if err != nil || proxys == nil {
		return nil, err
	}
	infoList, err := Client().HMGet(PROXY_INFO_MAP, ConvertInterface(proxys...)...)
	if err != nil || infoList == nil {
		return nil, err
	}
	var models []model.ProxyModel
	for _, v := range infoList {
		m := model.ProxyModel{}
		jsonErr := json.Unmarshal([]byte(v), &m)
		if jsonErr != nil {
			continue
		}
		models = append(models, m)
	}
	return models, nil
}

func getProxyKey(host string, port string, proto string) string {
	if proto == "" || proto == consts.PROTO_HTTP {
		return fmt.Sprintf("%s:%s", host, port)
	} else {
		return fmt.Sprintf("%s:%s:%s", host, port, proto)
	}
}

func updateProxySet(m *model.ProxyModel) {
	if m == nil {
		return
	}
	key := getProxyKey(m.Host, m.Port, m.Proto)
	if m.Status == consts.STATUS_YES {
		_, _ = Client().SAdd(PROXY_SUCCESS_SET, key)
		_, _ = Client().SRem(PROXY_FAIL_SET, key)
		_, _ = Client().SRem(PROXY_RECHECK_SET, key)
	} else if m.Status == consts.STATUS_RECHECK {
		_, _ = Client().SAdd(PROXY_RECHECK_SET, key)
		_, _ = Client().SRem(PROXY_SUCCESS_SET, key)
		_, _ = Client().SRem(PROXY_FAIL_SET, key)
	} else {
		_, _ = Client().SAdd(PROXY_FAIL_SET, key)
		_, _ = Client().SRem(PROXY_SUCCESS_SET, key)
		_, _ = Client().SRem(PROXY_RECHECK_SET, key)
	}
}

func deleteProxySet(host string, port string, proto string) {
	key := getProxyKey(host, port, proto)
	Client().SRem(PROXY_SUCCESS_SET, key)
	Client().SRem(PROXY_FAIL_SET, key)
	Client().SRem(PROXY_RECHECK_SET, key)
}
