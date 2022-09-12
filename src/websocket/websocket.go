package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"stress-testing-tool/src/tool"
	"time"
)

// 拼接websocket请求的地址
func HandlerWsUrl(url, token, jobId string) string {
	return fmt.Sprintf("ws://%s/ws/%s/%s", url, token, jobId)
}

// 发起websocket连接
func StartWsConn(url string) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	return conn, err
}

// 发送websocket信息
func SendWsMsg() {

}

func pushChan(duration time.Duration, succ bool, ch chan<- *tool.ResponseRs) {
	ch <- &tool.ResponseRs{
		Time:      duration,
		IsSucceed: succ,
	}
}
func WebSocketRequest(conn *websocket.Conn, msg1 string, ch chan<- *tool.ResponseRs) string {

	//fmt.Println("发送的消息:", msg1)
	var startTime = time.Now()
	writeErr := conn.WriteMessage(1, []byte(msg1))
	if writeErr != nil {
		pushChan(0, false, ch)
		fmt.Println("WriteMessage错误信息:", writeErr)
		return ""
	} else {
		//fmt.Println("等待接收数据")
		//defer wg.Done()
		_, _, readErr := conn.ReadMessage()
		//fmt.Println(len(msg), string(msg))
		//fmt.Println("等待接收数据的数据咋还没来")
		if readErr != nil {
			pushChan(0, false, ch)

			fmt.Println("ReadMessage错误信息:", readErr)
			return ""
		}
		//if err == io.EOF {
		//	continue
		//}

		//fmt.Println("获取到的信息:", requestTime, string(msg))
		//d, err := JsonToData(string(msg))
		//var succ = false
		//if err != nil {
		//	succ = false
		//} else {
		//	if d.Err == 200 {
		//		succ = true
		//	} else {
		//		succ = false
		//	}
		//}

		pushChan(tool.DiffNano(startTime), true, ch)

		return ""
		//return string(msg)
		//if err != nil {
		//	return
		//}
	}
}
