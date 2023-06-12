package http

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type FormData struct {
}

func NewFormData() *FormData {
	return &FormData{}
}

func (f *FormData) generatePayload(flagParam *FlagParam) (*bytes.Buffer, *multipart.Writer) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	if flagParam.PostFile != "" { //需要上传文件

		file, errFile2 := os.Open(flagParam.PostFile)
		defer file.Close()

		part2, errFile2 := writer.CreateFormFile("file", filepath.Base(flagParam.PostFile))
		_, errFile2 = io.Copy(part2, file)

		if errFile2 != nil {
			fmt.Println(errFile2)
			return nil, nil
		}
	}

	if flagParam.PostBody != "" {
		body := GetPostBody(flagParam)

		for k, v := range body {
			_ = writer.WriteField(k, v)
		}
		//_ = writer.WriteField("num", "2")
	}
	err := writer.Close()

	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return payload, writer
}

func (f *FormData) setHeader(req *http.Request, userReq *Request, writer *multipart.Writer) {
	for k, v := range userReq.Headers {
		req.Header.Add(k, v)
	}
	//req.Header.Add("x-token", "5656565656")
	req.Header.Set("Content-Type", writer.FormDataContentType())
}

func (f *FormData) GenerateRequest(userReq *Request, flagParam *FlagParam) *http.Request {
	payload, writer := f.generatePayload(flagParam)

	req, err := http.NewRequest(userReq.Method, userReq.URL, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	f.setHeader(req, userReq, writer)

	//isSucc, body := do(req)
	//fmt.Println(isSucc, string(body))
	return req
}
