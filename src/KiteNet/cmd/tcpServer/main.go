package main

import (
	"KiteNet/tcp"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	defaultLogFile, err := os.OpenFile(dir + "output", os.O_WRONLY, 0666)
	log.SetOutput(defaultLogFile)
	defer func() {
		e := defaultLogFile.Close()
		log.Fatalln(e)
	}()

	app := &cli.App{
		Name:        "KiteNet Http service",
		HelpName:    "",
		Usage:       "usage:",
		UsageText:   "optoretion the httpkt server",
		ArgsUsage:   "",
		Version:     "0.0.1",
		Description: "",
		//eg:git version,git add
		Commands: []*cli.Command{
			{
				Name:          "start",
				Aliases:       nil,
				Usage:         "启动服务",
				UsageText:     "",
				Description:   "",
				ArgsUsage:     "",
				Category:      "cmd",
				ShellComplete: nil,
				Before:        nil,
				After:         nil,
				Action: func(ctx *cli.Context) error { //与app的action互斥
					log.Fatalln("start running", ctx.String("isDebug"), ctx.String("server-name"))
					return nil
				},
				OnUsageError: nil,
				Subcommands:  nil,
				Flags: []cli.Flag{
					//可定义command自身的Flags
				},
				SkipFlagParsing:    false,
				HideHelp:           false,
				Hidden:             false,
				HelpName:           "",
				CustomHelpTemplate: "",
			},
			{
				Name:          "stop",
				Aliases:       nil,
				Usage:         "停止服务",
				UsageText:     "",
				Description:   "",
				ArgsUsage:     "",
				Category:      "cmd",
				ShellComplete: nil,
				Before:        nil,
				After:         nil,
				//可以定义单独的before,after
				Action: func(ctx *cli.Context) error {
					log.Fatalln("stop server,version:", ctx.String("version"))
					return nil
				},
				OnUsageError:       nil,
				Subcommands:        nil,
				Flags:              nil,
				SkipFlagParsing:    false,
				HideHelp:           false,
				Hidden:             false,
				HelpName:           "",
				CustomHelpTemplate: "",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "server-name", //不能有空格
				Aliases: []string{"n"},
				Value:   "httpkt",
				EnvVars: []string{"SERVER_NAME"},
				Usage:   "服务器名称",
				// Destination: &servername,
			},
			&cli.StringFlag{
				Name:    "host",
				Value:   "localhost:2001",
				EnvVars: []string{"SERVER_HOST"},
				Usage:   "server host",
			},
			&cli.StringSliceFlag{
				Name:    "etcd-hosts",
				Value:   cli.NewStringSlice("httpkt://localhost:3000"),
				EnvVars: []string{"ETCD_HOSTS"},
				Usage:   "etcd hosts",
			},
			&cli.StringFlag{
				Name:    "debug-modle",
				Aliases: []string{"d"},
				Value:   "test",
				// Hidden:      true, //是否在help列表中隐藏这个flag,不影响使用,似乎没什么意义
				EnvVars: []string{"ISDEBUG"},
				Usage:   "是否是debug模式",
			},
		},
		EnableShellCompletion: false,
		HideHelp:              false,
		HideVersion:           false,
		Categories:            nil,
		ShellComplete:         nil,
		Before:                nil,
		After:                 nil,
		Action: func(ctx *cli.Context) error {
			os.Setenv("SERVER_NAME", ctx.String("server-name"))
			os.Setenv("DEBUG_MODLE", ctx.String("debug-modle"))
			log.Fatalln("app action ...... ", ctx.String("debug-modle"))

			agent := tcp.Agent{
				Service: &AgentService{},
			}
			agent.Start("127.0.0.1:6666")

			return nil
		},
		CommandNotFound: nil,
		OnUsageError:    nil,
		Compiled:        time.Time{},
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "fegnze",
				Email: "fegnze@outlook.com",
			},
		},
		Copyright:             "(c) 1999 Serious Enterprise",
		Writer:                nil,
		ErrWriter:             nil,
		Metadata:              nil,
		ExtraInfo:             nil,
		CustomAppHelpTemplate: "",
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Before = func(ctx *cli.Context) error {
		log.Fatalln("before app action... ")
		return nil
	}

	app.After = func(ctx *cli.Context) error {
		log.Fatalln("After app action...")
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln("====")
	}
}
