package app

import (
	"sync/atomic"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"crypto/md5"
	"fmt"
	"config"
)

type AppConfig config.AppConfig

var (
	baseConf *string = new(string) //基础配置文件位置
	configMgr ConfigMgr
)

type ConfigChecker interface {
	Check() error
}

type ConfigMgr struct {
	configAtomic atomic.Value
	hook XPlugin
	initial bool
}

func BaseConf()(*string){
	return baseConf
}

func Config() *AppConfig {
	return configMgr.getConfig()
}

func ConfigInitialed() bool {
	return configMgr.initial
}

func InitConfig(cfg string) error {
	*baseConf = cfg
	fmt.Println(cfg)
	data, err := readFile(cfg)
	if err != nil {
		return err
	}
	config := new(AppConfig)
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return err
	}
	setConfig(config)
	//fmt.Printf("appconfig:%v\n", Config())
	var vv interface{} = configMgr.getConfig()
	if iCheck, ok := vv.(ConfigChecker); ok{
		if err = iCheck.Check(); err != nil {
			return err
		}
	}
	configMonitor.version = fmt.Sprintf("%v",md5.Sum(data))
	configMonitor.cur = data
	configMgr.initial = true
	return nil
}


func (p *ConfigMgr) getConfig() *AppConfig {
	return (*AppConfig)(config.GetConfig())
}

func setConfig(appConfig *AppConfig) {
	if configMgr.hook != nil {
		configMgr.hook.OnBeforeReloadConfig()
	}
	serverConfig := (*config.AppConfig)(appConfig)
	config.SetConfig(serverConfig)
	if configMgr.hook != nil {
		configMgr.hook.OnAfterReloadConfig()
	}
}

func readFile(file string) (data []byte, err error) {
	var f *os.File
	f, err = os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()
	data, err = ioutil.ReadAll(f)
	if err != nil {
		return
	}
	return
}

