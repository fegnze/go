package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var configStruct struct {
	LogLevel    int
	LogFilePath string
}

var (
	//LogLevel 需要记录的日志级别0:error,1:debug,2:info,3:verbose
	LogLevel = configStruct.LogLevel
	//LogFilePath 日志文件的存放路径
	LogFilePath = configStruct.LogFilePath
)

func init() {
	data, err := ioutil.ReadFile("conf/configs.json")
	if err != nil {
		log.Fatalf("conf:%v", err)
	}

	if err = json.Unmarshal(data, &configStruct); err != nil {
		log.Fatalf("conf:%v", err)
	}

	LogLevel = configStruct.LogLevel
	LogFilePath = configStruct.LogFilePath
}
