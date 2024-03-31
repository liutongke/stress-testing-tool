package http_client

import (
	"math"
	"net/http"
	"stress-testing-tool/src/tool"
	"sync"
	"time"
)

// 定义 RequestGenerator 接口
type RequestGenerator interface {
	GenerateRequest(flagParam *ABConfig) (*http.Response, time.Duration, error)
}

// FormDataGenerator 处理 "form-data" 类型的请求生成
type FormDataGenerator struct{}

func (g FormDataGenerator) GenerateRequest(flagParam *ABConfig) (*http.Response, time.Duration, error) {
	// 实现 form-data 请求生成逻辑...
	// 自定义请求头
	//customHeaders := map[string]string{
	//	"X-Custom-Header": "my-custom-value",
	//}
	// 示例调用
	return SendMultipartFormData(flagParam.RequestURL, flagParam.Method, map[string]string{"key": "FormData"}, map[string]string{"file": "./example.txt"}, flagParam.Headers)

}

// XWWWFormUrlencodedGenerator 处理 "x-www-form-urlencoded" 类型的请求生成
type XWWWFormUrlencodedGenerator struct{}

func (g XWWWFormUrlencodedGenerator) GenerateRequest(flagParam *ABConfig) (*http.Response, time.Duration, error) {
	// 实现 x-www-form-urlencoded 请求生成逻辑...
	//http_client.PrintResponse(http_client.SendFormURLEncoded("http://192.168.0.110:12349/form", url.Values{"key": []string{"FormURLEncoded-keke"}, "key1": []string{"key1"}}, customHeaders))
	return SendFormURLEncoded(flagParam.RequestURL, flagParam.Method, map[string]string{"key": "FormURLEncoded-test", "key1": "test"}, flagParam.Headers)

}

// RawDataGenerator 处理 "raw" 类型的请求生成
type RawDataGenerator struct{}

func (g RawDataGenerator) GenerateRequest(flagParam *ABConfig) (*http.Response, time.Duration, error) {
	// 实现 raw 数据请求生成逻辑...
	return SendRaw(flagParam.RequestURL, flagParam.Method, "application/json", []byte(`{"key": "raw"}`), flagParam.Headers)

}

// BinaryDataGenerator 处理 "binary" 类型的请求生成
type BinaryDataGenerator struct{}

func (g BinaryDataGenerator) GenerateRequest(flagParam *ABConfig) (*http.Response, time.Duration, error) {
	// 实现 binary 数据请求生成逻辑...
	return SendBinary(flagParam.RequestURL, flagParam.Method, "application/octet-stream", []byte{0x00, 0x01, 0x02}, flagParam.Headers)
}

// CreateGenerator 根据 ContentType 创建对应的 RequestGenerator
func CreateGenerator(contentType string) RequestGenerator {
	switch contentType {
	case "form-data":
		return FormDataGenerator{}
	case "x-www-form-urlencoded":
		return XWWWFormUrlencodedGenerator{}
	case "raw":
		return RawDataGenerator{}
	case "binary":
		return BinaryDataGenerator{}
	default:
		// 处理未知的 ContentType 或返回错误信息
	}
	return nil
}

// 使用 CreateGenerator 函数来处理请求
func ProcessRequest(flagParam *ABConfig) (*http.Response, time.Duration, error) {
	generator := CreateGenerator(flagParam.ContentType)
	if generator != nil {
		return generator.GenerateRequest(flagParam)
	}
	// 处理没有对应处理器的情况，或返回错误信息
	return nil, 0, nil
}

// http请求提前登录
func StartHTTPRequest(WgHTTPRequest *sync.WaitGroup, ch chan<- *tool.ResponseRs, flagParam *ABConfig) {

	defer func() {
		WgHTTPRequest.Done()
	}()
	userRunNum := int(math.Ceil(float64(flagParam.Requests / flagParam.Concurrency))) //每个用户发送的请求次数
	for i := 0; i < userRunNum; i++ {
		ch <- send(flagParam)

		//time.Sleep(requestInterval)
		//time.Sleep(1000 * time.Millisecond) // 等待一下，然后再发起下一个请求
	}
}

func send(flagParam *ABConfig) *tool.ResponseRs {
	return DoResponse(ProcessRequest(flagParam))
}
