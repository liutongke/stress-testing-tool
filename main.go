package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

//var wg sync.WaitGroup
//var wg1 sync.WaitGroup
//var list = make(map[string]interface{})

const (
	concurrence    = 1 //并发数量
	concurrenceNum = 1 //单个并发执行的次数

)

//var ip = "118.195.232.124:12223"

//var ip = "127.0.0.1:12223"
//var urlList = make(map[int]string) //请求地址的列表

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://192.168.0.105:9500/", nil)
	//conn, err := WsConn("ws://192.168.0.105:9500/")
	if err != nil {
		fmt.Println(err)
	}
	err = conn.WriteMessage(1, []byte("{\n    \"id\":\"123123123\",\n    \"path\":\"/login\",\n    \"data\": {\n        \"username\":\"hello\"\n    }\n}"))
	if err != nil {
		fmt.Println(err, "发送消息错误")
	}
	messageType, p, error := conn.ReadMessage()
	fmt.Println(messageType, string(p), error)
}

//func mains() {
	//ch := make(chan *RequestResults, 1000)
	//go ReceivingResults(ch)
	//wg1.Add(1)
	//for i := 0; i < concurrence; i++ { //单独先生成减轻数据库压力
	//	url := fmt.Sprintf("http://%s/auth/login", ip)
	//	//payload := fmt.Sprintf("{\"code\":\"wxdev%d\"}", i)
	//	//fmt.Println(payload)
	//	//data, err := PostJosn(url, payload)
	//	data, err := PostFormData(url, map[string]string{"code": fmt.Sprintf("wxdev%d", i)})
	//	if err != nil {
	//		fmt.Println(fmt.Sprintf("post json ------>%d err:%s", i, err.Error()))
	//	} else {
	//		fmt.Println(fmt.Sprintf("post json succ --------->%d payload:%s", i, data))
	//	}
	//	token = gjson.Get(data, "data.token").String()
	//
	//	createJob, _ := PostFormData(fmt.Sprintf("http://%s/user/job/create", ip), map[string]string{"job": "1", "jobCardId": "1"})
	//	fmt.Println(createJob)
	//	jobId := gjson.Get(createJob, "data.jobId").String()
	//	if len(jobId) <= 0 {
	//		fmt.Println(fmt.Sprintf("post json jobId err --------->%s", jobId))
	//	}
	//	fmt.Println(webSocketUrl(ip, token, jobId))
	//	urlList[i] = webSocketUrl(ip, token, jobId)
	//}
	//for i := 0; i < concurrence; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer func() {
	//			wg.Done()
	//		}()
	//		conn, err := WsConn(urlList[i])
	//		if err != nil {
	//			fmt.Println(fmt.Sprintf("websocket err --------->%d", i))
	//		} else {
	//			defer func() {
	//				fmt.Println("执行了defer close")
	//				wsCloseErr := conn.Close()
	//				if wsCloseErr != nil {
	//					fmt.Println(fmt.Sprintf("关闭长连接时出现错误：%s", wsCloseErr.Error()))
	//				}
	//			}()
	//			timer := time.NewTimer(1 * time.Second) //一秒后激活时间
	//			n := 0
	//			WebSocketRequest(conn, WsRequestData(i, "1.0.0", "/Select/Map", map[string]string{"mapId": "1"}), ch)
	//			for {
	//				select {
	//				case <-timer.C:
	//					timer.Reset(1 * time.Second) //重置倒计时
	//					n++
	//					fmt.Println(n)
	//					WebSocketRequest(conn, WsRequestData(i, "1.0.0", "/SevenDays/List", map[string]string{"test": "111111111111111111"}), ch)
						//roomId := gjson.Get(selectMapInfo, "data.roomId").String()
						//if roomId != "" {
						//	WebSocketRequest(conn, WsRequestData(i, "1.0.0", "/Monster/List", map[string]string{"roomId": roomId}), ch)
						//}
						//Read(conn, "{\"id\":123,\"path\":\"/Bag/List\",\"ver\":\"1.0.0\",\"data\":\"\"}")
						//Read(conn, "{\"id\":123,\"path\":\"/Map/List\",\"ver\":\"1.0.0\",\"data\":\"\"}")
//						if n >= concurrenceNum {
//							timer.Stop()                //到达指定次数结束时间
//							time.Sleep(2 * time.Second) //让信息处理缓一会儿
//							return
//						}
//					}
//				}
//			}
//		}(i)
//	}
//	wg.Wait()
//	close(ch)
//	wg1.Wait()
//	fmt.Println("-------success-------")
//}

//func ReceivingResults(ch <-chan *RequestResults) {
//	// 时间
//	var (
//		processingTime time.Duration //处理总时间
//		maxTime        time.Duration // 最大时长
//		minTime        time.Duration // 最小时长
//		successNum     uint64        // 成功处理数，code为0
//		failureNum     uint64        // 处理失败数，code不为0
//	)
//
//	for data := range ch {
//		processingTime += data.Time
//		if maxTime <= data.Time {
//			maxTime = data.Time
//		}
//		if minTime > data.Time && minTime != 0 {
//			minTime = data.Time
//		}
//		// 是否请求成功
//		if data.IsSucceed == true {
//			successNum = successNum + 1
//		} else {
//			failureNum = failureNum + 1
//		}
//	}
//	fmt.Println(fmt.Sprintf("最大请求时长:%s,最小请求时长:%s,成功的处理数:%d,失败的请求数:%d,处理总时长:%s", maxTime, minTime, successNum, failureNum, processingTime))
//	wg1.Done()
//}
