package src

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

	fmt.Println("-------success-------")
	echoHeader(maxTime, minTime, successNum, failureNum, processingTime, runTime, qps)
	fmt.Println("-------end-------")
	WgTask.Done()
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

// 打印进度
func echoProcess(num int) {
	if (num % (totalUserNum / 10)) == 0 {
		fmt.Printf("Completed %d requests\n", num)
	}
}

func echoHeader(maxTime, minTime time.Duration, successNum, failureNum uint64, processingTime, runTime time.Duration, qps float64) {
	//| 最大请求时长| 最小请求时长 | 成功的处理数 | 失败的请求数 | 处理总时长 | 处理用时 |
	//|-----------|------------|-------------|------------|----------|---------|
	fmt.Printf("最大请求时长:%s 最小请求时长:%s 成功的处理数:%d 失败的请求数:%d 处理总时长:%s 处理用时:%s qps:%d\n", maxTime, minTime, successNum, failureNum, processingTime, runTime, int(qps))
	fmt.Println("=============================================================================================================================")
	return
}
