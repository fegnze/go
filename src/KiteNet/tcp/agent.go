package tcp

import (
	"KiteNet/log"
	"KiteNet/utils"
	"net"

	"github.com/rs/xid"
)

//SysType -控制信号类型
type SysType int

//noinspection GoUnusedConst
const (
	SESSIONSYSTYPECLOSE SysType = iota
	SESSIONSYSTYPEERROR
)

//SysChanMsg -控制信号消息
type SysChanMsg struct {
	sessID  xid.ID
	sysType SysType
	msg     string
}

//Agent -
type Agent struct {
	Addr        *net.TCPAddr
	Sess        map[xid.ID]*Session
	sessSysChan chan SysChanMsg

	Service ServiceInterface
}

//Start -
func (agent *Agent) Start(addr string) {
	//解析地址
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	utils.CheckErr(err)
	utils.CheckErr(err)
	agent.Addr = tcpAddr
	//设置监听
	listener, err := net.ListenTCP("tcp", tcpAddr)
	utils.CheckErr(err)

	agent.sessSysChan = make(chan SysChanMsg)

	//等待连接
	go func() {
		for {
			glog.Debug("Wait connect ...")
			conn, err := listener.Accept()
			if err != nil {
				glog.Error(err)
				continue
			}
			id := xid.New()
			sess := &Session{
				Service: agent.Service,
				HD:      id,
				SysChan: agent.sessSysChan,
				Conn:    conn,
			}
			//绑定到具体的逻辑实体中
			agent.Service.BindSession(sess)
			go sess.Connect()
		}
	}()

	//接收控制信号
	select {
	case sysMsg := <-agent.sessSysChan:
		glog.Debug(sysMsg.sessID, sysMsg.msg)
	}
}
