package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"stress-testing-tool/src/tool"
	"time"
)

// HandlerWsUrl 拼接websocket请求的地址
func HandlerWsUrl(url, token, jobId string) string {
	return fmt.Sprintf("ws://%s/ws/%s/%s", url, token, jobId)
}

// StartWsConn 发起websocket连接
func StartWsConn(url string) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	return conn, err
}

// SendWsMsg 发送websocket信息
func SendWsMsg() {

}

func pushChan(duration time.Duration, succ bool, ch chan<- *tool.ResponseRs) {
	ch <- &tool.ResponseRs{
		IsSucc:      succ,
		DataLen:     0,
		RequestTime: duration,
	}
}

func WebSocketRequest(conn *websocket.Conn, wsSendData []byte) (isSucc bool, DataLen int, RequestTime time.Duration) {
	isSucc = false
	DataLen = 0
	RequestTime = 0
	var startTime = time.Now()

	writeErr := conn.WriteMessage(1, wsSendData)
	if writeErr != nil {
		return
	}

	_, msg, readErr := conn.ReadMessage()
	if readErr != nil {
		return
	}
	
	isSucc = true
	DataLen = len(msg)
	RequestTime = tool.DiffNano(startTime)
	return
}
