package process

import (
	"stress-testing-tool/src/tool"
	"sync"
)

func Start(userRunNum int, WgUser *sync.WaitGroup, ch chan<- *tool.ResponseRs, postFile string) {
	//defer func() {
	//	WgUser.Done()
	//}()
	//data, err := tool.GetFileData(postFile)
	//if err != nil {
	//	return
	//}
	//
	//for _, v := range string(data) {
	//
	//}
	//lastName := gjson.Get(data, "name.last")
	//for i := 0; i < userRunNum; i++ {
	//
	//	http.PostFormData(&http.Request{
	//		URL:       "",
	//		Form:      "",
	//		Method:    "",
	//		Headers:   nil,
	//		Body:      "",
	//		Verify:    "",
	//		Timeout:   0,
	//		Debug:     false,
	//		MaxCon:    0,
	//		HTTP2:     false,
	//		Keepalive: false,
	//		Code:      0,
	//	})
	//	ch <- &tool.ResponseRs{
	//		IsSucc:      isSucc,
	//		DataLen:     dataLen,
	//		RequestTime: requestTime,
	//	}

	//}
}
