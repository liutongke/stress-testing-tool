package src

import (
	"fmt"
	"time"
)

func ReceivingResults(ch <-chan *ResponseRs) {
	// 时间
	var (
		processingTime time.Duration //处理总时间
		maxTime        time.Duration // 最大时长
		minTime        time.Duration // 最小时长
		successNum     uint64        // 成功处理数，code为0
		failureNum     uint64        // 处理失败数，code不为0
		requestNum     int           //已经请求总数量
	)

	startTm := time.Now()
	for data := range ch {

		requestNum = requestNum + 1
		echoProcess(requestNum)

		processingTime += data.Time
		if maxTime <= data.Time {
			maxTime = data.Time
		}

		if minTime == 0 { //第一步需要初始化
			minTime = data.Time
		}

		if minTime > data.Time {
			minTime = data.Time
		}
		// 是否请求成功
		if data.IsSucceed == true {
			successNum = successNum + 1
		} else {
			failureNum = failureNum + 1
		}
	}

	runTime := DiffNano(startTm)

	//| 最大请求时长| 最小请求时长 | 成功的处理数 | 失败的请求数 | 处理总时长 | 处理用时 |
	//|-----------|------------|-------------|------------|----------|---------|
	//| %s        | %s         | %d          | %d         | %s       | %s      |

	fmt.Printf("|最大请求时长:%s|最小请求时长:%s|成功的处理数:%d|失败的请求数:%d|处理总时长:%s|处理用时:%s|", maxTime, minTime, successNum, failureNum, processingTime, runTime)
	WgTask.Done()
}

// 打印进度
func echoProcess(num int) {
	if (num % (totalUserNum / 10)) == 0 {
		fmt.Printf("Completed %d requests\n", num)
	}
}
