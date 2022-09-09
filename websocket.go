package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

//拼接websocket请求的地址
func webSocketUrl(url, token, jobId string) string {
	return fmt.Sprintf("ws://%s/ws/%s/%s", url, token, jobId)
}

func WsConn(url string) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	return conn, err
}

func WebSocketRequest(conn *websocket.Conn, msg1 string, ch chan<- *RequestResults) string {
	fmt.Println("发送的消息:", msg1)
	var startTime = time.Now()
	writeErr := conn.WriteMessage(1, []byte(msg1))
	if writeErr != nil {
		ch <- &RequestResults{
			Time:      0,
			IsSucceed: false,
		}
		fmt.Println("WriteMessage错误信息:", writeErr)
		return ""
	} else {
		fmt.Println("等待接收数据")
		//defer wg.Done()
		_, msg, readErr := conn.ReadMessage()
		fmt.Println("等待接收数据的数据咋还没来")
		if readErr != nil {
			//conn.Close()
			ch <- &RequestResults{
				Time:      0,
				IsSucceed: false,
			}
			fmt.Println("ReadMessage错误信息:", readErr)
			return ""
		}
		//if err == io.EOF {
		//	continue
		//}
		requestTime := DiffNano(startTime)
		fmt.Println("获取到的信息:", requestTime, string(msg))
		d, err := JsonToData(string(msg))
		var succ = false
		if err != nil {
			succ = false
		} else {
			if d.Err == 200 {
				succ = true
			} else {
				succ = false
			}
		}
		ch <- &RequestResults{
			Time:      requestTime,
			IsSucceed: succ,
		}
		return string(msg)
		//if err != nil {
		//	return
		//}
	}
}
