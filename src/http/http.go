package http

import (
	"fmt"
	"stress-testing-tool/src/tool"
	"sync"
	"time"
)

// http请求提前登录
func Http(userRunNum int, url string, WgUser *sync.WaitGroup, ch chan<- *tool.ResponseRs) {

	defer func() {
		WgUser.Done()
	}()

	for i := 0; i < userRunNum; i++ {

		startTime := time.Now()

		_, _, err := PostFormData(url, map[string]string{"code": fmt.Sprintf("wxdev%d", i)})

		if err != nil {
			//fmt.Println(err)
			//os.Exit(0)
			ch <- &tool.ResponseRs{
				Time:      tool.DiffNano(startTime),
				IsSucceed: false,
			}
		} else {
			ch <- &tool.ResponseRs{
				Time:      tool.DiffNano(startTime),
				IsSucceed: true,
			}
		}

	}
}
