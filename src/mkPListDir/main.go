package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var loger *log.Logger
var path string

func checkErr(err error) {
	if err != nil {
		loger.Println(err)
		panic(err)
	}
}

func main() {
	a := os.Args[0]
	a = strings.Replace(a, "\\\\", "/", -1)
	a = strings.Replace(a, "\\", "/", -1)
	index := strings.LastIndex(a, "/")
	path = a[:index]
	logFile, _ := os.OpenFile(path+"/mkPlistDir.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer func() {
		err := logFile.Close()
		if err != nil {
			panic(err)
		}
	}()
	loger = log.New(logFile, "", log.LstdFlags)

	if len(os.Args) <= 1 {
		loger.Panicln("未指定工作文件!")
		return
	}

	for i := 1; i < len(os.Args); i++ {
		b := os.Args[i]
		if _, err := os.Stat(b); os.IsExist(err) {
			loger.Panicln("指定文件不存在!")
			return
		}

		b = strings.Replace(b, "/", "\\", -1)
		parts := strings.Split(b, "\\")
		filename := parts[len(parts)-1]
		names := strings.Split(filename, ".")
		i := len(names) - 1
		endfix := names[i]
		name := strings.Replace(filename, "."+endfix, "", -1)
		name = "." + name + "_PList.Dir"

		path := strings.Replace(b, filename, "", -1)
		loger.Println("创建目录:", path+name)
		fmt.Println("创建目录:", path+name)
		err := os.Mkdir(path+name, os.ModePerm)
		checkErr(err)

		fmt.Println("移动文件:", b, path+name+"\\"+filename)
		loger.Println("移动文件:", b, path+name+"\\"+filename)
		err = os.Rename(b, path+name+"\\"+filename)
		checkErr(err)
	}

	loger.Println("执行完毕.")
	fmt.Println("执行完毕.")
}
