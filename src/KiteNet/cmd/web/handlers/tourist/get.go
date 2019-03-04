package tourist

import (
	"KiteNet/cmd/web/db"
	"KiteNet/cmd/web/global/code/ErrorCode"
	"KiteNet/httpkt"
	"KiteNet/log"
	"KiteNet/utils"
	"fmt"
	"net/http"
)

//GetReqProto 客户端请求数据,本地类型
type GetTouristReqProto struct {
	SubPackage string `json:"openid"`
	Region     string `json:"region"`
	UniqueCode string `json:"uniqueCode"`
	DeviceType string `json:"deviceType"`
	Origin     string `json:"origin"`
	DeviceId   string `json:"deviceId"`
	Channel    string `json:"channel"`
}

//GetTouristResProto 返回数据
type GetTouristResProto struct {
	New        string `json:"new"` //服务器IP
	Openid     string `json:"openid"`
	Nickname   string `json:"nickname"` //秘钥
	Account    string `json:"account"`
	RegionName string `json:"regionName"`
	RegionId   string `json:"regionId"`
}

//GetTouristHandler login/mix处理函数
func GetTouristHandler(w http.ResponseWriter, r *http.Request) {
	glog.Debug("mix/login/tourist/get HttpExec ...")
	reqProto := &GetTouristReqProto{}
	httpkt.ParseReqProto(r, reqProto)

	//TODO -此处是否存在无限索引风险?如何优化?
	//生成UniqueCode(UUID)
	openid := utils.CreateUniqueCode().String()
	for db.GetUserWithOpenID(openid) == nil {
		openid = utils.CreateUniqueCode().String()
	}

	//生成随机字符串
	account := utils.CreateRandString(5)
	account = fmt.Sprintf("%s-%s", reqProto.Channel, account)

	//将该条数据插入到数据库
	err := db.InsertUserAccount(account, "", "", openid)
	if err != nil {
		glog.Error(err)
		httpkt.ReturnError(w,ErrorCode.DBUnexpectedError,"GetTouristHandler,db.InsertUserAccount,DBUnexpectedError,openid:"+openid)
	}

	res := GetTouristResProto{
		New:        "1", //0-- 老用户,1-- 新用户,2-- 已经绑定过正式账号的游客用户
		Openid:     openid,
		Nickname:   account,
		Account:    account,
		RegionName: account,
		RegionId:   reqProto.Region,
	}

	httpkt.ReturnJSON(w, &res)
}
