package http

import (
	"net/http"
	"stress-testing-tool/src/tool"
	"time"
)

func HttpDo(userReq *Request, flagParam *FlagParam) *tool.ResponseRs {
	req := myRequest(userReq, flagParam)
	client := &http.Client{}

	startTime := time.Now()

	r, err := client.Do(req)

	requestTime := tool.DiffNano(startTime)

	if err != nil {
		return &tool.ResponseRs{
			IsSucc:      false,
			DataLen:     0,
			RequestTime: requestTime,
		}
	}
	defer r.Body.Close()
	//body, err := io.ReadAll(r.Body)
	//fmt.Println(string(body))
	if r.StatusCode != 200 {
		return &tool.ResponseRs{
			IsSucc:      false,
			DataLen:     0,
			RequestTime: requestTime,
		}
	}

	return &tool.ResponseRs{
		IsSucc:      true,
		DataLen:     0,
		RequestTime: requestTime,
	}
}

func myRequest(userReq *Request, flagParam *FlagParam) (myReq *http.Request) {
	switch flagParam.ContentType {
	case "application/x-www-form-urlencoded":
		myReq = StartXWWWFormUrlencoded(userReq, flagParam)
	case "application/json":
		myReq = StartFormData(userReq, flagParam)
	case "text/plain":
		//getText(body, req)
	case "multipart/form-data":
		//getFormData(body, req)
	default:
		// 处理未知的 ContentType
	}
	return
}
