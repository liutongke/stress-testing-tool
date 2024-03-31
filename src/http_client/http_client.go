package http_client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"stress-testing-tool/src/tool"
	"strings"
	"time"
)

// 发送请求的通用函数
//
//	func sendRequest(method, url string, contentType string, body io.Reader, headers map[string]string) (*http.Response, error) {
//		request, err := http.NewRequest(method, url, body)
//		if err != nil {
//			return nil, err
//		}
//
//		// 设置内容类型
//		if contentType != "" {
//			request.Header.Set("Content-Type", contentType)
//		}
//
//		// 设置额外的请求头
//		for key, value := range headers {
//			request.Header.Set(key, value)
//		}
//
//		client := &http.Client{}
//
//		// 请求前时间
//		startTime := time.Now()
//
//		// 发送请求
//		response, err := client.Do(request)
//
//		// 请求后时间
//		endTime := time.Now()
//
//		// 请求耗时统计
//		requestDuration := endTime.Sub(startTime)
//		fmt.Printf("请求花费的时间: %v\n", requestDuration)
//
//		// 将耗时转换为浮点数形式的毫秒
//		durationInMilliseconds := float64(requestDuration) / float64(time.Millisecond)
//		// 使用 fmt.Sprintf 保留两位小数
//		formattedDuration := fmt.Sprintf("%.4f", durationInMilliseconds)
//		fmt.Printf("请求耗时：%s 毫秒\n", formattedDuration)
//
//		return response, err
//	}
func sendRequest(method, requestURL string, contentType string, body io.Reader, headers map[string]string) (*http.Response, time.Duration, error) {
	request, err := http.NewRequest(method, requestURL, body)
	if err != nil {
		return nil, 0, err
	}

	// 设置内容类型
	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}

	// 设置额外的请求头
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{}

	// 请求前时间
	startTime := time.Now()

	// 发送请求
	response, err := client.Do(request)

	// 请求后时间
	endTime := time.Now()

	// 请求耗时统计，计算毫秒并格式化为两位小数
	requestDuration := endTime.Sub(startTime)
	//fmt.Println(requestDuration)
	//durationInMilliseconds := float64(requestDuration) / float64(time.Millisecond)
	//formattedDuration := fmt.Sprintf("%.4f", durationInMilliseconds)

	// 如果发生错误，返回错误并包含耗时信息
	if err != nil {
		return nil, requestDuration, err
	}

	// 请求成功完成，返回响应和耗时信息
	return response, requestDuration, nil
}

// 创建并发送 multipart/form-data 请求
func SendMultipartFormData(requestURL, method string, fields map[string]string, files map[string]string, headers map[string]string) (*http.Response, time.Duration, error) {
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 添加表单字段
	for field, value := range fields {
		if err := writer.WriteField(field, value); err != nil {
			return nil, 0, err
		}
	}

	// 添加文件
	for fieldname, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			return nil, 0, err
		}
		defer file.Close()

		part, err := writer.CreateFormFile(fieldname, filename)
		if err != nil {
			return nil, 0, err
		}
		if _, err = io.Copy(part, file); err != nil {
			return nil, 0, err
		}
	}

	contentType := writer.FormDataContentType()
	if err := writer.Close(); err != nil {
		return nil, 0, err
	}

	return sendRequest(method, requestURL, contentType, &requestBody, headers)
}

// 创建并发送 application/x-www-form-urlencoded 请求
func SendFormURLEncoded(requestURL, method string, fields map[string]string, headers map[string]string) (*http.Response, time.Duration, error) {
	data := url.Values{}

	for k, v := range fields {
		data.Set(k, v)
	}

	return sendRequest(method, requestURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()), headers)
}

// 创建并发送 raw 请求 (如 JSON)
func SendRaw(requestURL, method, contentType string, rawData []byte, headers map[string]string) (*http.Response, time.Duration, error) {
	return sendRequest(method, requestURL, contentType, bytes.NewReader(rawData), headers)
}

// 创建并发送 binary 请求
func SendBinary(requestURL, method, contentType string, binaryData []byte, headers map[string]string) (*http.Response, time.Duration, error) {
	return sendRequest(method, requestURL, contentType, bytes.NewReader(binaryData), headers)
}

// 打印响应内容
//
//	func PrintResponse(resp *http.Response, err error) {
//		if err != nil {
//			fmt.Println("Error making request:", err)
//			return
//		}
//		defer resp.Body.Close()
//
//		fmt.Printf("HTTP Status Code: %d\n", resp.StatusCode)
//		bodyBytes, err := ioutil.ReadAll(resp.Body)
//		if err != nil {
//			fmt.Println("Error reading response body:", err)
//			return
//		}
//		fmt.Printf("Response Body:\n%s\n", string(bodyBytes))
//	}
//
// 打印响应内容和请求耗时
func PrintResponse(resp *http.Response, duration time.Duration, err error) {
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// 打印请求耗时和 HTTP 状态码
	durationInMilliseconds := float64(duration) / float64(time.Millisecond)
	formattedDuration := fmt.Sprintf("%.4f", durationInMilliseconds)
	fmt.Printf("请求耗时：%s 毫秒\n", formattedDuration)
	//fmt.Printf("请求耗时：%s 毫秒\n", duration)
	fmt.Printf("HTTP Status Code: %d\n", resp.StatusCode)

	// 读取并打印响应内容
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Printf("Response Body:\n%s\n", string(bodyBytes))
}
func DoResponse(resp *http.Response, duration time.Duration, err error) *tool.ResponseRs {
	if err != nil {
		return &tool.ResponseRs{
			IsSucc:      false,
			DataLen:     0,
			Body:        fmt.Sprintf("Error making request:%v", err),
			RequestTime: duration,
		}
	}
	defer resp.Body.Close()

	//fmt.Printf("HTTP Status Code: %d\n", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return &tool.ResponseRs{
			IsSucc:      false,
			DataLen:     0,
			Body:        "",
			RequestTime: duration,
		}
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &tool.ResponseRs{
			IsSucc:      false,
			DataLen:     0,
			Body:        fmt.Sprintf("Error reading response body:%v", err),
			RequestTime: duration,
		}
	}

	return &tool.ResponseRs{
		IsSucc:      true,
		DataLen:     0,
		Body:        string(bodyBytes),
		RequestTime: duration,
	}
}
