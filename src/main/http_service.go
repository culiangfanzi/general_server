package main

import (
	"log"
	"os"
	"net/http"
	"config"
	"service"
)

func init() {
	log.SetOutput(os.Stdout)
}

type DeamonService struct {

}

func (p *DeamonService) Process()(err error) {
	http.HandleFunc("/test/adx", service.HandlerPmp)
	http.HandleFunc("/test/s2s", service.HandlerS2s)
	http.HandleFunc("/test/pdb", service.HandlerPdb)
	err = http.ListenAndServe(config.GetConfig().BindingConf.Addr, nil)
	if err != nil {
		log.Printf("server exited unexpectedly!!!, err:%v\n", err)
	}
	return err
}

