package config

import (
	"time"
)

//db
type db struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

//staticFileServe
type fs struct {
	LocalPath string `json:"local_path"`
	Prefix    string `json:"prefix"`
}

//web
type web struct {
	Name string `json:"name"`
	Ip   string `json:"ip"`
	Port string `json:"port"`
	DB   db     `json:"db"`
	FS   fs     `json:"statices"`
}

//world
type world struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

//log
type log struct {
	Path           string        `json:"path"`
	FilePrefixName string        `json:"filePrefixName"`
	FileDuration   time.Duration `json:"fileDuration"`
	SplitDuration  time.Duration `json:"splitDuration"`
}

//Configs 配置文件解析结构
type Configs struct {
	Web   web   `json:"web"`
	World world `json:"world"`
	Log   log   `json:"log"`
}
