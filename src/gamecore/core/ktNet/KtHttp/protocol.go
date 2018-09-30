package KtHttp

import (
	"gamecore/core/KtNet/KtHttp/11111"
	"net/http"
	"net/url"
)

//Proto 获取玩家角色名
type Proto struct {
}

func (proto *Proto) excute(query *url.Values) ProtoError {

	return ProtoError{
		Code: 1,
		Msg:  "error",
	}
}

func (proto *Proto) response(w *http.ResponseWriter) {
	test.Test()
}
