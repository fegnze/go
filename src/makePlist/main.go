package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main(){
	for i,entry := range os.Args {
		if i == 0 {
			continue
		}
		entry = strings.Replace(entry,"\\","/",-1)
		fmt.Println("获取到源路径:",entry)
		parts := strings.Split(entry,"/.")
		if len(parts) < 2 {
			fmt.Println("无效的源文件路径")
			continue
		}
		out := parts[0]

		name := strings.Split(parts[1],"_PList")[0]
		if len(name) < 2 {
			fmt.Println("无效的源文件路径")
			continue
		}
		fmt.Println("文件名:",name)
		fmt.Println(entry)
		cmd := fmt.Sprintf(
			"--data %s/%s.plist --sheet %s/%s.png --format cocos2d --trim-mode None %s/",
			out,name,
			out,name,
			entry)

		fmt.Println("执行命令:",cmd)

		args := strings.Split(cmd," ")
		exe := exec.Command("C:/Program Files (x86)/CodeAndWeb/TexturePacker/bin/TexturePacker.exe",args[0:]...)
		bytes,err := exe.Output()

		fmt.Println("=====",string(bytes))
		if err != nil {
			fmt.Println(err)
		}
	}
	time.Sleep(time.Second * 5)
}
