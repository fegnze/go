package protocols

import (
	"ktnet/core/kthttp"
	"net/http"
	"net/url"
)

//ProtoTest 获取玩家角色名
type ProtoTest struct {
	//KtHttp.Proto
	Name string `json:"name"`
	Age  rune   `json:"age"`
	Sex  string `json:"sex"`
}

func (proto *ProtoTest) excute(query *url.Values) KtHttp.ProtoError {

	return kthttp.ProtoError{
		Code: 1,
		Msg:  "error",
	}
}

func (proto *ProtoTest) response(w *http.ResponseWriter) {

}
