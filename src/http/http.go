package http

import (
	"stress-testing-tool/src/tool"
	"sync"
	"time"
)

// http请求提前登录
func Http(userRunNum int, WgUser *sync.WaitGroup, ch chan<- *tool.ResponseRs, userReq *Request) {

	defer func() {
		WgUser.Done()
	}()

	for i := 0; i < userRunNum; i++ {

		isSucc, dataLen, requestTime := send(userReq)

		ch <- &tool.ResponseRs{
			IsSucc:      isSucc,
			DataLen:     dataLen,
			RequestTime: requestTime,
		}

	}
}

func send(userReq *Request) (isSucc bool, dataLen int, requestTime time.Duration) {

	resp, requestTime, isSucc := HttpDo(userReq)
	//resp, requestTime, err := PostFormData(req)
	dataLen = int(resp.ContentLength)
	return
}
