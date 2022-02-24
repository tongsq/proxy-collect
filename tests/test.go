package tests

import (
	"fmt"

	"github.com/tongsq/go-lib/request"
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
	//d := config.Get()
	//fmt.Println(d.Redis.MaxIdle, d.Redis.Address)
	//service.ProxyService.DoGetProxy(service.NewGetProxyFanQie(), global.Pool)
	u := "http://10.4.12.12/inside/v1/ssq/set"
	h := &request.RequestHeaderDto{
		//UserAgent:               consts.USER_AGENT,
		UpgradeInsecureRequests: "1",
		//Host:                    "socket-apit.weipaitang.com",
	}
	data, err := request.WebPost(u, nil, h, nil)
	fmt.Printf("%#v", data)
	if err != nil || data.HttpCode != request.HTTP_CODE_OK {
		fmt.Println("test post request fail", err)
	} else {
		fmt.Println("test post request success", err)
	}
}
