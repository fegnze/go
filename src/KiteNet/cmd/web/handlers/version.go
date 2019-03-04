package handlers

//VersionHandle
import (
	"KiteNet/cmd/web/db"
	"KiteNet/httpkt"
	"KiteNet/log"
	"fmt"
	"net/http"
	"os"
)

//versionReqProto 客户端请求数据,本地类型
//type versionReqProto struct{}

//versionProto 版本信息回传数据
type versionProto struct {
	Version   string   `json:"version"`   //版本号
	URL       string   `json:"url"`       //网站地址
	UpdateURL string   `json:"updateUrl"` //更新地址
	Cdn       string   `json:"cdn"`       //cdn地址
	Cdns      []string `json:"cdns"`      //各渠道cdn地址
	Active    string   `json:"active"`    //活动地址
	PayURL    string   `json:"payUrl"`    //支付地址
	Notice    string   `json:"notice"`    //支付通知地址
	Static    string   `json:"static"`    //静态数据地址
	Build     string   `json:"build"`     //IOS版本
}

//VersionHandler 处理函数
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	glog.Debug("VersionHandler HttpExec ...")
	v := db.VersionData{}
	if err := db.GetGameVersion(&v); err != nil {
		glog.Error(err,"查询版本信息异常")
		//TODO 此处应该返回一个错误码
	}
	ver := v.Test
	debug := os.Getenv("DEBUG_MODLE")
	if debug == "false" {
		ver = v.Release
	}
	host := os.Getenv("SERVER_HOST")
	httpkt.ReturnJSON(w, &versionProto{
		Version:   ver,
		URL:       fmt.Sprintf("http://%s/worldship", host),
		Cdn:       fmt.Sprintf("http://%s/download", host),
		UpdateURL: fmt.Sprintf("http://%s/download", host),
		Cdns:      []string{fmt.Sprintf("http://%s/download", host), fmt.Sprintf("http://%s/download", host)},
		Active:    fmt.Sprintf("http://%s/worldship", host),
		PayURL:    fmt.Sprintf("http://%s/worldship", host),
		Notice:    fmt.Sprintf("http://%s/worldship", host),
		Static:    fmt.Sprintf("http://%s/worldship", host),
		Build:     "0",
	})
}
