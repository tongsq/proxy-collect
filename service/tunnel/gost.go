package tunnel

import (
	"crypto/tls"
	"errors"
	_ "net/http/pprof"
	"os"
	"time"

	"github.com/go-log/log"
	"github.com/tongsq/go-lib/logger"
	"github.com/tongsq/gost"
	"proxy-collect/config"
	"proxy-collect/dao"
	"proxy-collect/dto"
)

var (
	baseCfg       = &baseConfig{}
	NodeGroupList = []*gost.NodeGroup{}
)

func StartTunnels() {
	gost.SetLogger(&gost.LogLogger{})
	baseCfg.Debug = config.Get().Tunnel.Debug
	baseCfg.ServeNodes = config.Get().Tunnels
	baseCfg.ChainNodes = []dto.ProxyDto{}
	proxyList, err := getProxyList()
	if err != nil {
		logger.Error("get active ip fail", logger.Fields{"err": err})
	} else {
		baseCfg.ChainNodes = proxyList
	}
	baseCfg.ChinLevel = config.Get().Tunnel.TunnelLevel
	// NOTE: as of 2.6, you can use custom cert/key files to initialize the default certificate.
	tlsConfig, err := tlsConfig(defaultCertFile, defaultKeyFile, "")
	if err != nil {
		// generate random self-signed certificate.
		cert, err := gost.GenCertificate()
		if err != nil {
			log.Log(err)
			os.Exit(1)
		}
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
	} else {
		log.Log("load TLS certificate files OK")
	}

	gost.DefaultTLSConfig = tlsConfig

	if err := start(); err != nil {
		log.Log(err)
		os.Exit(1)
	}

	select {}
}

func start() error {
	gost.Debug = baseCfg.Debug

	var routers []router
	rts, err := baseCfg.route.GenRouters()
	if err != nil {
		return err
	}
	routers = append(routers, rts...)

	for _, route := range baseCfg.Routes {
		rts, err := route.GenRouters()
		if err != nil {
			return err
		}
		routers = append(routers, rts...)
	}

	if len(routers) == 0 {
		return errors.New("invalid config")
	}
	for i := range routers {
		go routers[i].Serve()
	}
	go StartRefreshNodeGroupList()
	return nil
}

func getProxyList() ([]dto.ProxyDto, error) {
	proxies, err := dao.ProxyDao.GetActiveList()
	if err != nil {
		logger.Error("get active ip fail", logger.Fields{"err": err})
		return nil, err
	}
	proxyList := []dto.ProxyDto{}
	for _, proxy := range proxies {
		var p = dto.NewProxyDto(proxy)
		proxyList = append(proxyList, p.ProxyDto)
	}
	return proxyList, nil
}

func StartRefreshNodeGroupList() {
	for {
		time.Sleep(time.Second * time.Duration(config.Get().Tunnel.Refresh))
		logger.Debug("start run RefreshNodeGroupList", nil)
		proxyList, err := getProxyList()
		if err != nil {
			logger.Error("get active ip fail", logger.Fields{"err": err})
			continue
		}
		if err = RefreshNodeGroupList(proxyList); err != nil {
			logger.Error("RefreshNodeGroupList fail", map[string]interface{}{"err": err})
		}

	}
}

func RefreshNodeGroupList(proxyList []dto.ProxyDto) error {
	// parse the base nodes
	nodes, err := parseChainNode(proxyList)
	if err != nil {
		return err
	}

	nid := 1 // node ID
	for i := range nodes {
		nodes[i].ID = nid
		nid++
	}
	for _, ngroup := range NodeGroupList {
		ngroup.SetNodes(nodes...)

		ngroup.SetSelector(nil,
			gost.WithFilter(
				&gost.FailFilter{
					MaxFails:    config.Get().Tunnel.MaxFails,
					FailTimeout: time.Duration(config.Get().Tunnel.FailTimeout) * time.Second,
				},
				&gost.InvalidFilter{},
			),
			gost.WithStrategy(gost.NewStrategy(config.Get().Tunnel.Strategy)),
		)
	}
	return nil
}
