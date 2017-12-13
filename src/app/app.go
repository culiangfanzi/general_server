package app

import (
	//"runtime"
)

type Processor interface {
	 Process() error
}

func Run(cfg string, hook XPlugin, processor Processor) error{
	var err error
	//runtime.GOMAXPROCS(runtime.NumCPU())
	if err = RegisterPlugin(hook); err != nil {
		return err
	}
	g_pluginMgr.OnBeforeLoadConfig()
	if err = InitConfig(cfg);err != nil {
		return err
	}
	g_pluginMgr.OnAfterLoadConfig()
	configMgr.hook = g_pluginMgr
	//init 打点库
	//init 资源状态库
	//init 管理库
	go ConfigMon().Process()
	if processor == nil {
		return nil
	}
	return processor.Process()
}
