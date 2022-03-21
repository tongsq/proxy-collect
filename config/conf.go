package config

import (
	"fmt"
	"io/ioutil"

	"github.com/tongsq/go-lib/logger"
	"gopkg.in/yaml.v2"
)

var conf *confDto
var YamlPath string = "conf.yaml"

func Get() *confDto {
	if conf == nil {
		yamls, err := ioutil.ReadFile(YamlPath)
		if err != nil {
			ReadConfigError(err)
		}
		c := confDto{}
		err = yaml.Unmarshal(yamls, &c)
		if err != nil {
			ReadConfigError(err)
		}
		conf = &c
		logger.Info("load config success", logger.Fields{"conf": conf})
	}
	return conf
}

func ReadConfigError(err error) {
	logger.Error("read config file error", logger.Fields{"err": err, "path": YamlPath})
	panic("read config file error")
}

type confDto struct {
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
	}
	Api struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	PoolSize        int    `yaml:"poolSize"`
	LocalIpDataPath string `yaml:"localIpDataPath"`
}

func (c confDto) String() string {
	return fmt.Sprintf("%#v", c)
}
