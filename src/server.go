package src

import (
	"stress-testing-tool/src/http"
	"stress-testing-tool/src/tool"
	"sync"
)

var (
	userNum      = 100                   //并发数量
	userRunNum   = 1000                 //单个并发执行的次数
	totalUserNum = userNum * userRunNum //总请求参数

	WgUser       sync.WaitGroup
	WgTask       sync.WaitGroup
	ResponseRsCh = make(chan *tool.ResponseRs, 1000)

	httpUrl = "http://192.168.0.105:9500/"
	wsUrl   = "ws://192.168.0.105:9500/"

	list    = make(map[string]interface{})
	urlList = make(map[int]string) //请求地址的列表

)

func Run() {

	go ReceivingResults(ResponseRsCh) //统计处理
	WgTask.Add(1)
	launchLink()

	WgUser.Wait()
	close(ResponseRsCh)
	WgTask.Wait()
}

// 开启服务
func launchLink() {

	for i := 0; i < userNum; i++ {
		WgUser.Add(1)
		go http.Http(userRunNum, httpUrl, &WgUser, ResponseRsCh)
	}

}
