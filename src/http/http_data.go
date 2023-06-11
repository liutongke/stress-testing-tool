package http

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
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
	Code      int               // 验证的状态码
	Req       *http.Request     //发送的请求信息
}

func GetMultipartFormData(body []byte) (io.Reader, string) {
	var content map[string]interface{}

	err := json.Unmarshal(body, &content)
	if err != nil {
		panic(err)
	}

	payload := &bytes.Buffer{}

	writer := multipart.NewWriter(payload)

	for k, v := range content {
		_ = writer.WriteField(k, v.(string))
	}

	err = writer.Close()
	if err != nil {
		panic(err)
	}

	return payload, writer.FormDataContentType()
}
