package wsocket

import "github.com/rs/xid"

//ProxyInterface -服务,用于归类逻辑模块和分发客户端请求
type ProxyInterface interface {
	BindSession(session *Session)
	UnBindSession(sessionID xid.ID)
	DispatchMsg(SessionID xid.ID, data []byte)
}
