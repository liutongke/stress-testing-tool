package http

import (
	"encoding/json"
	"fmt"
	"stress-testing-tool/src/tool"
)

func GetPostBody(flagParam *FlagParam) map[string]string {
	data, err := tool.GetFileData(flagParam.PostBody)
	if err != nil {
		return nil
	}

	var result map[string]string
	err = json.Unmarshal(data, &result)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}
	return result
}
