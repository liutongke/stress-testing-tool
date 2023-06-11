package http

import (
	"net/http"
	"net/url"
	"strings"
)

func generatePayload(flagParam *FlagParam) *strings.Reader {
	// 创建一个空的 url.Values 对象
	values := url.Values{}

	// 添加键值对数据
	body := GetPostBody(flagParam)
	for k, v := range body {
		values.Set(k, v)
	}

	//values.Set("num", "99999999")
	//values.Set("nick", "generatePayload")

	// 将 url.Values 对象转换为字符串
	queryString := values.Encode()

	//fmt.Println(queryString) // 输出: num=1111111111111&nick=keke
	return strings.NewReader(queryString)
}

func setHeader(req *http.Request, userReq *Request) {
	for k, v := range userReq.Headers {
		req.Header.Add(k, v)
	}
	//req.Header.Add("x-token", "b69db1d710f429675433bc0ce3135a47")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}

func StartXWWWFormUrlencoded(userReq *Request, flagParam *FlagParam) (err error) {
	payload := generatePayload(flagParam)

	req, err := http.NewRequest(userReq.Method, userReq.URL, payload)

	if err != nil {
		return
	}

	setHeader(req, userReq)

	userReq.Req = req
	//isSucc, body = do(req)
	return
}
