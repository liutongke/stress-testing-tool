package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Request 请求数据
type Request struct {
	URL       string            // URL
	Form      string            // http/webSocket/tcp
	Method    string            // 方法 GET/POST/PUT
	Headers   map[string]string // Headers
	Body      string            // body
	Verify    string            // 验证的方法
	Timeout   time.Duration     // 请求超时时间
	Debug     bool              // 是否开启Debug模式
	MaxCon    int               // 每个连接的请求数
	HTTP2     bool              // 是否使用http2.0
	Keepalive bool              // 是否开启长连接
	Code      int               //验证的状态码
}

type RequestResults struct {
	Time      time.Duration
	IsSucceed bool
}

//websocket请求的结构体
type WsRequest struct {
	Id   int         `json:"id"`   //消息id
	Ver  string      `json:"ver"`  //版本号
	Path string      `json:"path"` // 请求命令字
	Data interface{} `json:"data"` // 数据 json
}

//websocket返回的结构体
type WsResponse struct {
	Id   int         `json:"id"`   //消息id
	Err  int         `json:"err"`  // 返回的错误码
	Msg  string      `json:"msg"`  // 返回的信息
	Data interface{} `json:"data"` // 返回数据json
}

//生成websocket请求的数据集合
func WsRequestData(id int, ver, path string, data map[string]string) string {
	b, err := json.Marshal(&WsRequest{
		Id:   id,
		Ver:  ver,
		Path: path,
		Data: data,
	})
	if err != nil {
		panic("json Marshal err:" + err.Error())
	}
	return string(b)
}

//json数据转换
func JsonToData(payload string) (*WsResponse, error) {
	var t = &WsResponse{}
	err := json.Unmarshal([]byte(payload), t)
	if err != nil {
		fmt.Println(fmt.Sprintf("JsonToData err: %s", err.Error()))
		return nil, err
	}
	return t, nil
}
