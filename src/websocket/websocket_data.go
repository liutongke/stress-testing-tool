package websocket

import (
	"stress-testing-tool/src/tool"
	"stress-testing-tool/tmp"
)

func GetWebsocketData(filePath string, req *main.Request) {
	body, err := tool.GetFileData(filePath)
	if err != nil {
		panic(err)
	}
	req.Body = string(body)
	return
}
