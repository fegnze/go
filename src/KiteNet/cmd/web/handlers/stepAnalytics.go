package handlers

import (
	"KiteNet/httpkt"
	"KiteNet/log"
	"net/http"
)

//StepAnalyticsReqProto 客户端请求数据,本地类型
type StepAnalyticsReqProto struct {
	Region     string `json:"region"`
	Openid     string `json:"openid"`
	UID        string `json:"uid"`
	Channel    string `json:"channel"`
	Step       string `json:"step"`
	UniqueCode string `json:"uniqueCode"`
	DeviceType string `json:"deviceType"`
}

//StepAnalyticsResProto 返回数据
type StepAnalyticsResProto struct{}

//StepAnalyticsHandler 处理函数
func StepAnalyticsHandler(w http.ResponseWriter, r *http.Request) {
	glog.Debug("VersionHandle HttpExec ...")
	reqProto := &StepAnalyticsReqProto{}
	httpkt.ParseReqProto(r, reqProto)
	glog.Debug(reqProto)
	//TODO -执行服务逻辑

	httpkt.ReturnNull(w)
}
