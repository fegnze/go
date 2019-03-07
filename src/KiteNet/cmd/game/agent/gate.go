package agent

import (
	"KiteNet/wsocket/gate"
	"github.com/rs/xid"
	"time"
)

//ProxyService
type ProxyService struct {
	agentMap map[xid.ID]*Agent
}

var Proxy ProxyService

func init() {
	Proxy = ProxyService{
		agentMap: make(map[xid.ID]*Agent),
	}
}

//BindSession 绑定session
func (s *ProxyService) BindSession(session *gate.Session) {
	if s.agentMap[session.ID] == nil {
		s.agentMap[session.ID] = &Agent{
			session:        session,
			lastReciveTime: time.Now(),
		}
	}
}

//DispatchMsg 分发消息
func (s *ProxyService) DispatchMsg(sessionID xid.ID, data []byte) {
	//TODO 是否会阻塞协程
	s.agentMap[sessionID].Recive(data)
}

//SendMessage 发送数据给客户端
func (s *ProxyService) UnBindSession(sessionID xid.ID) {
	s.agentMap[sessionID].Release()
	delete(s.agentMap, sessionID)
}
