package httpkt

import (
	"KiteNet/log"
	"KiteNet/utils"
	"encoding/json"
	"net/http"
)

//HandleFunc 重定义
type HandleFunc func( http.ResponseWriter, *http.Request)

func (f HandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	e := r.ParseForm()
	utils.CheckErr(e)

	f(w,r)
}

//ResInfo 数据返回格式
type ResInfo struct {
	Code    int         `json:"code"`
	Content interface{} `json:"content"`
	Msg     string      `json:"msg"`
	Client  int			`json:"client"`
}

//ReturnJSON 返回给json数据
func ReturnJSONWithClientMsg(w http.ResponseWriter, data interface{},msg string,client int) {
	glog.Debug("回传数据:", data)
	glog.Deepin(4)
	ret := &ResInfo{
		Code:    0,
		Content: data,
		Msg: msg,
		Client: client,
	}
	buf, err := json.Marshal(ret)
	if err != nil {
		ReturnError(w, -3, "数据序列化错误!")
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	i,e:=w.Write(buf)
	utils.CheckErrA(i,e)
}

//ReturnJSON 返回给json数据
func ReturnJSON(w http.ResponseWriter, data interface{}) {
	glog.Deepin(5)
	ReturnJSONWithClientMsg(w,data,"",0)
}

//ReturnError 返回给json数据
func ReturnErrorWithClientMsg(w http.ResponseWriter, code int, msg string,client int) {
	ret := &ResInfo{
		Code: code,
		Msg:  msg,
		Client: client,
	}
	buf, _ := json.Marshal(ret)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	i, e := w.Write(buf)
	glog.Deepin(4)
	utils.CheckErrA(i,e)
}

//ReturnError 返回给json数据
func ReturnError(w http.ResponseWriter, code int, msg string) {
	glog.Deepin(5)
	ReturnErrorWithClientMsg(w,code,msg,0)
}

//ReturnNull 单纯返回
func ReturnNull(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}
