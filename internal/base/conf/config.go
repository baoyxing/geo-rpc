package conf

import (
	"github.com/baoyxing/micro-extend/pkg/config/kitex_conf"
	"github.com/baoyxing/micro-extend/pkg/config/log"
	"github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kr/pretty"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Service       kitex_conf.Service `json:"service" mapstructure:"service" yaml:"service"`
	Server        kitex_conf.Server  `json:"server" mapstructure:"server" yaml:"server"`
	Logger        log.Logger         `json:"logger" mapstructure:"logger" yaml:"logger"`
	Ip2RegionPath string             `json:"ip2_region_path" mapstructure:"ip2_region_path" yaml:"ip2_region_path"` //
}

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	prefix := "conf"
	confFileRelPath := filepath.Join(prefix, filepath.Join(GetEnv(), "conf.yaml"))
	content, err := ioutil.ReadFile(confFileRelPath)
	if err != nil {
		panic(err)
	}
	conf = new(Config)
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		klog.Error("parse yaml error - %v", err)
		panic(err)
	}
	if err := validator.Validate(conf); err != nil {
		klog.Error("validate config error - %v", err)
		panic(err)
	}
	pretty.Printf("%+v\n", conf)
}

func GetEnv() string {
	e := os.Getenv("GO_ENV")
	if len(e) == 0 {
		return "test"
	}
	return e
}
