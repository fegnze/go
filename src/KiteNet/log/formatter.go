package glog

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"runtime"
	"sort"
	"strings"
	"time"
)

const (
	nocolor = 0
	red     = 31
	green   = 32
	yellow  = 33
	blue    = 34
	gray    = 37
)

var (
	baseTimestamp time.Time
	isTerminal    bool
)

func init() {
	baseTimestamp = time.Now()
	isTerminal = logrus.IsTerminal()
}

func miniTS() int {
	return int(time.Since(baseTimestamp) / time.Second)
}

type KtFormatter struct {
	// Set to true to bypass checking for a TTY before outputting colors.
	ForceColors bool

	// Force disabling colors.
	DisableColors bool

	// Disable timestamp logging. useful when output is redirected to logging
	// system that already adds timestamps.
	DisableTimestamp bool

	// Enable logging the full timestamp when a TTY is attached instead of just
	// the time passed since beginning of execution.
	FullTimestamp bool

	// TimestampFormat to use for display when a full timestamp is printed
	TimestampFormat string

	// The fields are sorted by default for a consistent output. For applications
	// that log extremely frequently and don't use the JSON formatter this may not
	// be desired.
	DisableSorting bool
}

func prefixFieldClashes(data logrus.Fields) {
	_, ok := data["time"]
	if ok {
		data["fields.time"] = data["time"]
	}

	_, ok = data["msg"]
	if ok {
		data["fields.msg"] = data["msg"]
	}

	_, ok = data["level"]
	if ok {
		data["fields.level"] = data["level"]
	}
}

func (f *KtFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var keys []string = make([]string, 0, len(entry.Data))
	for k := range entry.Data {
		keys = append(keys, k)
	}

	if !f.DisableSorting {
		sort.Strings(keys)
	}

	b := &bytes.Buffer{}

	prefixFieldClashes(entry.Data)

	isColorTerminal := isTerminal && (runtime.GOOS != "windows")
	isColored := (f.ForceColors || isColorTerminal) && !f.DisableColors
	if isTerminal && (runtime.GOOS == "windows") {
		isColored = false
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = logrus.DefaultTimestampFormat
	}
	if isColored {
		f.printColored(b, entry, keys, timestampFormat)
	} else {
		if !f.DisableTimestamp {
			f.appendKeyValue(b, "time", entry.Time.Format(timestampFormat))
		}
		f.appendKeyValue(b, "level", "["+strings.ToUpper(entry.Level.String())+"]")
		f.appendKeyValue(b, "msg", entry.Message[1:len(entry.Message)-1])
		for _, key := range keys {
			f.appendKeyValue(b, key, entry.Data[key])
		}
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *KtFormatter) printColored(b *bytes.Buffer, entry *logrus.Entry, keys []string, timestampFormat string) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel:
		levelColor = green
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = nocolor
	}

	levelText := strings.ToUpper(entry.Level.String())[0:4]

	////s := "          " + "          " + "          " +
	//	"          " + "          " + "          " +
	//	"          " + "          " + "          "
	//颜色语法:   \x1b[%d;%dm要显示颜色的内容\x1b[0m
	//0即是无色,第一个%d是底色(省略默认为0色),第二个%d是字色,结束重新赋值成0色
	if !f.FullTimestamp {
		_, err := fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%04d] \x1b[%dm	%s\x1b[0m", levelColor, levelText, miniTS(), levelColor, entry.Message[1:len(entry.Message)-1])
		if err != nil {
			log.Panicln(err)
		}
	} else {
		_, err := fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%s] \x1b[%dm	%s\x1b[0m", levelColor, levelText, entry.Time.Format(timestampFormat),
			levelColor, entry.Message[1:len(entry.Message)-1])
		if err != nil {
			log.Panicln(err)
		}

	}

	lineCount := 1
	for _, k := range keys {
		v := entry.Data[k]
		c := levelColor
		if k == "line" {
			c = yellow
		}
		str := ""
		if lineCount >= 1 && len(keys) > 1 {
			str = "\n							" + str
		}else {
			str = "	"
		}
		ss := fmt.Sprintf("%s=[ %+v ]", k, v)
		str = str + fmt.Sprintf("%50s",ss)
		_, err := fmt.Fprintf(b, "\x1b[%dm  %s", c, str)
		if err != nil {
			log.Panicln(err)
		}
		lineCount ++
	}
}

func needsQuoting(text string) bool {
	for _, ch := range text {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '.') {
			return false
		}
	}
	return true
}

func (f *KtFormatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {

	//b.WriteString(key)
	//b.WriteByte('=')

	switch value := value.(type) {
	case string:
		if key == "line" {
			value = "(" + value + ")"
		}
		if needsQuoting(value) {
			b.WriteString(value)
		} else {
			_, err := fmt.Fprintf(b, "%s", value)
			log.Println(err)
		}
	case error:
		errmsg := value.Error()
		if needsQuoting(errmsg) {
			b.WriteString("[")
			b.WriteString(errmsg)
			b.WriteString("]")
		} else {
			_, err := fmt.Fprintf(b, "[%s]", value)
			log.Println(err)
		}
	default:
		_, err := fmt.Fprint(b, value)
		log.Println(err)
	}

	b.WriteByte(' ')
}
