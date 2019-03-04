package httpkt

import (
	"KiteNet/utils"
	"encoding/json"
	"net/http"
)

//ParseReqProto 解析request数据到自定义协议中
func ParseReqProto(req *http.Request, proto interface{}) {
	if req.Form == nil {
		return
	}

	hasValue := false
	m := make(map[string]interface{}, 0)
	for k, v := range req.Form {
		m[k] = v[0]
		hasValue = true
	}
	if !hasValue {
		return
	}

	if j, err := json.Marshal(m); err == nil {
		err := json.Unmarshal(j, proto)
		utils.CheckErr(err)
	}
}
