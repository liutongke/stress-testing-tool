package http

import (
	"stress-testing-tool/src/tool"
	"sync"
	"time"
)

// http请求提前登录
func Http(userRunNum int, url string, WgUser *sync.WaitGroup, ch chan<- *tool.ResponseRs) {

	defer func() {
		WgUser.Done()
	}()

	for i := 0; i < userRunNum; i++ {

		isSucc, dataLen, requestTime := send(url)

		ch <- &tool.ResponseRs{
			IsSucc:      isSucc,
			DataLen:     dataLen,
			RequestTime: requestTime,
		}

	}
}

func send(url string) (isSucc bool, dataLen int, requestTime time.Duration) {
	resp, requestTime, err := PostFormData(url, map[string]string{"code": "wxdev"}, nil)
	isSucc = true
	if err != nil {
		isSucc = false
	}
	dataLen = int(resp.ContentLength)
	return
}
