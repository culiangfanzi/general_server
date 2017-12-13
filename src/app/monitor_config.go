package app

import (
	"time"
	"gopkg.in/redis.v3"
	"fmt"
	"crypto/md5"
	"gopkg.in/yaml.v2"
)

const (
	typeRedis = "redis"
	typeFile = "file"
)

const(
	refreshInterval = 10
)

var (
	configMonitor *ConfigMonitor = &ConfigMonitor{}
)

type ConfigMonitor struct {
	version string
	old     []byte
	cur     []byte
	conn	*redis.Client
}

func ConfigMon() *ConfigMonitor {
	return configMonitor
}

//监听配置信息，后台任务
func (p *ConfigMonitor) Process() {
	//读配置
	//判断类型
	var data []byte
	var err error
	for {
		<-time.After(time.Second * refreshInterval)
		data, err = p.ReadLocalConfig()
		newV := fmt.Sprintf("%v",md5.Sum(data))
		if newV == p.version {
			continue
		}
		appConfig := new(AppConfig)
		err = yaml.Unmarshal(data, appConfig)
		if err != nil {
			//to do log
			fmt.Printf("config unmarshal err:%v\n", err)
			continue
		}
		if string(data) == "" {
			//to do log
			fmt.Printf("config data empty:%v\n", err)
			continue
		}
		var vv interface{} = appConfig
		if iCheck, ok := vv.(ConfigChecker); ok{
			if err = iCheck.Check(); err != nil {
				//to do log
				fmt.Printf("config checker failed:%v\n", err)
				continue
			}
		}
		p.version = newV
		setConfig(appConfig)
		p.old = p.cur
		p.cur = data
	}
}

func (p *ConfigMonitor) ReadLocalConfig() ([]byte, error) {
	file := *baseConf
	return readFile(file)
}


