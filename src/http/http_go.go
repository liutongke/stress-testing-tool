package http

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

func PostJosn(url, data string) (string, error) {
	//url := "http://127.0.0.1:12223/auth/login"
	payload := strings.NewReader(data)
	//payload := strings.NewReader(`{“code":"wxdev1111"}`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), err
}

var token string

func PostFormData(url string, data map[string]string) (resp *http.Response, requestTime uint64, err error) {
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for k, v := range data {
		_ = writer.WriteField(k, v)
	}
	err = writer.Close()
	if err != nil {
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return
	}

	req.Header.Add("token", token)

	req.Header.Add("Content-Type", "application/json")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Close = true //DisableKeepAlives
	resp, err = client.Do(req)
	defer resp.Body.Close()
	return
}
