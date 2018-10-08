package main

import (
	"ktnet/ktcore"
	"ktnet/ktcore/ktlog"
	"ktnet/server/services"
	"time"
)

func release() {
	ktcore.Release()
	time.Sleep(10 * time.Second)
}

func main() {
	defer release()

	ktcore.Init("configs.json")

	ktlog.Info("Hello GO")
	ktlog.Verbose("This is a verbose", "...")
	ktlog.Warning("This is a debug %d,%s", 1, "***")
	ktlog.Error("This is a Error log", 1, 2, 3)
	//ktlog.Panic("无法解决的异常,阻断程序运行")

	ktlog.OpenLogN("login")
	ktlog.InfoN("login", "这是另一个logger")

	services.Init()
	ktlog.Info("-----------")

}
