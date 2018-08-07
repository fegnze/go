package glog

import (
	"testing"
)

func TestDebug(t *testing.T) {
	OpenLog()
	Verbose("第%d条", 1, 2)
	Verbose("第2条")
	Info("第3条")
	Debug("第4条")
	Verbose("第5条")
	Debug("第6条")
	Error("===========")
	Debug("第7条", 111, 222)

	CloseLog()
	Debug("第8条")
}

func BenchmarkDebug(t *testing.B) {
	OpenLog()
	Debug("第%d条", 1)
	Debug("第2条")
	Debug("第3条")
	Debug("第4条")
	Debug("第5条")
	Debug("第6条")
	Debug("第7条", 111, 222)
	CloseLog()
	Debug("第8条")
}
