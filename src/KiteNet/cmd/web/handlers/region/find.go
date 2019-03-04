package region

import (
	"KiteNet/cmd/web/db"
	"KiteNet/httpkt"
	"KiteNet/log"
	"net/http"
)

//FindReqProto 客户端请求数据,本地类型
type FindReqProto struct{}

//FindResProto 返回数据
type FindResProto struct {
	RegionID    int    `json:"regionId"`
	RegionName  string `json:"regionName"`
	RegionState int    `json:"regionState"`
	ServerIP    string `json:"serverIp"`
	ServerPort  string `json:"serverPort"`
}

//FindHandler 处理函数
func FindHandler(w http.ResponseWriter, r *http.Request) {
	glog.Debug("regn version HttpExec ...")

	list := db.GetRegionList()
	//TODO -写入返回消息
	//ip := os.Getenv("WORLD_IP")
	//port := os.Getenv("WORLD_PORT")
	array := []FindResProto{
		{
			RegionID:    list[0].ID,
			RegionName:  list[0].Name,
			RegionState: list[0].State,
			//ServerIP:    "127.0.0.1",
			//ServerPort:  "15443",
		},
		{
			RegionID:    list[1].ID,
			RegionName:  list[1].Name,
			RegionState: list[1].State,
			//ServerIP:    "127.0.0.1",
			//ServerPort:  "15443",
		},
	}
	httpkt.ReturnJSON(w, &array)
}
