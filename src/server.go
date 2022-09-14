package src

import (
	"math"
	"stress-testing-tool/src/http"
	"stress-testing-tool/src/request"
	"stress-testing-tool/src/tool"
	"stress-testing-tool/src/websocket"
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

func Run(req *http.Request, userNum, totalUserNum int) {
	go ReceivingResults(ResponseRsCh) //统计处理
	WgTask.Add(1)
	launchLink(req, userNum, totalUserNum)

	WgUser.Wait()
	close(ResponseRsCh)
	WgTask.Wait()
}

// 开启服务
func launchLink(req *http.Request, userNum, totalUserNum int) {

	for i := 0; i < userNum; i++ {
		WgUser.Add(1)

		userRunNum := int(math.Ceil(float64(totalUserNum / userNum))) //每个用户发送的请求次数

		switch req.Form {
		case request.FormTypeHTTP:
			go http.Http(userRunNum, &WgUser, ResponseRsCh, req)
		case request.FormTypeWebSocket:
			go websocket.Websocket(userRunNum, &WgUser, ResponseRsCh, req)
		default: //暂时不支持的类型
			WgUser.Done()
		}
	}

}
