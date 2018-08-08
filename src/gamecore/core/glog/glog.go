package glog

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

//Error for log
func Error(formatStr string, v ...interface{}) {
	if logLevel < lvERROR {
		return
	}

	_, file, line, _ := runtime.Caller(1)
	callpos := " \tcallpos=>(" + file + strconv.Itoa(line) + ")"

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

//Debug for log
func Debug(formatStr string, v ...interface{}) {
	if logLevel < lvDEBUG {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	callpos := " \tcallpos=>(" + file + strconv.Itoa(line) + ")"

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

//Info for log
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

//Verbose for log
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

//CloseLog close logfile
func CloseLog() {
	if logFile != nil {
		fmt.Println("close log file ... ")
		logFile.Close()
	}
}

//OpenLog 1
func OpenLog() {
	if logger != nil {
		Error("logger is already exist ...")
	}
	var err error
	if _, err1 := os.Stat(logFilePath); os.IsNotExist(err1) {
		//fmt.Println("文件不存在，创建文件")
		logFile, err = os.Create(logFilePath)
		if err != nil {
			// fmt.Println(logFilePath)
			log.Fatalf("%v\n,open log file err ... %s", err, logFilePath)
		}
	} else {
		//fmt.Println("文件存在")
		logFile, err = os.OpenFile(logFilePath, os.O_APPEND, 0666)
	}

	//log.LUTC|log.Lmicroseconds ,世界统一时间|毫秒时间
	logger = log.New(logFile, "", log.Ldate|log.Ltime)
	//fmt.Println("创建logger。")
}

func isFileExist(file string) bool {
	if _, err := os.Stat(file); os.IsExist(err) {
		return true
	}
	return false
}

func init() {
	OpenLog()
}
