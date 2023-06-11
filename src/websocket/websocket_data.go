package websocket

import (
	"stress-testing-tool/src/http"
	"stress-testing-tool/src/tool"
)

func GetWebsocketData(filePath string, req *http.Request) {
	body, err := tool.GetFileData(filePath)
	if err != nil {
		panic(err)
	}
	req.Body = string(body)
	return
}
