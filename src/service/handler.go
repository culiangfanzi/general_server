package service

import(
	"net/http"
	"log"
	"sync"
)

const(
	TYPE_HANDLER_S2S	= 1
	TYPE_HANDLER_PMP	= 2

	TYPE_VIDEO		= 1
	TYPE_NOVIDEO		= 0
)

var(
	_succCnt = int32(0)
	_failCnt = int32(0)
	lockLock = sync.Mutex{}
)

type InnerRequest struct {
	IsVideo int
	HandlerId int
}

func HandlerS2s(w http.ResponseWriter, r *http.Request) {
	tr := new(InnerRequest)
	tr.HandlerId = TYPE_HANDLER_S2S
	req := &AdProcessor{
		RawReq: r,
		InnerReq: tr,
	}
	w.Write(req.Process())
}

func HandlerPmp(w http.ResponseWriter, r *http.Request) {
	tr := new(InnerRequest)
	tr.HandlerId = TYPE_HANDLER_PMP
	req := &AdProcessor{
		RawReq: r,
		InnerReq: tr,
	}
	w.Write(req.Process())
}

func HandlerPdb(w http.ResponseWriter, r *http.Request) {
	tr := new(InnerRequest)
	tr.HandlerId = TYPE_HANDLER_PMP
	req := &AdProcessor{
		RawReq: r,
		InnerReq: tr,
	}
	w.Write(req.Process())
}

type AdProcessor struct {
	RawReq *http.Request
	InnerReq *InnerRequest
}


func (p *AdProcessor) Process() []byte {
	var flag bool = true
	log.Printf("req data:%v\n", p.RawReq.RequestURI)
	status := execControl(p.InnerReq.HandlerId)
	lockLock.Lock()
	defer func() {
		lockLock.Unlock()
		log.Printf("req end:%v\n", flag)
	}()
	if status {
		_succCnt += 1
	}else {
		_failCnt += 1
		flag = false
		return []byte(`<VAST version="3.0"></VAST>`)
	}
	return loadVast(p.InnerReq.HandlerId)
}

func logSucc() {
	for {
		lockLock.Lock()
		m := _failCnt
		n := _succCnt
		_failCnt = 0
		_succCnt = 0
		lockLock.Unlock()
		log.Printf("succ_cnt:%d fail_cnt:%d\n", m, n)
	}

}
