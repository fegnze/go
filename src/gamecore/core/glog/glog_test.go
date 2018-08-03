package glog

import (
	"testing"
)

func TestDebug(t *testing.T) {
	Debug()
}

func BenchmarkDebug(t *testing.B) {
	Debug()
}
