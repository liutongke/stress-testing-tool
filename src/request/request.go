package request

import (
	"stress-testing-tool/src/http"
	"strings"
)

// 支持协议
const (
	FormTypeHTTP      = "http"
	FormTypeWebSocket = "webSocket"
	FormTypeGRPC      = "grpc"
)

func NewRequest(userNum, totalUserNum int, url string, keepalive int, postFile string, contentType string) (request *http.Request, err error) {
	var (
		form string
		req  http.Request
	)
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		form = FormTypeHTTP
	} else if strings.HasPrefix(url, "ws://") || strings.HasPrefix(url, "wss://") {
		form = FormTypeWebSocket
	}

	req.Form = form
	req.Method = "GET"
	req.URL = url
	req.Headers = map[string]string{"Content-Type": "application/json"}
	//http.GetHeader("application/x-www-form-urlencoded", "./post.txt", &req)
	//http.GetHeader("application/json", "./post.txt", &req)
	//http.GetHeader("text/plain", "./post.txt", &req)
	//http.GetHeader("multipart/form-data", "./post.txt", &req)

	if postFile != "" && contentType != "" {
		http.GetHeader(contentType, postFile, &req)
	}

	return &req, nil
}
