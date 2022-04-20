package common

import (
	"fmt"
	"regexp"

	"proxy-collect/config"
	"proxy-collect/consts"
	"proxy-collect/dto"
)

func CheckProxyFormat(host string, port string) bool {
	ok, _ := regexp.Match(`^[\d\.]+$`, []byte(host))
	if !ok {
		return false
	}
	ok, _ = regexp.Match(`^\d+$`, []byte(port))
	return ok
}

func GetProxyUrl(p *dto.ProxyDto) string {
	proto := p.Proto
	if proto == "" {
		proto = consts.PROTO_HTTP
	}
	if p.User == "" {
		return fmt.Sprintf("%s://%s:%s", proto, p.Host, p.Port)
	} else {
		return fmt.Sprintf("%s://%s:%s@%s:%s", proto, p.User, p.Password, p.Host, p.Port)
	}
}

func GetTunnelUrl(tunnel *config.TunnelConfig) string {
	return fmt.Sprintf("%s://%s:%s", tunnel.Proto, tunnel.Host, tunnel.Port)
}
