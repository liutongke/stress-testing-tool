package http_client

import (
	"stress-testing-tool/src/tool"
	"sync"
)

var (
	WgHTTPRequest      sync.WaitGroup
	WgHTTPStressTester sync.WaitGroup
	ResponseRsCh       = make(chan *tool.ResponseRs, 1000)
)

func Run(flagParam *ABConfig) {
	go ReceivingResults(ResponseRsCh) //统计处理
	WgHTTPStressTester.Add(1)
	launchLink(flagParam)

	WgHTTPRequest.Wait()
	close(ResponseRsCh)
	WgHTTPStressTester.Wait()
}

// 开启服务
func launchLink(flagParam *ABConfig) {
	for i := 0; i < flagParam.Concurrency; i++ {
		WgHTTPRequest.Add(1)
		go StartHTTPRequest(&WgHTTPRequest, ResponseRsCh, flagParam)
		//switch userReq.Form {
		//case http.FormTypeHTTP:
		//	go http_client.StartHTTPRequest(WgHTTPRequest, ResponseRsCh, flagParam)
		//case http.FormTypeWebSocket:
		//go ws.Websocket(userRunNum, &WgHTTPRequest, ResponseRsCh, userReq, flagParam)
		//case http.FormTypeProcess: //流程测试
		//go process.Start(userRunNum, &WgUser, ResponseRsCh, postFile)
		//default: //暂时不支持的类型
		//	WgHTTPRequest.Done()
	}
}

//}
