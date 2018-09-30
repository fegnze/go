package main

import (
	"gamecore/core/Ktlog"
	"gamecore/core/conf"
	"gamecore/server/services"
)

func main() {
	//Ktlog.OpenLog()
	// Ktlog.Info("This is a info ... ")
	// Ktlog.Verbose("This is a verbose ", "log")
	// Ktlog.Debug("This is a debug %d,%s", 1, "***")
	// //Ktlog.Error("This is a Error log", 1, 2, 3)
	conf.Parse("conf/configs.json")
	Ktlog.OpenLog()

	test := "Hello GO"
	Ktlog.Info(test)

	services.StartHTTPServer(":3000")
	Ktlog.Info("-----------")
}
