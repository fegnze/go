package tcp

import (
	"github.com/rs/xid"
)

//ServiceInterface -服务,用于归类逻辑模块和分发客户端请求
type ServiceInterface interface {
	BindSession(session *Session)
	Dispatch(SessionID xid.ID, data []byte)
}
