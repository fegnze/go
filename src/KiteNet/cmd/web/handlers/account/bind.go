package account

//游客登录时的账号绑定
import (
	"KiteNet/cmd/web/db"
	"KiteNet/cmd/web/global/code/ErrorCode"
	"KiteNet/httpkt"
	"KiteNet/log"
	"net/http"
)

//bindReqProto 游客登录时的账号绑定,客户端请求数据,本地类型
type bindReqProto struct {
	Openid     string `json:"openid"`
	Account    string `json:"account"`
	Password   string `json:"password"`
	Channel    string `json:"channel"`
	DeviceId   string `json:"deviceId"`
	DeviceType string `json:"deviceType"`
	SubPackage string `json:"subPackage"`
}

//bindResProto 游客登录时的账号绑定,返回数据
type bindResProto struct {
	Openid string `json:"openid"`
}

//BindHandler 游客登录时的账号绑定,处理函数
func BindHandler(w http.ResponseWriter, r *http.Request) {
	glog.Debug("BindHandler ... 游客账号登陆时账号绑定")
	reqProto := &bindReqProto{}
	httpkt.ParseReqProto(r, reqProto)
	//验证下账号密码格式
	if reqProto.Account == "" {
		httpkt.ReturnError(w, ErrorCode.ParameterNull, "open id is null")
	}
	//更新数据库
	err := db.UpdateUserAccount(reqProto.Account, reqProto.Password, "", reqProto.Openid)
	if err != nil {
		glog.Error(err)
		httpkt.ReturnError(w, ErrorCode.DBUnexpectedError, "db.UpdateUserAccount,DBUnexpectedError,openid:"+reqProto.Openid)
		return
	}

	//TODO -写入返回消息
	httpkt.ReturnJSON(w, &bindResProto{
		Openid: reqProto.Openid,
	})
}
