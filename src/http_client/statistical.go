package http_client

import (
	"fmt"
	"stress-testing-tool/src/tool"
	"time"
)

var exportStatisticsTime = 1 * time.Second

// 时间
var (
	processingTime time.Duration //处理总时间
	maxTime        time.Duration // 最大时长
	minTime        time.Duration // 最小时长
	successNum     uint64        // 成功处理数，code为0
	failureNum     uint64        // 处理失败数，code不为0
	requestNum     int           //已经请求总数量
)

func ReceivingResults(ch <-chan *tool.ResponseRs) {
	var stopChan = make(chan bool)
	startTm := time.Now()

	ticker := time.NewTicker(exportStatisticsTime)
	go tickerEcho(ticker, startTm, stopChan) // 定时输出一次计算结果

	for data := range ch {

		requestNum = requestNum + 1
		//echoProcess(requestNum)

		processingTime += data.RequestTime
		if maxTime <= data.RequestTime {
			maxTime = data.RequestTime
		}

		if minTime == 0 { //第一步需要初始化
			minTime = data.RequestTime
		}

		if minTime > data.RequestTime {
			minTime = data.RequestTime
		}
		// 是否请求成功
		if data.IsSucc == true {
			successNum = successNum + 1
		} else {
			failureNum = failureNum + 1
		}
	}

	stopChan <- true

	runTime := tool.DiffNano(startTm)
	qps := float64(successNum*1e9) / float64(runTime)

	//| 最大请求时长| 最小请求时长 | 成功的处理数 | 失败的请求数 | 处理总时长 | 处理用时 |
	//|-----------|------------|-------------|------------|----------|---------|
	//| %s        | %s         | %d          | %d         | %s       | %s      |

	echoHeader(maxTime, minTime, successNum, failureNum, processingTime, runTime, qps)
	WgHTTPStressTester.Done()
}

// 定时输出一次结果
func tickerEcho(ticker *time.Ticker, startTm time.Time, stopChan <-chan bool) {
	for {
		select {
		case <-ticker.C:
			//endTime := uint64(time.Now().UnixNano())

			qps := float64(successNum*1e9) / float64(tool.DiffNano(startTm))
			echoHeader(maxTime, minTime, successNum, failureNum, processingTime, 0, qps)

		case <-stopChan:
			// 处理完成
			return
		}
	}
}

func echoHeader(maxTime, minTime time.Duration, successNum, failureNum uint64, processingTime, runTime time.Duration, qps float64) {
	//fmt.Println("统计信息")
	//fmt.Println("----------------------")
	fmt.Printf("最大请求时长: %v\n", maxTime)
	fmt.Printf("最小请求时长: %v\n", minTime)
	fmt.Printf("成功的处理数: %d\n", successNum)
	fmt.Printf("失败的请求数: %d\n", failureNum)
	fmt.Printf("处理总时长:   %v\n", processingTime)

	// 如果 runTime 为 0（还未完成或未传值时），就不显示处理用时
	if runTime != 0 {
		fmt.Printf("处理用时:     %v\n", runTime)
	}

	fmt.Printf("QPS:          %.2f\n", qps)
	fmt.Println("----------------------")
}
