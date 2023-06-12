package src

import (
	"math"
	"stress-testing-tool/src/http"
	"stress-testing-tool/src/tool"
	"stress-testing-tool/src/ws"
	"sync"
)

var (
	//userNum      = 100                  //并发数量
	//userRunNum   = 100                  //单个并发执行的次数
	//totalUserNum = userNum * userRunNum //总请求参数

	WgUser       sync.WaitGroup
	WgTask       sync.WaitGroup
	ResponseRsCh = make(chan *tool.ResponseRs, 1000)

	//httpUrl = "http://192.168.0.105:9500/"
	//wsUrl   = "ws://192.168.0.105:9500/"
	//
	//list    = make(map[string]interface{})
	//urlList = make(map[int]string) //请求地址的列表

)

func Run(userReq *http.Request, flagParam *http.FlagParam) {
	go ReceivingResults(ResponseRsCh) //统计处理
	WgTask.Add(1)
	launchLink(userReq, flagParam)

	WgUser.Wait()
	close(ResponseRsCh)
	WgTask.Wait()
}

// 开启服务
func launchLink(userReq *http.Request, flagParam *http.FlagParam) {

	for i := 0; i < flagParam.UserNum; i++ {
		WgUser.Add(1)
		userRunNum := int(math.Ceil(float64(flagParam.TotalUserNum / flagParam.UserNum))) //每个用户发送的请求次数
		switch userReq.Form {
		case http.FormTypeHTTP:
			go http.Http(userRunNum, &WgUser, ResponseRsCh, userReq, flagParam)
		case http.FormTypeWebSocket:
			go ws.Websocket(userRunNum, &WgUser, ResponseRsCh, userReq, flagParam)
		case http.FormTypeProcess: //流程测试
			//go process.Start(userRunNum, &WgUser, ResponseRsCh, postFile)
		default: //暂时不支持的类型
			WgUser.Done()
		}
	}

}
