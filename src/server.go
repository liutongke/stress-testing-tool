package src

import (
	"math"
	"stress-testing-tool/src/model"
	"stress-testing-tool/src/tool"
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

func Run(req *model.Request, userNum, totalUserNum int) {
	go ReceivingResults(ResponseRsCh) //统计处理
	WgTask.Add(1)
	launchLink(req, userNum, totalUserNum)

	WgUser.Wait()
	close(ResponseRsCh)
	WgTask.Wait()
}

// 开启服务
func launchLink(req *model.Request, userNum, totalUserNum int) {

	for i := 0; i < userNum; i++ {
		WgUser.Add(1)

		userRunNum := int(math.Ceil(float64(totalUserNum / userNum)))

		go model.Http(userRunNum, &WgUser, ResponseRsCh, req)
	}

}
