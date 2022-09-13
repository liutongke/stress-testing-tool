package request

import (
	"stress-testing-tool/src/model"
)

func NewRequest(userNum, totalUserNum int, url string, keepalive int, postFile string, contentType string) (request *model.Request, err error) {
	var req model.Request

	req.Method = "GET"
	req.URL = url
	req.Headers = map[string]string{"Content-Type": "application/json"}
	//http.GetHeader("application/x-www-form-urlencoded", "./post.txt", &req)
	//http.GetHeader("application/json", "./post.txt", &req)
	//http.GetHeader("text/plain", "./post.txt", &req)
	//http.GetHeader("multipart/form-data", "./post.txt", &req)

	if postFile != "" && contentType != "" {
		model.GetHeader(contentType, postFile, &req)
	}

	return &req, nil
}
