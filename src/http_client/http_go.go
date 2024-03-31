package http_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
	"stress-testing-tool/src/tool"
	"strings"
	"sync"
	"time"
)

// RequestGenerator 定义接口
type RequestGenerator interface {
	GenerateRequest(flagParam *ABConfig) (*http.Response, time.Duration, error)
}

// FormDataGenerator 处理 form-data 类型的请求生成
type FormDataGenerator struct{}

func (g FormDataGenerator) GenerateRequest(flagParam *ABConfig) (*http.Response, time.Duration, error) {

	var (
		formData      map[string]string
		filesFormData map[string]string
		err           error
	)

	// 提取JSON文件内容到formData
	if flagParam.Postfile != "" {
		formData, err = jsonFormDataFromFile(flagParam.Postfile)
		if err != nil {
			return nil, 0, err
		}
	}

	// 提取JSON文件内容到filesFormData
	if flagParam.Files != "" {
		filesFormData, err = jsonFormDataFromFile(flagParam.Files)
		if err != nil {
			return nil, 0, err
		}
	}

	// 假设SendMultipartFormData函数可接受文件的map[string]string
	return SendMultipartFormData(flagParam.RequestURL, flagParam.Method, formData, filesFormData, flagParam.Headers)
}

// XWWWFormUrlencodedGenerator 处理 x-www-form-urlencoded 类型的请求生成
type XWWWFormUrlencodedGenerator struct{}

func (g XWWWFormUrlencodedGenerator) GenerateRequest(flagParam *ABConfig) (*http.Response, time.Duration, error) {
	// 从文件读取表单字段并生成x-www-form-urlencoded数据
	var (
		formData map[string]string
		err      error
	)

	// 提取JSON文件内容到formData
	if flagParam.Postfile != "" {
		formData, err = jsonFormDataFromFile(flagParam.Postfile)
		if err != nil {
			return nil, 0, err
		}
	}
	// 假设SendFormURLEncoded函数可接受map[string]string
	return SendFormURLEncoded(flagParam.RequestURL, flagParam.Method, formData, flagParam.Headers)
}

// jsonFormDataFromFile 读取JSON文件，并转换为表单数据
func jsonFormDataFromFile(filePath string) (map[string]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read from file %s: %w", filePath, err)
	}

	// 解析JSON文件内容到map中
	var formData map[string]string
	err = json.Unmarshal(data, &formData)
	if err != nil {
		return nil, fmt.Errorf("unable to parse JSON file %s: %w", filePath, err)
	}

	return formData, nil
}

// RawDataGenerator 处理 raw 类型的请求生成
type RawDataGenerator struct{}

func (g RawDataGenerator) GenerateRequest(flagParam *ABConfig) (*http.Response, time.Duration, error) {
	contentType := getContentType(flagParam.Headers, "application/json")

	requestBody, err := os.ReadFile(flagParam.Postfile)
	if err != nil {
		return nil, 0, fmt.Errorf("error reading request body file: %w", err)
	}

	return SendRaw(flagParam.RequestURL, flagParam.Method, contentType, requestBody, flagParam.Headers)
}

// BinaryDataGenerator 处理 binary 类型的请求生成
type BinaryDataGenerator struct{}

func (g BinaryDataGenerator) GenerateRequest(flagParam *ABConfig) (*http.Response, time.Duration, error) {
	requestBody, err := os.ReadFile(flagParam.Files)

	if err != nil {
		return nil, 0, fmt.Errorf("error reading request body file: %w", err)
	}

	return SendBinary(flagParam.RequestURL, flagParam.Method, "application/octet-stream", requestBody, flagParam.Headers)
}

// CreateGenerator 根据 ContentType 创建对应的 RequestGenerator
func CreateGenerator(contentType string) (RequestGenerator, error) {
	switch contentType {
	case "form-data":
		return &FormDataGenerator{}, nil
	case "x-www-form-urlencoded":
		return &XWWWFormUrlencodedGenerator{}, nil
	case "raw":
		return &RawDataGenerator{}, nil
	case "binary":
		return &BinaryDataGenerator{}, nil
	default:
		return nil, errors.New("unsupported content type")
	}
}

// ProcessRequest 使用 CreateGenerator 函数来处理请求
func ProcessRequest(flagParam *ABConfig) (*http.Response, time.Duration, error) {
	generator, err := CreateGenerator(flagParam.ContentType)
	if err != nil {
		return nil, 0, err
	}
	return generator.GenerateRequest(flagParam)
}

// http请求提前登录
func StartHTTPRequest(WgHTTPRequest *sync.WaitGroup, ch chan<- *tool.ResponseRs, flagParam *ABConfig) {
	defer WgHTTPRequest.Done()
	userRunNum := int(math.Ceil(float64(flagParam.Requests / flagParam.Concurrency)))
	for i := 0; i < userRunNum; i++ {
		ch <- send(flagParam)
	}
}

func send(flagParam *ABConfig) *tool.ResponseRs {
	return DoResponse(ProcessRequest(flagParam))
}

// getContentType 获取或设置 'Content-Type' 的函数
func getContentType(headers map[string]string, defaultValue string) string {
	for k, v := range headers {
		if strings.EqualFold(k, "Content-Type") {
			return v
		}
	}
	return defaultValue
}
