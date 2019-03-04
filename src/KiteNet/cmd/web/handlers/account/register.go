package account

//Register 自有账号注册
import (
	"KiteNet/cmd/web/db"
	"KiteNet/cmd/web/global/code/ErrorCode"
	"KiteNet/httpkt"
	"KiteNet/log"
	"KiteNet/utils"
	"net/http"
)

//registerReqProto 自有账号注册,客户端请求数据,本地类型
type registerReqProto struct {
	Account    string `json:"account"`
	Password   string `json:"password"`
	Channel    string `json:"channel"`
	DeviceId   string `json:"deviceId"`
	DeviceType string `json:"deviceType"`
	SubPackage string `json:"subPackage"`
}

//registerResProto 自有账号注册,返回数据
type registerResProto struct {
	OpenId string `json:"openid"`
}

//RegisterHandler 自有账号注册,处理函数
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	glog.Debug("RegisterHandler ... 自有账号注册")
	reqProto := &registerReqProto{}
	httpkt.ParseReqProto(r, reqProto)
	//检查参数是否有效
	if reqProto.Account == "" {
		httpkt.ReturnError(w,ErrorCode.ParameterNull,"account is null")
	}
	//TODO -此处是否存在无限索引风险?如何优化?
	//生成UniqueCode(UUID)
	openid := utils.CreateUniqueCode().String()
	for db.GetUserWithOpenID(openid) == nil {
		openid = utils.CreateUniqueCode().String()
	}

	err := db.InsertUserAccount(reqProto.Account,reqProto.Password,"",openid)
	if err != nil {
		glog.Error(err)
		httpkt.ReturnError(w,ErrorCode.DBUnexpectedError,"RegisterHandler,db.InsertUserAccount,DBUnexpectedError,openid:"+openid)
	}

	httpkt.ReturnJSON(w, &registerResProto{OpenId:openid})
}
