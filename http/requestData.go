package http

import (
	"fmt"
	"github.com/tidwall/gjson"
	"stress-testing-tool"
)

// http请求提前登录
func ReadyHttp() {
	for i := 0; i < concurrence; i++ { //单独先生成减轻数据库压力
		url := fmt.Sprintf("http://%s/auth/login", main.ip)
		//payload := fmt.Sprintf("{\"code\":\"wxdev%d\"}", i)
		//fmt.Println(payload)
		//data, err := PostJosn(url, payload)
		data, err := PostFormData(url, map[string]string{"code": fmt.Sprintf("wxdev%d", i)})
		if err != nil {
			fmt.Println(fmt.Sprintf("post json ------>%d err:%s", i, err.Error()))
		} else {
			fmt.Println(fmt.Sprintf("post json succ --------->%d payload:%s", i, data))
		}
		token = gjson.Get(data, "data.token").String()

		createJob, _ := PostFormData(fmt.Sprintf("http://%s/user/job/create", main.ip), map[string]string{"job": "1", "jobCardId": "1"})
		fmt.Println(createJob)
		jobId := gjson.Get(createJob, "data.jobId").String()
		if len(jobId) <= 0 {
			fmt.Println(fmt.Sprintf("post json jobId err --------->%s", jobId))
		}
		fmt.Println(webSocketUrl(main.ip, token, jobId))
		main.urlList[i] = webSocketUrl(main.ip, token, jobId)
	}
}
