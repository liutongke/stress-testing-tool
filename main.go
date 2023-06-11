package main

//func main() {
//	fmt.Println(welcome)
//	//http.StartFormData("./f03293d19752707e9fc91e93061be8f0.jpg")
//	http.StartXWWWFormUrlencoded()
//}

import (
	"flag"
	"fmt"
	"stress-testing-tool/src"
	"stress-testing-tool/src/http"
)

var (
	userNum      int    //并发数量
	totalUserNum int    //总请求次数
	url          string //请求地址
	keepalive    int    //是否复用TCP 1开启0关闭
	postBody     string //post请求体
	postFile     string //post上传的文件
	contentType  string //请求体类型

	headerFile string //携带的请求头
) // 定义几个变量，用于接收命令行的参数值

func init() {
	//go run main.go -u http -c 100 -n 10000
	// &user 就是接收命令行中输入 -u 后面的参数值，其他同理
	flag.StringVar(&url, "u", "url", "压测地址")
	flag.IntVar(&userNum, "c", 0, "用户数量")
	flag.IntVar(&totalUserNum, "n", 0, "发起请求数量")
	flag.IntVar(&keepalive, "k", 0, "复用连接")
	flag.StringVar(&postBody, "p", "", "选填：postBody，发送POST请求体数据")
	flag.StringVar(&postFile, "f", "", "选填：postfile，发送POST请求时需要上传的文件")
	flag.StringVar(&contentType, "t", "", "即content-type，用于设置Content-Type请求头信息,post请求必选项")
	flag.StringVar(&headerFile, "h", "", "携带的请求头")
	// 解析命令行参数写入注册的flag里
	flag.Parse()
	// 输出结果
	//userRunNum := math.Ceil(float64(totalUserNum / userNum)) //单个用户请求次数
	//fmt.Println(userNum, totalUserNum, url, keepalive, userRunNum)
}

func main() {

	if userNum == 0 || totalUserNum == 0 || url == "" {
		fmt.Printf("示例: go run main.go -c 100 -n 10000 -u http://192.168.0.105:9500/ \n")
		fmt.Printf("压测地址必填 \n")
		fmt.Printf("当前请求参数: -c %d -n %d -u %s \n", userNum, totalUserNum, url)
		flag.Usage()
		return
	}

	flagParam := &http.FlagParam{
		Url:          url,
		UserNum:      userNum,
		TotalUserNum: totalUserNum,
		Keepalive:    keepalive,
		PostBody:     postBody,
		PostFile:     postFile,
		ContentType:  contentType,
		HeaderFile:   headerFile,
	}

	userReq, err := http.NewRequest(flagParam)

	fmt.Println(welcome)

	if err != nil {
		panic("初始化数据失败")
	}

	src.Run(userReq, flagParam)
	//go run main.go -c 10 -n 1000 -u http://192.168.0.105:9500/ -t application/x-www-form-urlencoded -p C:\Users\keke\dev\docker\stress-testing-tool/post.txt
}

var welcome = "                               \\\\\\\\\\\\\\\n                            \\\\\\\\\\\\\\\\\\\\\\\\\n                          \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\n  -----------,-|           |C>   // )\\\\\\\\|\n           ,','|          /    || ,'/////|\n---------,','  |         (,    ||   /////\n         ||    |          \\\\  ||||//''''|\n         ||    |           |||||||     _|\n         ||    |______      `````\\____/ \\\n         ||    |     ,|         _/_____/ \\\n         ||  ,'    ,' |        /          |\n         ||,'    ,'   |       |         \\  |\n_________|/    ,'     |      /           | |\n_____________,'      ,',_____|      |    | |\n             |     ,','      |      |    | |\n             |   ,','    ____|_____/    /  |\n             | ,','  __/ |             /   |\n_____________|','   ///_/-------------/   |\n              |===========,'"
