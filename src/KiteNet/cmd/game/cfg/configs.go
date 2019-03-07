package cfg

import (
	"KiteNet/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

//httpconf
type httpconf struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

//websocket
type websocket struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

//game
type game struct {
	Name      string    `json:"name"`
	Httpconf  httpconf  `json:"http"`
	Websocket websocket `json:"websocket"`
}

//logConfs
type logConfs struct {
	Path           string        `json:"path"`
	FilePrefixName string        `json:"filePrefixName"`
	FileDuration   time.Duration `json:"fileDuration"`
	SplitDuration  time.Duration `json:"splitDuration"`
}

//Configs
var Configs struct {
	Game game     `json:"server"`
	Log  logConfs `json:"log"`
}

//Parse
func Parse(file string) {
	log.Println("配置文件:", file)
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln("读取配置文件错误:", err)
	}
	err = json.Unmarshal(utils.JsonNormalize(buf), &Configs)
	if err != nil {
		log.Fatalln("解析配置文件错误:", err)
	}
	log.Println("配置:", Configs)
}
