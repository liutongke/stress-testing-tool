package main

import (
	"flag"
	"fmt"
	"stress-testing-tool/src/http_client"
)

//var (
//	userNum      int    // 并发用户数
//	totalUserNum int    // 请求总次数
//	url          string // 压测的目标地址
//	keepAlive    int    // 是否复用TCP连接，1为开启，0为关闭
//	postBody     string // POST请求的请求体
//	postFile     string // POST请求上传的文件
//	contentType  string // 请求体的Content-Type类型
//	headerFile   string // 额外携带的请求头文件
//)

//func init() {
//	flag.StringVar(&url, "u", "http://example.com", "压测的目标地址。")
//	flag.IntVar(&userNum, "c", 0, "并发用户数。")
//	flag.IntVar(&totalUserNum, "n", 0, "请求总次数。")
//	flag.IntVar(&keepAlive, "k", 0, "是否复用TCP连接：1为开启，0为关闭。")
//	flag.StringVar(&postBody, "p", "", "POST请求的请求体数据（可选）。")
//	flag.StringVar(&postFile, "f", "", "POST请求中需要上传的文件（可选）。")
//	flag.StringVar(&contentType, "t", "application/json", "请求体的Content-Type类型（POST请求需要指定）。")
//	flag.StringVar(&headerFile, "h", "", "携带的额外请求头文件（可选）。")
//
//	flag.Parse()

// 可在此执行其他初始化操作

// 示例：如需计算并打印每个用户的请求次数，去掉以下代码的注释
// userRunNum := math.Ceil(float64(totalUserNum) / float64(userNum))
// fmt.Println("每个用户将大约执行", userRunNum, "次请求。")
//}

func main() {
	cfg := http_client.NewFlagConfig()
	//fmt.Printf("配置参数: %+v\n", cfg)
	if cfg.Requests == 0 || cfg.Concurrency == 0 || cfg.RequestURL == "" {
		fmt.Printf("示例: go run main.go -c 100 -n 10000 -url http://192.168.0.105:9500/ \n")
		fmt.Printf("压测地址必填 \n")
		fmt.Printf("当前请求参数: -c %d -n %d -url %s \n", cfg.Concurrency, cfg.Requests, cfg.RequestURL)
		flag.Usage()
		return
	}

	fmt.Println(welcome)

	http_client.Run(cfg)
	//go run .\main.go -n 10 -c 1 -T raw -H "Accept: application/json" -H "Cache-Control: no-cache" -p ./header.txt -k -url http://192.168.0.110:12349/get -m GET
}

var welcome = "                               \\\\\\\\\\\\\\\n                            \\\\\\\\\\\\\\\\\\\\\\\\\n                          \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\n  -----------,-|           |C>   // )\\\\\\\\|\n           ,','|          /    || ,'/////|\n---------,','  |         (,    ||   /////\n         ||    |          \\\\  ||||//''''|\n         ||    |           |||||||     _|\n         ||    |______      `````\\____/ \\\n         ||    |     ,|         _/_____/ \\\n         ||  ,'    ,' |        /          |\n         ||,'    ,'   |       |         \\  |\n_________|/    ,'     |      /           | |\n_____________,'      ,',_____|      |    | |\n             |     ,','      |      |    | |\n             |   ,','    ____|_____/    /  |\n             | ,','  __/ |             /   |\n_____________|','   ///_/-------------/   |\n              |===========,'"
