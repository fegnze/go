package tcp

import (
	"KiteNet/log"
	"KiteNet/utils"
	"io"
	"net"

	"github.com/rs/xid"
)

//noinspection GoUnusedConst
const (
	maxDataSize int = 2048
	oneDataSize int = 512
)

//sessionStateType -会话连接状态
type sessionStateType int

//IDLE,CONNECTING,DISCONNECT,INERROR
//noinspection GoUnusedConst
const (
	IDLE sessionStateType = iota
	CONNECTING
	DISCONNECT
	INERROR
)

//protoHead -
//var protoHead = "zhh7107"
//var protoHeadSize = len(protoHead)

//HeadProto -
type HeadProto struct {
	Head     string `json:"head"`
	ProtoID  int    `json:"protoID"`
	DataSize int    `json:"dataSize"`
	Data     []byte `json:"data"`
}

//Session -会话
type Session struct {
	//HD 句柄
	HD xid.ID
	//WriteChan 写入管道
	WriteChan chan []byte
	//ReadChan 读取管道
	ReadChan chan []byte
	//State 连接状态
	State sessionStateType
	//Conn socket连接
	Conn net.Conn
	//控制信号
	SysChan chan SysChanMsg
	//保存的service指针
	Service ServiceInterface
}

//Connect 创建会话
func (ss *Session) Connect() {
	if ss.Service == nil || ss.Conn == nil || ss.SysChan == nil || ss.HD == xid.NilID() {
		return
	}

	ss.State = CONNECTING
	//等待数据
	glog.Debug("client %v connect success.\n", ss.HD)
	//TODO 关于连接超时的处理
	// ss.Conn.SetReadDeadline(time.Now().Add(time.Second * 1000))
	defer ss.Close()

	request := make([]byte, 24)
	for {
		readline, err := ss.Conn.Read(request)

		//socket被关闭
		if err != nil {
			if err == io.EOF {
				glog.Debug("读取到结尾EOF")
			} else if net.Error.Timeout(err.(net.Error)) {
				glog.Error(err,"连接超时")
				//TODO 异常处理
				break
			} else if net.Error.Temporary(err.(net.Error)) {
				glog.Error(err,"临时报告,DNS异常")
				//TODO 异常处理
				//临时报告是否已知DNS错误是临时的。 这并不总是知道的; 由于临时错误，DNS查找可能会失败，并返回临时返回false的DNSError。
				break
			}
		} else {
			glog.Debug(string(request[:readline]))
			// time.Sleep(time.Second)
			//TODO 超出容量怎么办
			// if len(data)+oneDataSize > maxDataSize {
			// 	glog.Debug("单个数据包超过最大值")
			// } else {
			// 	data = append(data, request...)
			// }

			ss.Service.Dispatch(ss.HD, request)
			request = make([]byte, oneDataSize)
		}
	}
}

//Close 断开连接
func (ss *Session) Close() {
	e := ss.Conn.Close()
	utils.CheckErr(e)
	ss.SysChan <- SysChanMsg{
		sessID:  ss.HD,
		sysType: SESSIONSYSTYPEERROR,
		msg:     "会话终止",
	}
}

//SendMessage 发送消息
func (ss *Session) SendMessage(msg []byte) error {
	n, e := ss.Conn.Write(msg)
	utils.CheckErrA(n, e)
	return nil
}
