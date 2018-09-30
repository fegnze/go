package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type conf struct {
	ResDir string
}

func main() {
	var loger *log.Logger
	if _, e := os.Stat("./log"); e != nil {
		logFile, _ := os.Create("./log")
		defer logFile.Close()
		loger = log.New(logFile, "", 0)
	} else {
		logFile, _ := os.OpenFile("./log", os.O_APPEND, 0666)
		defer logFile.Close()
		loger = log.New(logFile, "", 0)
	}

	resDir := "."
	//1.从配置文件读取png目录和输出目录
	conFile := "./ini.json"
	file, err := ioutil.ReadFile(conFile)
	if err != nil {
		loger.Panicln("[ERROR:]读取配置文件失败...", err.Error())
	} else {
		var cfst conf
		if err = json.Unmarshal(file, &cfst); err != nil {
			loger.Panicln("[ERROR:]解析配置文件失败...", err.Error())
		} else {
			resDir = cfst.ResDir
		}
	}
	loger.Println("resDir:" + resDir)

	//2.读取ccb
	args := []string{}
	if we := filepath.Walk("./ccb", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			loger.Panicln("[ERROR]遍历ccb文件" + path + "有误," + err.Error())
		}
		if info.IsDir() {
			loger.Println("[WARNING]遍历ccb文件夹," + path + "忽略文件夹,")
			return nil
		}
		if strings.Contains(info.Name(), ".ccb") && !strings.Contains(info.Name(), ".ccbi") {
			args = append(args, path)
		}
		return nil
	}); we != nil {
		loger.Panicln("[ERROR]遍历ccb文件夹失败" + we.Error())
	}
	// args := []string{"", `E:\WorkSpace\Kapai\Games\sks-heti\ui_ccb\ui\Resources\layout\interface\accouterexchange01.ccb`, `E:\WorkSpace\Kapai\Games\sks-heti\ui_ccb\ui\Resources\layout\interface\accouterinfo01.ccb`}
	// for _, f := range args[1:] {
	for _, f := range args[1:] {
		loger.Println("ccbfile:" + f)
		f = strings.Replace(f, "\\", "/", -1)
		ccbfile, err := ioutil.ReadFile(f)
		if err != nil {
			loger.Panicln("读取ccb文件"+f+"失败", err.Error())
		}

		ccbstr := string(ccbfile)
		//3.筛选ccb中的png文件
		reg, _ := regexp.Compile(`[\w|\/]+.png`)
		pngs := reg.FindAllString(ccbstr, -1)
		tmps := strings.Split(f, "/")
		ccbn := tmps[len(tmps)-1]
		ccbnamestr := strings.Split(ccbn, ".")
		ccbname := ccbnamestr[0]

		//4.从png目录中拉取目标文件
		for _, png := range pngs {
			path := resDir + "/" + png
			pngfile, err := ioutil.ReadFile(path)
			if err != nil {
				loger.Println("[ERROR]读取png文件" + path + "失败")
			} else {
				fp := "./output/" + ccbname + "/" + png
				fdindex := strings.LastIndex(fp, "/")
				fd := fp[:fdindex]
				os.MkdirAll(fd, 0777)
				out := "./output/" + ccbname + "/" + png
				if err := ioutil.WriteFile(fp, pngfile, os.ModePerm); err != nil {
					loger.Panicln("写入文件png文件"+out+"失败", err.Error())
				} else {
					loger.Println("拷贝文件" + path + "到" + out)
				}
			}
		}
	}

}
