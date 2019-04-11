package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var loger *log.Logger
var path string

func init() {

}

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
	logFile, _ := os.OpenFile(path+"/mergefiles.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer func() {
		err := logFile.Close()
		if err != nil {
			panic(err)
		}
	}()
	loger = log.New(logFile, "", log.LstdFlags)

	output, err := os.OpenFile(path+"/TemplateLang.lua", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		loger.Println(err)
	}
	defer func() {
		err := output.Close()
		checkErr(err)
	}()
	for i, entry := range os.Args {
		if i == 0 {
			continue
		}
		entry = strings.Replace(entry, "\\", "/", -1)
		loger.Printf("读取文件:%s\n", entry)

		f, err := os.Open(entry)
		checkErr(err)
		data, err := ioutil.ReadAll(f)
		checkErr(err)

		n, err := output.Write(data)
		checkErr(err)
		loger.Printf("字节:%d\n", n)
		fmt.Println("字节:", n)

		h := []byte("\n\n")
		output.Write(h)
	}
	loger.Println("执行完毕.")
	fmt.Println("执行完毕.")
}
