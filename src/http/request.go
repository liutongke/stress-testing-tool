package http

import (
	"strings"
)

// 支持协议
const (
	FormTypeHTTP      = "http"
	FormTypeWebSocket = "webSocket"
	FormTypeProcess   = "process"
	FormTypeGRPC      = "grpc"
)

// FlagParam 命令行数据
type FlagParam struct {
	Url          string
	UserNum      int
	TotalUserNum int
	Keepalive    int
	PostBody     string
	PostFile     string
	ContentType  string
}

func NewRequest(flagParam *FlagParam) (userReq *Request, err error) {

	form := getRequestType(flagParam.Url)

	method := "POST"
	if flagParam.PostBody == "" {
		method = "GET"
	}

	userReq = &Request{
		URL:       flagParam.Url,
		Form:      form,
		Method:    method,
		Headers:   map[string]string{"Content-Type": "application/json"},
		Keepalive: flagParam.Keepalive != 1, //true关闭 false开启
	}
	//http.GetHeader("application/x-www-form-urlencoded", "./post.txt", &req)
	//http.GetHeader("application/json", "./post.txt", &req)
	//http.GetHeader("text/plain", "./post.txt", &req)
	//http.GetHeader("multipart/form-data", "./post.txt", &req)

	switch flagParam.ContentType {
	case "application/x-www-form-urlencoded":
		err = StartXWWWFormUrlencoded(userReq, flagParam)

	case "application/json":
		err = StartFormData(userReq, flagParam)
	case "text/plain":
		//getText(body, req)
	case "multipart/form-data":
		//getFormData(body, req)
	default:
		// 处理未知的 ContentType
	}
	//if form == FormTypeWebSocket {
	//	websocket.GetWebsocketData(postFile, req)
	//}
	//if postFile != "" && contentType != "" {
	//	if form == FormTypeHTTP {
	//		http.GetHeader(contentType, postFile, req)
	//	}
	//	if form == FormTypeWebSocket {
	//		websocket.GetWebsocketData(postFile, req)
	//	}
	//}

	return
}

// 获取请求类型
func getRequestType(url string) string {
	var form string
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		form = FormTypeHTTP
	} else if strings.HasPrefix(url, "ws://") || strings.HasPrefix(url, "wss://") {
		form = FormTypeWebSocket
	}
	return form
}
