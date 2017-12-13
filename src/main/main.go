package main

import (
	"fmt"
	"app"
	"flag"
	"time"
)

var(
	cfg *string
	cmd *string
	//sType *string
	//fType *string
)

func init() {
	cfg = flag.String("c", "", "config file")
	flag.Parse()
	if *cfg == "" {
		*cfg = `C:\Users\zxh\ad_logman\config\scm_config.yaml`
	}
	return
}


func main() {
	fmt.Println(time.Now().Format("2006/01/02 15:04:05.999"))
	var err error
	hook := &XMainHook{}
	processor := &DeamonService{}
	if err = hook.OnStartup(); err != nil {
		fmt.Println("startup app failed, err:", err)
		return
	}
	err = app.Run(*cfg, hook, processor)
	if err != nil {
		fmt.Println("run app failed, err:", err)
	}
	hook.OnShutdown()
}



