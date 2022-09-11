package src

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	WgUser       sync.WaitGroup
	WgTask       sync.WaitGroup
	ResponseRsCh = make(chan *ResponseRs, 1000)

	ip    = "192.168.0.105:9500"
	wsUrl = "ws://192.168.0.105:9500/"

	list    = make(map[string]interface{})
	urlList = make(map[int]string) //请求地址的列表

	wsData = "{\n    \"id\":\"123123123\",\n    \"path\":\"/login\",\n    \"data\": {\n        \"username\":\"hello\"\n    }\n}"
)

const (
	userNum      = 10                   //并发数量
	userRunNum   = 10                   //单个并发执行的次数
	totalUserNum = userNum * userRunNum //总请求参数
)

func Run() {
	go ReceivingResults(ResponseRsCh) //统计处理
	WgTask.Add(1)
	LaunchWebsocket()
}

func LaunchWebsocket() {

	for i := 1; i <= userNum; i++ {
		WgUser.Add(1)
		go func(i int) {

			defer func() {
				WgUser.Done()
			}()

			conn, err := StartWsConn(wsUrl)

			if err != nil {
				fmt.Println(fmt.Sprintf("websocket err --------->%d", i))
			} else {
				defer func() {
					err := conn.Close()
					if err != nil {
						fmt.Println(fmt.Sprintf("关闭长连接时出现错误：%s", err.Error()))
					}
				}()

				timer := time.NewTimer(1 * time.Second) //一秒后激活时间
				n := 0

				for {
					select {
					case <-timer.C:
						timer.Reset(1 * time.Second) //重置倒计时
						n++

						WebSocketRequest(conn, WsRequestData(&WsRequest{
							Id:   strconv.Itoa(n),
							Path: "/",
							Data: map[string]string{"queue": strconv.Itoa(n), "user": strconv.Itoa(i)},
						}), ResponseRsCh)

						if n >= userRunNum {
							timer.Stop()                //到达指定次数结束时间
							time.Sleep(2 * time.Second) //让信息处理缓一会儿
							return
						}
					}
				}
			}
		}(i)
	}

	WgUser.Wait()
	close(ResponseRsCh)
	WgTask.Wait()
	fmt.Println("-------success-------")
}
