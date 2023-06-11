package http

import (
	"stress-testing-tool/src/tool"
	"sync"
)

// http请求提前登录
func Http(userRunNum int, WgUser *sync.WaitGroup, ch chan<- *tool.ResponseRs, userReq *Request, flagParam *FlagParam) {

	defer func() {
		WgUser.Done()
	}()

	for i := 0; i < userRunNum; i++ {
		ch <- send(userReq, flagParam)
	}
}

func send(userReq *Request, flagParam *FlagParam) *tool.ResponseRs {
	return HttpDo(userReq, flagParam)
}
