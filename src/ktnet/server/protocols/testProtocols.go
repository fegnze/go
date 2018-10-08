package protocols

import (
	"ktnet/ktcore/kthttp"
	"net/http"
	"net/url"
)

//ProtoTest 获取玩家角色名
type ProtoTest struct {
	kthttp.Proto
	Name string `json:"name"`
	Age  rune   `json:"age"`
	Sex  string `json:"sex"`
}

func (proto *ProtoTest) excute(query *url.Values) kthttp.ProtoError {

	return kthttp.ProtoError{
		Code: 1,
		Msg:  "error",
	}
}

func (proto *ProtoTest) response(w *http.ResponseWriter) {

}
