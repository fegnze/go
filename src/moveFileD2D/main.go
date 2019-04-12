package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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

func copyFile(orign string,dest string){
	content,err:= ioutil.ReadFile(orign)
	checkErr(err)
	if err != nil{
		return
	}

	err = ioutil.WriteFile(dest,content,0644)
	checkErr(err)
}

func main() {
	a := os.Args[0]
	a = strings.Replace(a, "\\\\", "/", -1)
	a = strings.Replace(a, "\\", "/", -1)
	index := strings.LastIndex(a, "/")
	path = a[:index]
	logFile, _ := os.OpenFile(path+"/moveFileD2D.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer func() {
		err := logFile.Close()
		if err != nil {
			panic(err)
		}
	}()
	loger = log.New(logFile, "", log.LstdFlags)

	b := os.Args[1]
	if len(os.Args) <= 1 {
		loger.Panicln("未指定工作文件!")
		return
	}
	if _,err:=os.Stat(b);os.IsExist(err){
		loger.Panicln("指定文件不存在!")
		return
	}
	confs,err:=os.Open(b)
	if err != nil{
		loger.Panicln("指定文件无法打开!",b)
		return
	}
	defer func() {
		err := confs.Close()
		checkErr(err)
	}()

	orignpath := ""
	destPath := ""

	reader := bufio.NewReader(confs)
	for{
		if line,prefix,err := reader.ReadLine();err==nil{
			checkErr(err)
			if prefix {
				loger.Panicln("这行什么这么长?")
				continue
			}else{
				if destPath == "" {
					destPath = string(line)
					continue
				}
				if orignpath == "" {
					orignpath = string(line)
					continue
				}

				str := string(line)
				parts := strings.Split(str,"\t")
				if len(parts) == 1 {
					orign := orignpath+"/" + parts[0]
					namestr:=strings.Replace(parts[0],"/","\\",-1)
					ns := strings.Split(namestr,"\\")
					name:=ns[len(ns)-1]
					dest := destPath+"/"+name
					loger.Println("该条目未配置成对值,作为单文件拷贝到输出目录",orign)
					fmt.Println("该条目未配置成对值,作为单文件拷贝到输出目录",orign)

					if _,err:=os.Stat(orign);os.IsNotExist(err){
						loger.Println("操作文件不存在")
						fmt.Println("操作文件不存在")
						continue
					}

					copyFile(orign,dest)
					continue
				}

				orign := orignpath+"/" + parts[1]
				tmp:= parts[0]
				names:=strings.Split(tmp,".")
				name := names[0]
				dest := destPath+"/."+name+"_PList.Dir/"+ parts[0]
				destBk := destPath+"/."+name+"_PList.Dir/backup_"+ parts[0]

				orign = strings.Replace(orign,"/","\\",-1)
				dest = strings.Replace(dest,"/","\\",-1)
				destBk = strings.Replace(destBk,"/","\\",-1)

				loger.Println("执行:",fmt.Sprintf("%s",orign),dest)
				//fmt.Println("执行:",orign,dest)

				_,e1:= os.Stat(orign)
				_,e2:= os.Stat(dest)
				if os.IsNotExist(e1) {
					loger.Println("文件不存在:",orign)
					continue
				}
				if os.IsNotExist(e2){
					loger.Println("文件不存在:",dest)
				}

				fmt.Println("move:",orign,dest)
				copyFile(fmt.Sprintf("%s",dest),destBk)

				copyFile(orign,dest)
			}
		} else {
			break
		}
	}
	loger.Println("执行完毕.")
	fmt.Println("执行完毕.")
}
