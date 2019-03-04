package account

import (
	"KiteNet/cmd/web/global"
	"KiteNet/cmd/web/global/code/ErrorCode"
	"KiteNet/httpkt"
	"KiteNet/log"
	"encoding/json"
	"fmt"
	"net/http"
)

//MixLoginReqProto 客户端请求数据,本地类型
type MixLoginReqProto struct {
	Openid     string `json:"openid"`
	Region     string `json:"region"`
	UniqueCode string `json:"uniqueCode"`
	DeviceType string `json:"deviceType"`
	Channel    string `json:"channel"`
	Origin     string `json:"origin"`
}

//MixLoginWordResProto world返回的数据
type MixLoginWorldResProto struct {
	Id     string `json:"id"`
	Secret string `json:"secret"`
	Server string `json:"server"`
	Port   string `json:"port"`
	Level  string `json:"level"`
	Name   string `json:"name"`
	New    string `json:"new"`
	Sid    string `json:"sid"`
}

//MixLoginResProto 返回数据
type MixLoginResProto struct {
	Server string `json:"server"` //服务器IP
	Port   string `json:"port"`
	Secret string `json:"secret"` //秘钥
	UID    string `json:"uid"`
	Region string `json:"region"`
	Openid string `json:"openid"`
	IsNew  string `json:"isnew"`
}

//MixLoginHandler login/mix处理函数
func MixLoginHandler(w http.ResponseWriter, r *http.Request) {
	glog.Debug("VersionHandle HttpExec ...")
	var reqProto = &MixLoginReqProto{}
	httpkt.ParseReqProto(r, reqProto)

	//发送请求到world服务器
	req := fmt.Sprintf("ul,win32_cn_999_%s_%s,%s,%d,%s", reqProto.Openid,reqProto.Region,reqProto.Channel, 0, reqProto.Region)

	session := global.WorldSS.NewSession()
	defer session.Close()
	response := session.Write([]byte(req))

	glog.Debug("world ul",string(response[0:]))

	//TODO -写入返回消息
	res := MixLoginWorldResProto{}
	err := json.Unmarshal(response[0:],&res)
	if err != nil {
		glog.Error(err,"MixLoginHandler 解析world数据异常")
		httpkt.ReturnError(w,ErrorCode.UnexpectedError,"RegisterHandler,UnexpectedError,json.Unmarshal,openid:"+reqProto.Openid)
		return
	}

	resArr := [...]MixLoginResProto{
		{
			Server: res.Server,
			Port:   res.Port,
			Secret: res.Secret,
			UID:    res.Id,
			Region: reqProto.Region,
			Openid: reqProto.Openid,
			IsNew:  res.New,
		},
	}

	httpkt.ReturnJSON(w, &resArr)
}
