package ktlog

var defaultName = "core"

//OpenLogWithCodeAndMsg 开启并写入部分日志
func OpenLogWithCodeAndMsg(code int, msg string) {
	OpenLogWithCodeAndMsgN(defaultName, code, msg)
}

//CloseLog 关闭日志的文件记录
func CloseLog() {
	for name := range logFile {
		CloseLogN(name)
	}
}

//OpenLog 打开日志的文件记录
func OpenLog(level int, file string) {
	OpenLogN(defaultName)
}

//Panic 程序无法处理的错误,中断执行
func Panic(formatStr string, v ...interface{}) {
	PanicN(defaultName, formatStr, v...)
}

//Error 捕获到的逻辑异常,不中断运行
func Error(formatStr string, v ...interface{}) {
	ErrorN(defaultName, formatStr, v...)
}

//Warning 警告,输出格式包含文件名和行号
func Warning(formatStr string, v ...interface{}) {
	WarningN(defaultName, formatStr, v...)
}

//Info 基本的需要记录的日志
func Info(formatStr string, v ...interface{}) {
	InfoN(defaultName, formatStr, v...)
}

//Verbose 额外的不比体现在日志流程的日志
func Verbose(formatStr string, v ...interface{}) {
	VerboseN(defaultName, formatStr, v...)
}
