package ktcore

import (
	"ktnet/ktcore/ktconf"
	"ktnet/ktcore/ktlog"
)

//InitLog ...
func InitLog(code int, msg string) {
	ktlog.Init(ktconf.LogLevel, ktconf.LogFilePath)
	ktlog.OpenLogWithCodeAndMsg(code, msg)
}

//InitConf for configs
func InitConf(file string) (int, string) {
	return ktconf.Parse(file)
}

//Init ...
func Init(configFile string) {
	InitLog(InitConf(configFile))
}
