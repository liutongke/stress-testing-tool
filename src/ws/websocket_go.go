package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	http2 "net/http"
	"stress-testing-tool/src/http"
	"stress-testing-tool/src/tool"
	"time"
)

type Ws struct {
}

func NewWs() *Ws {
	return &Ws{}
}

func GetWsHeader(flagParam *http.FlagParam) map[string]string {
	data, err := tool.GetFileData(flagParam.HeaderFile)
	if err != nil {
		fmt.Println(err, flagParam.HeaderFile)
		return nil
	}

	var result map[string]string
	err = json.Unmarshal(data, &result)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}
	return result
}

func (w *Ws) setHeader(flagParam *http.FlagParam) http2.Header {
	userHeader := GetWsHeader(flagParam)
	header := http2.Header{}
	for k, v := range userHeader {
		header.Add(k, v)
	}
	//req.Header.Add("x-token", "5656565656")
	//req.Header.Set("Content-Type", writer.FormDataContentType())
	return header
}

// HandlerWsUrl 拼接websocket请求的地址
func HandlerWsUrl(url, token, jobId string) string {
	return fmt.Sprintf("ws://%s/ws/%s/%s", url, token, jobId)
}

// StartWsConn 发起websocket连接
func StartWsConn(url string, flagParam *http.FlagParam) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, NewWs().setHeader(flagParam))
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
