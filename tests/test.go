package main

import (
	"fmt"
	"proxy-collect/config"
	"proxy-collect/global"
	"proxy-collect/service"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func main() {
	d := config.Get()
	fmt.Println(d.Redis.MaxIdle, d.Redis.Address)
	service.ProxyService.DoGetProxy(service.NewGetProxyFanQie(), global.Pool)
}
