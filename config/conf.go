package config

import (
	"fmt"
	"github.com/tongsq/go-lib/logger"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

var conf *ConfDto
var YamlPath = "conf.yaml"
var configRefreshHandlers []func(old, new *ConfDto)

func Get() *ConfDto {
	if conf == nil {
		yamls, err := ioutil.ReadFile(YamlPath)
		if err != nil {
			ReadConfigError(err)
		}
		c := ConfDto{}
		err = yaml.Unmarshal(yamls, &c)
		if err != nil {
			ReadConfigError(err)
		}
		conf = &c
		logger.Info("load config success", logger.Fields{"conf": conf})
	}
	return conf
}

func StartLoadConfig() {
	LoadConfig()
	go func() {
		t := time.NewTicker(time.Second * 10)
		for {
			<-t.C
			LoadConfig()
		}
	}()
}

func RegisterConfigRefreshHandler(f func(old, new *ConfDto)) {
	configRefreshHandlers = append(configRefreshHandlers, f)
}

func LoadConfig() {
	yamls, err := ioutil.ReadFile(YamlPath)
	if err != nil {
		ReadConfigError(err)
	}
	c := ConfDto{}
	err = yaml.Unmarshal(yamls, &c)
	if err != nil {
		ReadConfigError(err)
	}
	if conf != nil {
		for _, h := range configRefreshHandlers {
			h(conf, &c)
		}
	}
	conf = &c
	logger.Debug("load config success", logger.Fields{"conf": conf})
}

func ReadConfigError(err error) {
	logger.Error("read config file error", logger.Fields{"err": err, "path": YamlPath})
	panic("read config file error")
}

type ConfDto struct {
	Dao   string `yaml:"dao"`
	Redis struct {
		MaxIdle   int    `yaml:"MaxIdle"`
		MaxActive int    `yaml:"MaxActive"`
		Network   string `yaml:"Network"`
		Address   string `yaml:"Address"`
		Password  string `yaml:"Password"`
	}
	Database struct {
		Dialect string `yaml:"Dialect"`
		Args    string `yaml:"Args"`
		MaxIdle int    `yaml:"MaxIdle"`
		MaxOpen int    `yaml:"MaxOpen"`
	}
	Api struct {
		Host  string `yaml:"host"`
		Port  string `yaml:"port"`
		Token string `yaml:"token"`
	}
	PoolSize        int    `yaml:"poolSize"`
	LocalIpDataPath string `yaml:"localIpDataPath"`
	RecheckCount    int64  `yaml:"recheckCount"`
	MaxPing         int64  `yaml:"maxPing"`
	UpdateIpInfo    bool   `yaml:"updateIpInfo"`
	Log             struct {
		LogLevel     logger.Level `yaml:"logLevel"`
		ErrorLogFile string       `yaml:"errorLogFile"`
	}
	Tunnel struct {
		TunnelLevel int    `yaml:"tunnelLevel"`
		Refresh     int64  `yaml:"refresh"`
		Debug       bool   `yaml:"debug"`
		Strategy    string `yaml:"strategy"`
		MaxFails    int    `yaml:"maxFails"`
		FailTimeout int    `yaml:"failTimeout"`
	}
	Tunnels []TunnelConfig `yaml:"tunnels"`
	Getters []Getter       `yaml:"getters"`
}

func (c ConfDto) String() string {
	return fmt.Sprintf("%#v", c)
}
