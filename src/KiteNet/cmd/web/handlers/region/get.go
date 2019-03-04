package region

import (
	"KiteNet/cmd/web/db"
	"KiteNet/httpkt"
	"KiteNet/log"
	"net/http"
)

//GetReqProto 客户端请求数据,本地类型
type GetReqProto struct{}

//GetResProto 返回数据
type GetResProto struct {
	RegionID    int    `json:"regionId"`
	RegionName  string `json:"regionName"`
	RegionState int    `json:"regionState"`
	ServerIP    string `json:"serverIp"`
	ServerPort  string `json:"serverPort"`
}

//GetHandler 处理函数
func GetHandler(w http.ResponseWriter, r *http.Request) {
	glog.Debug("region version HttpExec ...")
	//reqProto := &GetReqProto{}
	//httpkt.ParseReqProto(r, reqProto)
	//TODO -执行服务逻辑
	list := db.GetRegionList()
	//TODO -写入返回消息
	region := GetResProto{
		RegionID:    list[0].ID,
		RegionName:  list[0].Name,
		RegionState: list[0].State,
		//ServerIP:    "127.0.0.1",
		//ServerPort:  "15443",
	}
	httpkt.ReturnJSON(w, &region)
}
