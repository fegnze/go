package sdk

//QuickSdkLogin quick sdk 登陆
import (
	"KiteNet/cmd/web/global"
	"KiteNet/cmd/web/global/code/ErrorCode"
	"KiteNet/httpkt"
	"KiteNet/log"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

//reqProto quick sdk 登陆,客户端请求数据,本地类型
type reqProto struct {
	Region     string `json:"region"`
	Channel    string `json:"channel"`
	UserId     string `json:"userId"`
	UserName   string `json:"userName"`
	Token      string `json:"token"`
	DeviceId   string `json:"deviceId"`
	DeviceType string `json:"deviceType"`
}

//resProto quick sdk 登陆 返回数据
type resProto struct {
	Server string `json:"server"`
	Port   string `json:"port"`
	Secret string `json:"secret"`
	Uid    string `json:"uid"`
	Region string `json:"region"`
	Openid string `json:"openid"`
	IsNew  int    `json:"isnew"`
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

//QuickSdkLoginHandler quick sdk 登陆,处理函数
func QuickSdkLoginHandler(w http.ResponseWriter, r *http.Request) {
	glog.Debug("QuickSdSdkLoginHandler ... quick sdk 登陆")
	reqProto := &reqProto{}
	httpkt.ParseReqProto(r, reqProto)

	openID := reqProto.UserId
	if openID == "" || reqProto.Region == "" || reqProto.Channel == "" {
		httpkt.ReturnError(w,ErrorCode.ParameterNull,
			"QuickSdkLoginHandler,ParameterNull,openid:"+openID+",reqProto.Region:"+reqProto.Region+",reqProto.Channel:"+reqProto.Channel)
	}

	//发送请求到world服务器
	req := fmt.Sprintf("ul,android_cn_999_%s_%s,%s,%d,%s", openID, reqProto.Region, reqProto.Channel, 0, "0")

	session := global.WorldSS.NewSession()
	defer session.Close()
	response := session.Write([]byte(req))
	glog.Debug("world ul", string(response[0:]))

	//TODO -写入返回消息
	res := MixLoginWorldResProto{}
	err := json.Unmarshal(response[0:], &res)
	if err != nil {
		glog.Error(err, "MixLoginHandler 解析world数据异常")
		httpkt.ReturnError(w, ErrorCode.UnexpectedError, "QuickSdkLoginHandler,UnexpectedError,json.Unmarshal,openid:"+openID)
		return
	}

	isnew,err := strconv.Atoi(res.New)
	if err != nil {
		isnew = -1
	}

	resArr := []resProto{
		{
			Server: res.Server,
			Port:   res.Port,
			Secret: res.Secret,
			Uid:    res.Id,
			Region: reqProto.Region,
			Openid: openID,
			IsNew:  isnew,
		},
	}

	//TODO -写入返回消息
	httpkt.ReturnJSON(w, &resArr)
}
