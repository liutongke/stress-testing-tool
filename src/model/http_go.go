package model

import (
	"fmt"
	"net/http"
	"stress-testing-tool/src/tool"
	"time"
)

var token string

func PostFormData(request *Request) (r *http.Response, requestTime time.Duration, err error) {

	method := request.Method
	url := request.URL
	body := request.Body

	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return
	}

	for k, v := range request.Headers {
		req.Header.Set(k, v)
	}

	req.Close = true //DisableKeepAlives
	fmt.Println(request)
	startTime := time.Now()

	r, err = client.Do(req)

	requestTime = tool.DiffNano(startTime)
	r.Body.Close()

	return
}
