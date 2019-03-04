package client

import "KiteNet/utils"

//Session socket客户端会话
type Session struct{
	sid string
	wch writeChan //client 写入管道
	rch receiveChan //session 私有接收管道
	sysch controlChan //client 控制管道
}

//发送信息
func (s *Session)Write(msg []byte)[]byte{
	utils.CheckNilMsg(s.sid)
	utils.CheckNil(s.rch)
	utils.CheckNil(s.wch)
	utils.CheckNil(s.sysch)

	//向服务器发送数据
	msg = append(msg,byte(','))
	msg = append(msg,[]byte(s.sid)...)
	msg = append(msg,byte('\n'))
	s.wch <- writeProto{
		sid: s.sid,
		msg: msg,
	}

	return <- s.rch
}

//关闭连接
func (s *Session)Close(){
	if s.sysch != nil {
		s.sysch <- controlProto{
			sid:    s.sid,
			signal: CloseSig,
		}
	}
}