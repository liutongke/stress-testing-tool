package main

import (
	"flag"
	"fmt"
	"math"
	"stress-testing-tool/src/http"
)

var (
	userNum      int //并发数量
	totalUserNum int //总请求次数
	url          string
	keepalive    int
) // 定义几个变量，用于接收命令行的参数值

func inits() {

	//go run main.go -u http -c 100 -n 10000

	// &user 就是接收命令行中输入 -u 后面的参数值，其他同理
	flag.StringVar(&url, "u", "url", "压测地址")
	flag.IntVar(&userNum, "c", 0, "用户数量")
	flag.IntVar(&totalUserNum, "n", 0, "发起请求数量")
	flag.IntVar(&keepalive, "k", 0, "复用连接")
	// 解析命令行参数写入注册的flag里
	flag.Parse()
	// 输出结果
	userRunNum := math.Ceil(float64(totalUserNum / userNum)) //单个用户请求次数
	fmt.Println(userNum, totalUserNum, url, keepalive, userRunNum)
}

func main() {
	var req http.Request
	req.Method = "GET"
	fmt.Println(welcome)
	//src.Run()
	//http.GetHeader("application/x-www-form-urlencoded", "./post.txt", &req)
	//http.GetHeader("application/json", "./post.txt", &req)
	//http.GetHeader("text/plain", "./post.txt", &req)
	//http.GetHeader("multipart/form-data", "./post.txt", &req)
	http.PostFormData("http://192.168.0.105:9500/", map[string]string{"code": "wxdev"}, &req)
}

var welcome = "                               \\\\\\\\\\\\\\\n                            \\\\\\\\\\\\\\\\\\\\\\\\\n                          \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\n  -----------,-|           |C>   // )\\\\\\\\|\n           ,','|          /    || ,'/////|\n---------,','  |         (,    ||   /////\n         ||    |          \\\\  ||||//''''|\n         ||    |           |||||||     _|\n         ||    |______      `````\\____/ \\\n         ||    |     ,|         _/_____/ \\\n         ||  ,'    ,' |        /          |\n         ||,'    ,'   |       |         \\  |\n_________|/    ,'     |      /           | |\n_____________,'      ,',_____|      |    | |\n             |     ,','      |      |    | |\n             |   ,','    ____|_____/    /  |\n             | ,','  __/ |             /   |\n_____________|','   ///_/-------------/   |\n              |===========,'"
