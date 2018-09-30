package main

import (
	"ktnet/core/ktlog"
	"ktnet/server/services"
)

func main() {
	//Ktlog.OpenLog()
	// Ktlog.Info("This is a info ... ")
	// Ktlog.Verbose("This is a verbose ", "log")
	// Ktlog.Debug("This is a debug %d,%s", 1, "***")
	// //Ktlog.Error("This is a Error log", 1, 2, 3)
	conf.Parse("conf/configs.json")
	ktlog.OpenLog()

	test := "Hello GO"
	ktlog.Info(test)

	services.StartHTTPServer(":3000")
	ktlog.Info("-----------")
}
