package service

import(
	"net/http"
	"log"
	"sync/atomic"
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
	tr.HandlerId = 1
	req := &AdProcessor{
		RawReq: r,
		InnerReq: tr,
	}
	w.Write(req.Process())
}

func HandlerPmp(w http.ResponseWriter, r *http.Request) {
	tr := new(InnerRequest)
	tr.HandlerId = 2
	req := &AdProcessor{
		RawReq: r,
		InnerReq: tr,
	}
	w.Write(req.Process())
}

func HandlerPdb(w http.ResponseWriter, r *http.Request) {
	tr := new(InnerRequest)
	tr.HandlerId = 2
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
	log.Printf("req data:%v\n", p.RawReq.Form)
	status := execControl(p.InnerReq.HandlerId)
	lockLock.Lock()
	if status {
		atomic.AddInt32(&_succCnt, 1)
	}else {
		atomic.AddInt32(&_failCnt, 1)
	}
	lockLock.Unlock()
	return loadVast(p.InnerReq.IsVideo)
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
