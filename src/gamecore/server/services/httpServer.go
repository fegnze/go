package services

import (
	"gamecore/core/ktNet/KtHttp"
	"gamecore/server/protocols"
	"net/http"
)

//StartHTTPServer 启动http服务,监听addr
func StartHTTPServer(addr string) {
	sv := KtHttp.RegistServer(&http.Server{Addr: addr})
	sv.RegistRout("/test", &protocols.ProtoTest{}, "get")
	sv.StartService()
}
