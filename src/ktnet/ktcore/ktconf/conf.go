package ktconf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//CoreCfgs struct
var CoreCfgs struct {
	LogLevel    int    `json:"logLevel"`
	LogFilePath string `json:"logFilePath"`
}

//CoreCfgs entry
var (
	LogLevel    = CoreCfgs.LogLevel
	LogFilePath = CoreCfgs.LogFilePath
)

//Cfgs map
var Cfgs map[string]interface{}

//Parse 解析配置文件
func Parse(file string) (int, string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return -1, fmt.Sprintf("ConfPareError:%v", err)
	}

	if err = json.Unmarshal(data, &CoreCfgs); err != nil {
		return -1, fmt.Sprintf("ConfPareError:%v", err)
	}

	//set CoreCfgs entry
	LogLevel = CoreCfgs.LogLevel
	LogFilePath = CoreCfgs.LogFilePath

	if err = json.Unmarshal(data, &Cfgs); err != nil {
		return -1, fmt.Sprintf("ConfPareError:%v", err)
	}

	return 0, "Parse configs succesed..."
}
