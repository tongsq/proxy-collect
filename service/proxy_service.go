package service

import (
	"fmt"
	"reflect"
	"regexp"
	"sync"
	"time"

	"github.com/tongsq/go-lib/component"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/go-lib/request"
	"proxy-collect/config"
	"proxy-collect/consts"
	"proxy-collect/dao"
	"proxy-collect/dto"
	"proxy-collect/service/ip"
)

func NewProxyService() *proxyService {
	return &proxyService{}
}

type proxyService struct {
}

func (s proxyService) TransferProxyDto(proxy *dto.ProxyDto) *request.ProxyDto {
	return &request.ProxyDto{
		Host:     proxy.Host,
		Port:     proxy.Port,
		Proto:    proxy.Proto,
		User:     proxy.User,
		Password: proxy.Password,
	}
}

func (s *proxyService) CheckIpStatus(proxy *dto.ProxyDto) bool {
	u := "https://www.baidu.com"
	h := &request.HeaderDto{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36",
	}
	_, err := request.WebGetProxy(u, h, nil, s.TransferProxyDto(proxy))
	return err == nil
}

func (s *proxyService) CheckProxyAndSave(p dto.ProxyDto) {
	result := s.CheckIpStatus(&p)
	if result {
		logger.Success("ip is success", logger.Fields{"proxy": p})
	} else {
		logger.Debug("ip is fail", logger.Fields{"proxy": p})
	}
	var status int8 = consts.STATUS_YES
	if !result {
		status = consts.STATUS_NO
	}

	proxyModel, err := dao.ProxyDao.GetOne(p.Host, p.Port, p.Proto)
	if err != nil {
		logger.Error("get model fail:%s", logger.Fields{"err": err})
		return
	}
	if proxyModel == nil {
		if status == consts.STATUS_NO {
			return
		}
		_, err = dao.ProxyDao.Create(p, status)
		if err != nil {
			logger.Error("create proxy model fail", map[string]interface{}{"proxy": p})
		}
		return
	}
	if status == consts.STATUS_YES {
		if proxyModel.CheckCount <= 20 {
			proxyModel.CheckCount = proxyModel.CheckCount + 1
		}
		if proxyModel.ActiveTime == 0 || proxyModel.Status == consts.STATUS_NO {
			proxyModel.ActiveTime = time.Now().Unix()
		}
	} else {
		proxyModel.ActiveTime = 0
		//set yes to no, need recheck
		if proxyModel.Status == consts.STATUS_YES && config.Get().RecheckCount > 0 {
			proxyModel.CheckCount = config.Get().RecheckCount
			status = consts.STATUS_RECHECK
		} else if proxyModel.Status == consts.STATUS_RECHECK {
			proxyModel.CheckCount = proxyModel.CheckCount - 1
			if proxyModel.CheckCount >= 0 {
				status = consts.STATUS_RECHECK
			}
		}
		if proxyModel.CheckCount < 0 && status == consts.STATUS_NO {
			logger.Debug("start delete fail proxy", map[string]interface{}{"proxy": proxyModel})
			if err := dao.ProxyDao.Delete(p.Host, p.Port, p.Proto); err != nil {
				logger.Error("delete proxy fail", logger.Fields{"host": p.Host, "port": p.Port})
			}
			return
		}
		proxyModel.CheckCount = proxyModel.CheckCount - 1
	}
	if p.Source != "" {
		proxyModel.Source = p.Source
	}
	//if ip is ok, update ip info
	if result && proxyModel.City == "" {
		ipInfo := ip.GetIpInfo(p.Host, p.Port)
		proxyModel.City = ipInfo.City
		proxyModel.Country = ipInfo.Country
		proxyModel.Isp = ipInfo.Isp
		proxyModel.Region = ipInfo.Region
	}
	proxyModel.Status = status
	proxyModel.UpdateTime = time.Now().Unix()
	proxyModel.User = p.User
	proxyModel.Password = p.Password
	_ = dao.ProxyDao.Save(proxyModel)
}

func (s *proxyService) DoGetProxy(getProxyService ProxyGetterInterface, pool *component.Pool) {
	for _, requestUrl := range getProxyService.GetUrlList() {
		contentBody := getProxyService.GetContentHtml(requestUrl)
		if contentBody == "" {
			time.Sleep(time.Second * 5)
			continue
		}
		proxyList := getProxyService.ParseHtml(contentBody)
		logger.Info("get ip list:", logger.Fields{"list": proxyList})
		var wg sync.WaitGroup = sync.WaitGroup{}
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			logger.FInfo("wait 10s ...")
			time.Sleep(time.Second * 10)
		}(&wg)
		for _, proxyArr := range proxyList {
			p := s.ParseProxyArr(proxyArr)
			p.Source = reflect.TypeOf(getProxyService).String()[14:]
			pool.RunTask(func() { s.CheckProxyAndSave(p) })
		}

		wg.Wait()
	}
}

func (s *proxyService) CheckProxyFormat(host string, port string) bool {
	ok, _ := regexp.Match(`^[\d\.]+$`, []byte(host))
	if !ok {
		return false
	}
	ok, _ = regexp.Match(`^\d+$`, []byte(port))
	return ok
}

func (s *proxyService) ParseProxyArr(proxyArr []string) dto.ProxyDto {
	p := dto.ProxyDto{
		Proto: consts.PROTO_HTTP,
	}
	for i, val := range proxyArr {
		switch i {
		case 0:
			p.Host = val
		case 1:
			p.Port = val
		case 2:
			p.Proto = val
		case 3:
			p.User = val
		case 4:
			p.Password = val
		}
	}
	return p
}

func (s *proxyService) GetProxyUrl(p dto.ProxyDto) string {
	if p.Proto == "" {
		p.Proto = consts.PROTO_HTTP
	}
	if p.User == "" {
		return fmt.Sprintf("%s://%s:%s", p.Proto, p.Host, p.Port)
	} else {
		return fmt.Sprintf("%s://%s:%s@%s:%s", p.Proto, p.User, p.Password, p.Host, p.Port)
	}
}
