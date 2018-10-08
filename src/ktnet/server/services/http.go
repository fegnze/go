package services

import (
	"ktnet/ktcore/kthttp"
	"ktnet/server/protocols"
	"net/http"
)

//InitHTTP 启动http服务,监听addr
func InitHTTP(addr string) {
	sv := kthttp.RegistServer(&http.Server{Addr: addr})
	sv.RegistRout("/test", &protocols.ProtoTest{}, "get")
	sv.StartService()
}
