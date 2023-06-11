package http

import (
	"fmt"
	"io"
	"net/http"
	"stress-testing-tool/src/tool"
	"strings"
	"time"
)

var token string

func PostFormData(request *Request) (r *http.Response, requestTime time.Duration, err error) {

	method := request.Method
	url := request.URL

	client := &http.Client{}

	req, err := http.NewRequest(method, url, postPyload(request))

	if err != nil {
		return
	}

	for k, v := range request.Headers {
		req.Header.Set(k, v)
	}

	req.Close = true //DisableKeepAlives

	startTime := time.Now()

	r, err = client.Do(req)

	requestTime = tool.DiffNano(startTime)

	if err != nil {
		return
	}
	r.Body.Close()

	return
}

func HttpDo(userReq *Request) (r *http.Response, requestTime time.Duration, isSucc bool) {
	client := &http.Client{}

	startTime := time.Now()
	
	r, err := client.Do(userReq.Req)

	requestTime = tool.DiffNano(startTime)

	if r.StatusCode != 200 || err != nil {
		isSucc = false
		return
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	fmt.Println(string(body))
	if err != nil {
		return
	}
	return
}

func postPyload(request *Request) io.Reader {
	if strings.Contains(request.Headers["Content-Type"], "multipart/form-data") {
		payload, formDataContentType := GetMultipartFormData([]byte(request.Body))
		request.Headers["Content-Type"] = formDataContentType
		return payload
	}

	return strings.NewReader(request.Body)
}
