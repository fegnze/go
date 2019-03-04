package httpkt

import (
	"KiteNet/log"
	"fmt"
	"net/http"
	"strings"
	"time"
)

//RoutMap 路由表类型定义
type RoutMap map[string]HandleFunc

//App 包装http server,每个App会创建自身的Mux,指定独立的路由表
type App struct {
	Name           string
	Addr           string
	Handler        *http.ServeMux
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int

	routs           RoutMap
	defaultRootFunc HandleFunc

	FileServePath string //文件托管服务的绑定路径
	FileServePrefix string //文件托管服务的访问前缀
}

//

/*Start 开启该Http服务器
*参数 rtm RoutMap 路由表
*路由表规则, 1,类型:httpkt.RoutMap (map[string]httpkt.IHandler)
*			2,可注册具体路径(第一个字符未"/"),eg:"/login/login"
*			3,若指定非路径,但request参数中含有cmd字段,则通过"/"分发cmd
*可通过SetRootDefaultFunc指定根路径解析,所有请求会先经过这个函数,后续考虑增加路径传递
 */
func (app *App) Start(rtm RoutMap) error {
	app.routs = rtm
	favicon := app.routs["/favicon.ico"]
	if favicon == nil {
		app.routs["/favicon.ico"] = FaviconHandler
	}

	for k,v := range app.routs {
		if !(strings.Index(k, "/") == 0) {
			continue
		}

		app.Handler.Handle(k,v)
	}

	if app.FileServePath != "" {
		glog.Info("注册文件服务 ...")
		if app.FileServePrefix == "" {
			app.FileServePrefix = "/download/"
		}
		//这种方式会创建一个filehandler,执行一个带重定向的serveFile,
		//比如说如果目录中有index.xml,/download/重定向到 /index.html
		app.Handler.Handle("/",http.StripPrefix(app.FileServePrefix, http.FileServer(http.Dir(app.FileServePath))))
	}

	server := &http.Server{
		Addr:           app.Addr,
		Handler:        app.Handler,
		ReadTimeout:    app.ReadTimeout,
		WriteTimeout:   app.WriteTimeout,
		MaxHeaderBytes: app.MaxHeaderBytes,
	}

	glog.Info(fmt.Sprintf("listening at \"%s\"",app.Addr))
	return server.ListenAndServe()
}

//SetRootDefaultFunc ,所有请求会先经过这个函数,后续考虑增加路径传递
func (app *App) SetRootDefaultFunc(f HandleFunc) {
	app.defaultRootFunc = f
}
