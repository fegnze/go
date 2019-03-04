package main

import (
	"KiteNet/cmd/web/config"
	"KiteNet/cmd/web/db"
	"KiteNet/cmd/web/global"
	"KiteNet/cmd/web/handlers"
	. "KiteNet/httpkt"
	"KiteNet/log"
	"KiteNet/tcp/client"
	"KiteNet/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	dir, err := filepath.Abs("./")
	fmt.Println(dir)
	if err != nil {
		panic(err)
	}
	defaultLogFile, err := os.OpenFile(dir+"/output", os.O_WRONLY, 0666)
	log.SetOutput(defaultLogFile)
	defer func() {
		e := defaultLogFile.Close()
		log.Println(e)
	}()

	app := &cli.App{
		Name:        "KiteNet Http service",
		Usage:       "usage:",
		UsageText:   "operation the httpkt server",
		ArgsUsage:   "",
		Version:     "0.0.1",
		Description: "a web app for new game",
		//子命令eg:git version,git add
		Commands: []*cli.Command{
			{
				Name:     "start",
				Usage:    "启动服务",
				Category: "cmd",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "conf",
						Aliases: []string{"c"},
						Usage:   "配置文档的完整路径",
						//Value:   "C:/Users/Administrator/go/src/KiteNet/cmd/web/config/configs.json",
					},
					&cli.BoolFlag{
						Name:    "debug",
						Aliases: []string{"d"},
						EnvVars: []string{"ISDEBUG"},
						Usage:   "是否是debug模式",
					},
				},
				Action: func(ctx *cli.Context) error { //与app的action互斥
					file := ctx.String("conf")
					log.Println("配置文件:", file)
					//读取配置
					conf := &config.Configs{}
					global.Conf = conf
					buf, err := ioutil.ReadFile(file)
					utils.CheckErr(err)
					err = json.Unmarshal(utils.JsonNormalize(buf), conf)
					utils.CheckErr(err)
					log.Println("读取配置:", conf)

					//初始化日志系统
					glog.InitWithConfig(conf.Log.Path,
						conf.Log.FilePrefixName,
						time.Hour*conf.Log.FileDuration,
						time.Hour*conf.Log.SplitDuration,
						ctx.Bool("debug"))
					e := defaultLogFile.Close()
					utils.CheckErr(e)

					//写入全局环境变量
					isdebug := ctx.String("debug")
					os.Setenv("SERVER_NAME", conf.Web.Name)
					os.Setenv("DEBUG_MODLE", isdebug)
					os.Setenv("SERVER_HOST", conf.Web.Ip+":"+conf.Web.Port)
					os.Setenv("WORLD_IP", conf.World.Ip)
					os.Setenv("WORLD_PORT", conf.World.Port)

					//初始化DB服务
					glog.Info("初始化db")
					db.Init(&db.Config{
						User:     conf.Web.DB.User,
						Password: conf.Web.DB.Password,
						Addr:     conf.Web.DB.Host,
						DBName:   conf.Web.DB.DBName,
					})
					defer db.Close()

					//初始化world服务器socket
					glog.Info("连接world")
					global.WorldSS = &client.Client{}
					defer global.WorldSS.Close()
					global.WorldSS.Connect(fmt.Sprintf("%s:%s", conf.World.Ip, conf.World.Port))

					//初始化HTTP服务
					glog.Info("初始化HTTP服务")
					app := &App{
						Name:           conf.Web.Name,
						Addr:           conf.Web.Ip + ":" + conf.Web.Port,
						Handler:        http.NewServeMux(),
						ReadTimeout:    10 * time.Second,
						WriteTimeout:   10 * time.Second,
						MaxHeaderBytes: 1 << 20,

						FileServePath:   conf.Web.FS.LocalPath,
						FileServePrefix: conf.Web.FS.Prefix,
					}

					//注册路由表,并开始监听http服务
					return app.Start(handlers.RoutMap)
				},
			},
			{
				Name:     "stop",
				Usage:    "停止服务",
				Category: "cmd",
				//可以定义单独的before,after
				Action: func(ctx *cli.Context) error {
					glog.Info("stop server,version:", ctx.String("version"))
					return nil
				},
			},
		},
		EnableShellCompletion: false,
		Authors: []*cli.Author{
			{
				Email: "fegnze@outlook.com",
				Name:  "fegnze",
			},
		},
		Copyright: "(c) 1999 Serious Enterprise",
	}

	//指令排序
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	//解析命令
	if err := app.Run(os.Args); err != nil {
		log.Println("cli启动失败")
	}
}
