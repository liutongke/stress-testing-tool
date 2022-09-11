package http

import (
	"bytes"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), err
}

var token string

func PostFormData(url string, data map[string]string) (string, error) {
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for k, v := range data {
		_ = writer.WriteField(k, v)
	}
	err := writer.Close()
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return "", err
	}

	req.Header.Add("token", token)

	req.Header.Add("Content-Type", "application/json")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	return string(body), err
}
