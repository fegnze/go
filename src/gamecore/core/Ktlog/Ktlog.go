package Ktlog

import (
	"fmt"
	"gamecore/core/conf"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const (
	lvERROR = iota
	lvDEBUG
	lvINFO
	lvVERBOSE
)

const (
	prefixERROR   = "[ERROR#!!]"
	prefixDEBUG   = "[#DEBUG#]"
	prefixINFO    = "[INFO]"
	prefixVERBOSE = "[VERBOSE]"
)

var logLevel = conf.LogLevel
var logFilePath = conf.LogFilePath
var logFile *os.File
var logger *log.Logger

//Panic 程序无法处理的错误,终端执行
func Panic(formatStr string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	callpos := " \tcallpos=>(" + file + ":" + strconv.Itoa(line) + ")"

	logger.SetPrefix(prefixERROR)
	if len(v) > 0 {
		if strings.Index(formatStr, "%") > -1 {
			str := fmt.Sprintf(formatStr, v...) + callpos
			logger.Panicln(str)
		} else {
			str := fmt.Sprint(v...)
			str = formatStr + "," + str + callpos
			logger.Panicln(str)
		}

	} else {
		logger.Panicln(formatStr + callpos)
	}
}

//Error 捕获到的逻辑异常,不中断运行
func Error(formatStr string, v ...interface{}) {
	if logLevel < lvERROR {
		return
	}

	_, file, line, _ := runtime.Caller(1)
	callpos := " \tcallpos=>(" + file + ":" + strconv.Itoa(line) + ")"

	logger.SetPrefix(prefixERROR)
	if len(v) > 0 {
		if strings.Index(formatStr, "%") > -1 {
			str := fmt.Sprintf(formatStr, v...) + callpos
			logger.Println(str)
		} else {
			str := fmt.Sprint(v...)
			str = formatStr + "," + str + callpos
			logger.Println(str)
		}

	} else {
		logger.Println(formatStr + callpos)
	}
}

//Debug 调试模式,输出格式包含文件名和行号
func Debug(formatStr string, v ...interface{}) {
	if logLevel < lvDEBUG {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	callpos := " \tcallpos=>(" + file + ":" + strconv.Itoa(line) + ")"

	logger.SetPrefix(prefixDEBUG)
	if len(v) > 0 {
		if strings.Index(formatStr, "%") > -1 {
			str := fmt.Sprintf(formatStr, v...) + callpos
			logger.Println(str)
		} else {
			str := fmt.Sprint(v...)
			str = formatStr + "," + str + callpos
			logger.Println(str)
		}

	} else {
		logger.Println(formatStr + callpos)
	}
}

//Info 基本的需要记录的日志
func Info(formatStr string, v ...interface{}) {
	if logLevel < lvINFO {
		return
	}
	logger.SetPrefix(prefixINFO)
	if len(v) > 0 {
		if strings.Index(formatStr, "%") > -1 {
			str := fmt.Sprintf(formatStr, v...)
			logger.Println(str)
		} else {
			str := fmt.Sprint(v...)
			str = formatStr + "," + str
			logger.Println(str)
		}

	} else {
		logger.Println(formatStr)
	}
}

//Verbose 额外的不比体现在日志流程的日志
func Verbose(formatStr string, v ...interface{}) {
	if logLevel < lvVERBOSE {
		return
	}
	logger.SetPrefix(prefixVERBOSE)
	if len(v) > 0 {
		if strings.Index(formatStr, "%") > -1 {
			str := fmt.Sprintf(formatStr, v...)
			logger.Println(str)
		} else {
			str := fmt.Sprint(v...)
			str = formatStr + "," + str
			logger.Println(str)
		}

	} else {
		logger.Println(formatStr)
	}
}

//CloseLog 关闭日志的文件记录
func CloseLog() {
	if logFile != nil {
		fmt.Println("close log file ... ")
		logFile.Close()
	}
}

//OpenLog 打开日志的文件记录
func OpenLog() {
	if logger != nil {
		Error("logger is already exist ...")
	}
	var err error
	if _, err1 := os.Stat(logFilePath); os.IsNotExist(err1) {
		logFile, err = os.Create(logFilePath)
		if err != nil {
			log.Fatalf("%v\n,open log file err ... %s", err, logFilePath)
		}
	} else {
		logFile, err = os.OpenFile(logFilePath, os.O_APPEND, 0666)
	}

	//log.LUTC|log.Lmicroseconds ,世界统一时间|毫秒时间
	logger = log.New(logFile, "", log.Ldate|log.Ltime)
}

//文件是否存在
func isFileExist(file string) bool {
	if _, err := os.Stat(file); os.IsExist(err) {
		return true
	}
	return false
}
