package glog

import (
	"fmt"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"time"
)

var isdebug = false
var defaultDeep = 2
var deep = 2

//初始化
func InitWithConfig(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration, debug bool) {

	baseLogFile := path.Join(logPath, logFileName)
	if logFileName == "" {
		baseLogFile = baseLogFile + "/"
	}
	debugWriter, err := rotatelogs.New(
		baseLogFile+"%Y%m%d%H%M_debug"+".log",
		rotatelogs.WithLinkName(baseLogFile+"now_debug.log"), //生成软连接,指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),                        //文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime),            //日志切割时间间隔
	)
	writer, err := rotatelogs.New(
		baseLogFile+"%Y%m%d%H%M"+".log",
		rotatelogs.WithLinkName(baseLogFile+"now.log"), //生成软连接,指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),                  //文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime),      //日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("Init log with config error.%+v", errors.WithStack(err))
	}
	//此处设置的Formatter只对文件输出生效
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: debugWriter,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &KtFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	isdebug = debug
	//此处设置的Formatter对非文件输出生效
	logrus.SetFormatter(&KtFormatter{
		ForceColors:      isdebug, //设置为true以在输出颜色之前绕过检查TTY。
		DisableColors:    false,   //禁用颜色
		DisableTimestamp: false,   //禁用时间戳
		FullTimestamp:    true,    //在连接TTY时启用记录完整时间戳，而不是仅记录自开始执行以来经过的时间。
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableSorting:   false,
	})
	level := logrus.InfoLevel
	if isdebug {
		level = logrus.DebugLevel
	}
	logrus.SetLevel(level)
	logrus.SetOutput(os.Stdout)
	//设置文件输出属性
	logrus.AddHook(lfHook)
}

//Deepin 设置临时深度
func Deepin(d int){
	if deep == defaultDeep {
		deep = d
	}
}

//获取行号
func GetLine(skip int) (ret string) {
	_, file, line, ok := runtime.Caller(skip)
	if ok {
		//temp := strings.Split(file, "/src/")
		//if len(temp) > 1 {
		//	file = temp[1]
		//}
		ret = fmt.Sprintf("%s:%d", file, line)
	}
	return
}

//格式化msg, 将其中的结构体转换成字符串,原理:利用switch type将地址转换类型为实际的值
func format(args ...interface{})interface{} {
	for k,v := range args {
		switch value := v.(type) {
		case string:
		case error:
		case int:
		default:
			args[k] = fmt.Sprintf("%+v",value)
		}
	}
	deep = defaultDeep
	return args
}

func Debug(args ...interface{}) {
	logrus.WithField("line", GetLine(deep)).Debug(format(args ...))
	deep = defaultDeep
}

func Info(args ...interface{}) {
	if isdebug {
		logrus.WithField("line", GetLine(deep)).Info(format(args ...))
	} else {
		logrus.Info(format(args ...))
	}
	deep = defaultDeep
}

func Warning(args ...interface{}) {
	if isdebug {
		logrus.WithField("line", GetLine(deep)).Warning(format(args ...))
	} else {
		logrus.Warning(format(args ...))
	}
	deep = defaultDeep
}

func Panic(args ...interface{}) {
	if isdebug {
		logrus.WithField("line", GetLine(deep)).Panic(format(args ...))
	} else {
		logrus.Panic(format(args ...))
	}
	deep = defaultDeep
}

func Error(err error, args ...interface{}) {
	if err != nil {
		logrus.WithError(err).WithField("line", GetLine(deep)).Error(format(args ...))
	} else {
		logrus.WithField("line", GetLine(deep)).Error(format(args ...))
	}
	deep = defaultDeep
}

func Fatal(args ...interface{}) {
	if isdebug {
		logrus.WithField("line", GetLine(deep)).Fatal(format(args ...))
	} else {
		logrus.Fatal(format(args ...))
	}
	deep = defaultDeep
}
