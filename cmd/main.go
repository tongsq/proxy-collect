package main

import (
	"flag"
	"os"

	"github.com/tongsq/go-lib/logger"
	"proxy-collect/bootstrap"
	"proxy-collect/config"
)

var (
	configFile string
	servers    bootstrap.StringList
)

func main() {
	flag.Var(&servers, "S", "start server type")
	flag.StringVar(&configFile, "C", "conf.yaml", "")
	flag.Parse()
	if flag.NFlag() == 0 || len(servers) == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}
	config.YamlPath = configFile
	//init
	bootstrap.Bootstrap()

	for _, server := range servers {
		if server == bootstrap.ServerALl {
			go bootstrap.StartApiServer()
			go bootstrap.StartJobServer()
			break
		} else if server == bootstrap.ServerJob {
			go bootstrap.StartJobServer()
		} else if server == bootstrap.ServerApi {
			go bootstrap.StartApiServer()
		} else {
			logger.Error("unknown server", nil)
			os.Exit(1)
		}
	}
	select {}
}
