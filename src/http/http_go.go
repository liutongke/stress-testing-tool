package http

import (
	"net/http"
	"stress-testing-tool/src/tool"
	"time"
)

var token string

func PostFormData(url string, data map[string]string, request *Request) (response *http.Response, requestTime time.Duration, err error) {

	client := &http.Client{}
	req, err := http.NewRequest(request.Method, url, request.Body)

	if err != nil {
		return
	}

	for k, v := range request.Headers {
		req.Header.Set(k, v)
	}

	req.Close = true //DisableKeepAlives

	startTime := time.Now()

	response, err = client.Do(req)

	requestTime = tool.DiffNano(startTime)
	defer response.Body.Close()

	return
}
