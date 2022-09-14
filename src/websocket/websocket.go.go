package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"stress-testing-tool/src/http"
	"stress-testing-tool/src/tool"
	"sync"
	"time"
)

func Websocket(userRunNum int, WgUser *sync.WaitGroup, ch chan<- *tool.ResponseRs, req *http.Request) {

	defer func() {
		WgUser.Done()
	}()

	conn, err := StartWsConn(req.URL)
	if err != nil {
		fmt.Println(fmt.Sprintf("websocket err --------->"))
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println(fmt.Sprintf("关闭长连接时出现错误：%s", err.Error()))
		}
	}()

	sendMsg(userRunNum, conn, ch, req)
}

func sendMsg(userRunNum int, conn *websocket.Conn, ch chan<- *tool.ResponseRs, req *http.Request) {

	timer := time.NewTimer(1 * time.Second) //一秒后激活时间
	n := 0

	for {
		select {
		case <-timer.C:
			timer.Reset(1 * time.Second) //重置倒计时
			n++

			isSucc, dataLen, requestTime := WebSocketRequest(conn, req.Body)

			ch <- &tool.ResponseRs{
				IsSucc:      isSucc,
				DataLen:     dataLen,
				RequestTime: requestTime,
			}

			if n >= userRunNum {
				timer.Stop()                //到达指定次数结束时间
				time.Sleep(2 * time.Second) //让信息处理缓一会儿
				return
			}
		}
	}
}
