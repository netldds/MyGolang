package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/imdario/mergo"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"reflect"
)

type NH struct {
	Name string
}
type ServerConfig struct {
	name string
}

func LoadConfig(env string, dir string) ServerConfig {
	var configName string
	if env == "product" {
		configName = "product.yml"
	} else if env == "test" {
		configName = "test.yml"
	} else {
		configName = "development.yml"
	}

	localConfigFile := path.Join(dir, "local.yml")
	configFile := path.Join(dir, configName)

	var localObj ServerConfig
	if _, err := os.Stat(localConfigFile); !os.IsNotExist(err) {
		raw, err2 := ioutil.ReadFile(localConfigFile)
		if err2 != nil {
			glog.Infof("Read local.yml error: %v", err)
			panic("failed to parse config file ")
		}
		err2 = yaml.Unmarshal(raw, &localObj)
		if err2 != nil {
			glog.Infof("local.yml format error: %v", err)
			panic("failed to parse config file ")
		}
	}

	var configObj ServerConfig
	if _, err := os.Stat(configFile); !os.IsNotExist(err) {
		raw, err2 := ioutil.ReadFile(configFile)
		if err2 != nil {
			glog.Infof("Rread %s error: %v", configFile, err2)
			panic("failed to parse config file ")
		}
		err2 = yaml.Unmarshal(raw, &configObj)
		if err2 != nil {
			glog.Infof("%s format error: %v", configFile, err)
			panic("failed to parse config file ")
		}
	}

	err := mergo.MergeWithOverwrite(&configObj, localObj)
	if err != nil {
		glog.Infof("Merge config faileD: %v", err)
	}
	return configObj
}

func main() {
	s1 := ServerConfig{name: "123"}
	v1 := reflect.ValueOf(&s1).Elem()
	f1 := v1.Type().Field(0)
	fmt.Println(f1.PkgPath)
	fmt.Println(f1.Anonymous)
}
