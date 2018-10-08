package kthttp

import (
	"net/http"
	"net/url"
)

//Proto 获取玩家角色名
type Proto struct {
}

func (proto *Proto) excute(query *url.Values) ProtoError {

	return ProtoError{
		Code: -100,
		Msg:  "base protocol",
	}
}

func (proto *Proto) response(w *http.ResponseWriter) {

}
