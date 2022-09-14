package http

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"stress-testing-tool/src/tool"
	"strings"
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
}

//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
//req.Header.Add("Content-Type", "application/json")
//req.Header.Add("Content-Type", "text/plain")
//multipart/form-data; boundary=300c39bc6b1b8366edd2ac1835ec4b0bd6daaa98800e305d5a443d224f67

func GetHeader(ContentType, filePath string, req *Request) {
	body, err := tool.GetFileData(filePath)

	req.Method = "POST"

	if err != nil {
		panic("get local body error!")
	}

	if ContentType == "application/x-www-form-urlencoded" {
		getUrlencoded(body, req)
	} else if ContentType == "application/json" {
		getJson(body, req)
	} else if ContentType == "text/plain" {
		getText(body, req)
	} else if ContentType == "multipart/form-data" {
		getFormData(body, req)
	}

}

// req.Header.Add("Content-Type", "application/json")
func getJson(body []byte, req *Request) *Request {
	//payload := strings.NewReader(string(body))
	payload := string(body)
	req.Headers = map[string]string{"Content-Type": "application/json"}
	req.Body = payload
	return req
}

// req.Header.Add("Content-Type", "text/plain")
func getText(body []byte, req *Request) *Request {
	//payload := strings.NewReader(string(body))
	payload := string(body)
	req.Headers = map[string]string{"Content-Type": "text/plain"}
	req.Body = payload
	return req
}

// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
func getUrlencoded(body []byte, req *Request) *Request {

	var content map[string]string

	err := json.Unmarshal(body, &content)
	if err != nil {
		panic(err)
	}

	var list []string

	for k, v := range content { //"keke=123&username=nimo"
		list = append(list, k+"="+v)
	}

	str := strings.Join(list, `&`)

	//payload := strings.NewReader(str)

	req.Headers = map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	req.Body = str
	return req
}

// multipart/form-data; boundary=300c39bc6b1b8366edd2ac1835ec4b0bd6daaa98800e305d5a443d224f67
func getFormData(body []byte, req *Request) *Request {
	_, formDataContentType := GetMultipartFormData(body)
	req.Headers = map[string]string{"Content-Type": formDataContentType}
	req.Body = string(body)
	return req
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
