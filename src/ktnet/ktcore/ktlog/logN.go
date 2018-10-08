package ktlog

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

//log level
const (
	LvPanic = -1
	LvERROR = iota
	LvWARNING
	LvINFO
	LvVERBOSE
)

const (
	prefixERROR   = "[ERROR#!!]"
	prefixWarning = "[#Warning#]"
	prefixINFO    = "[INFO]"
	prefixVERBOSE = "[VERBOSE]"
)

var logLevel int
var logFilePath string
var logFile = make(map[string]*os.File)
var logger = make(map[string]*log.Logger)

//Init ...
func Init(level int, path string) {
	logLevel = level
	logFilePath = path
}

//PanicN 程序无法处理的错误,中断执行
func PanicN(name string, formatStr string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	callpos := " \tcallpos=>(" + file + ":" + strconv.Itoa(line) + ")"

	logger[name].SetPrefix(prefixERROR)
	if len(v) > 0 {
		if strings.Index(formatStr, "%") > -1 {
			str := fmt.Sprintf(formatStr, v...) + callpos
			logger[name].Panicln(str)
			fmt.Println(str)
		} else {
			str := fmt.Sprint(v...)
			str = formatStr + "," + str + callpos
			logger[name].Panicln(str)
			fmt.Println(str)
		}

	} else {
		logger[name].Panicln(formatStr + callpos)
		fmt.Println(formatStr + callpos)
	}
}

//ErrorN 捕获到的逻辑异常,不中断运行
func ErrorN(name string, formatStr string, v ...interface{}) {
	if logLevel < LvERROR {
		return
	}

	_, file, line, _ := runtime.Caller(1)
	callpos := " \tcallpos=>(" + file + ":" + strconv.Itoa(line) + ")"

	logger[name].SetPrefix(prefixERROR)
	if len(v) > 0 {
		if strings.Index(formatStr, "%") > -1 {
			str := fmt.Sprintf(formatStr, v...) + callpos
			logger[name].Println(str)
			fmt.Println(str)
		} else {
			str := fmt.Sprint(v...)
			str = formatStr + "," + str + callpos
			logger[name].Println(str)
			fmt.Println(str)
		}

	} else {
		logger[name].Println(formatStr + callpos)
		fmt.Println(formatStr + callpos)
	}
}

//WarningN 警告,输出格式包含文件名和行号
func WarningN(name string, formatStr string, v ...interface{}) {
	if logLevel < LvWARNING {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	callpos := " \tcallpos=>(" + file + ":" + strconv.Itoa(line) + ")"

	logger[name].SetPrefix(prefixWarning)
	if len(v) > 0 {
		if strings.Index(formatStr, "%") > -1 {
			str := fmt.Sprintf(formatStr, v...) + callpos
			logger[name].Println(str)
			fmt.Println(str)
		} else {
			str := fmt.Sprint(v...)
			str = formatStr + "," + str + callpos
			logger[name].Println(str)
			fmt.Println(str)
		}

	} else {
		logger[name].Println(formatStr + callpos)
		fmt.Println(formatStr + callpos)
		fmt.Println(formatStr + callpos)
	}
}

//InfoN 基本的需要记录的日志
func InfoN(name string, formatStr string, v ...interface{}) {
	if logLevel < LvINFO {
		return
	}
	logger[name].SetPrefix(prefixINFO)
	if len(v) > 0 {
		if strings.Index(formatStr, "%") > -1 {
			str := fmt.Sprintf(formatStr, v...)
			logger[name].Println(str)
		} else {
			str := fmt.Sprint(v...)
			str = formatStr + "," + str
			logger[name].Println(str)
			fmt.Println(str)
		}

	} else {
		logger[name].Println(formatStr)
		fmt.Println(formatStr)
	}
}

//VerboseN 额外的不比体现在日志流程的日志
func VerboseN(name string, formatStr string, v ...interface{}) {
	if logLevel < LvVERBOSE {
		return
	}
	logger[name].SetPrefix(prefixVERBOSE)
	if len(v) > 0 {
		if strings.Index(formatStr, "%") > -1 {
			str := fmt.Sprintf(formatStr, v...)
			logger[name].Println(str)
			fmt.Println(str)
		} else {
			str := fmt.Sprint(v...)
			str = formatStr + "," + str
			logger[name].Println(str)
			fmt.Println(str)
		}

	} else {
		logger[name].Println(formatStr)
	}
}

//CloseLogN 关闭日志的文件记录
func CloseLogN(name string) {
	if logFile[name] != nil {
		fmt.Println("close log file ... ")
		logFile[name].Close()
	}
}

//OpenLogN 打开日志的文件记录
func OpenLogN(name string) {
	if logLevel <= 0 {
		logLevel = LvVERBOSE
	}

	if logFilePath == "" {
		logFilePath = "./log/"
	}

	if logger[name] != nil {
		Error(name, ""+name+" logger is already exist ...")
		return
	}
	var err error
	if _, err1 := os.Stat(logFilePath + name + ".log"); os.IsNotExist(err1) {
		logFile[name], err = os.Create(logFilePath + name + ".log")
		if err != nil {
			log.Fatalf("%v\n,open log file err ... %s", err, logFilePath)
		}
	} else {
		logFile[name], err = os.OpenFile(logFilePath+name+".log", os.O_APPEND, 0666)
	}

	//log.LUTC|log.Lmicroseconds ,世界统一时间|毫秒时间
	logger[name] = log.New(logFile[name], "", log.Ldate|log.Ltime)
}

//OpenLogWithCodeAndMsgN 开启并写入部分日志
func OpenLogWithCodeAndMsgN(name string, code int, msg string) {
	OpenLogN(name)
	if code == LvPanic {
		PanicN(name, msg)
	} else if code == LvERROR {
		ErrorN(name, msg)
	} else if code == LvWARNING {
		WarningN(name, msg)
	} else if code == LvINFO {
		InfoN(name, msg)
	} else {
		VerboseN(name, msg)
	}
}

//文件是否存在
func isFileExist(file string) bool {
	if _, err := os.Stat(file); os.IsExist(err) {
		return true
	}
	return false
}
