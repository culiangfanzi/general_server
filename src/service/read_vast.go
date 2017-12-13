package service

import (
	"config"
	"log"
	"io/ioutil"
	"math/rand"
	"time"
)

type RawVast  []byte

var(
	typeVastMap = make(map[int][]RawVast)
)


func InitVast() {
	for _, videoFile := range config.GetConfig().VastConf.VideoList {
		d, err := ioutil.ReadFile(videoFile)
		if err != nil {
			log.Printf("load file failed: %v", err)
		} else {
			typeVastMap[TYPE_HANDLER_S2S] = append(typeVastMap[TYPE_HANDLER_S2S], RawVast(d))
		}
	}

	for _, noVideoFile := range config.GetConfig().VastConf.NoVideoList {
		d, err := ioutil.ReadFile(noVideoFile)
		if err != nil {
			log.Printf("load file failed: %v", err)
		} else {
			typeVastMap[TYPE_HANDLER_PMP] = append(typeVastMap[TYPE_HANDLER_PMP], RawVast(d))
		}
	}
	//log.Printf("all vast data:%v\n", typeVastMap)
}

func loadVast(isVideo int) []byte {
	availables := typeVastMap[isVideo]
	if len(availables) <= 0 {
		return nil
	}
	index := rand.Intn(len(availables))
	return []byte(availables[index])
}

func execControl(t int) bool {
	ctr := config.DetailControlConfig{}
	switch t {
	case TYPE_HANDLER_S2S:
		ctr = config.GetConfig().ControlConf.S2cControl
	case TYPE_HANDLER_PMP:
		ctr = config.GetConfig().ControlConf.S2cControl
	}
	randCost := rand.Intn(ctr.AvgCost * 2)
	tt := ctr.AvgCost + randCost

	<- time.After(time.Millisecond * time.Duration(tt))

	sucMax := rand.Intn(100)
	rate := float32(sucMax)/100.0
	log.Printf("sleep time :%d, rate:%v\n", tt, rate)
	if rate < ctr.SuccRate {
		return true
	}else {
		return false
	}
}

