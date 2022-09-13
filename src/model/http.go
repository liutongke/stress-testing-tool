package model

import (
	"stress-testing-tool/src/tool"
	"sync"
	"time"
)

// http请求提前登录
func Http(userRunNum int, WgUser *sync.WaitGroup, ch chan<- *tool.ResponseRs, req *Request) {

	defer func() {
		WgUser.Done()
	}()

	for i := 0; i < userRunNum; i++ {

		isSucc, dataLen, requestTime := send(req)

		ch <- &tool.ResponseRs{
			IsSucc:      isSucc,
			DataLen:     dataLen,
			RequestTime: requestTime,
		}

	}
}

func send(req *Request) (isSucc bool, dataLen int, requestTime time.Duration) {

	resp, requestTime, err := PostFormData(req)
	isSucc = true
	if err != nil {
		isSucc = false
	}
	dataLen = int(resp.ContentLength)
	return
}
