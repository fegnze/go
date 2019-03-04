package client

import (
	"KiteNet/log"
	"KiteNet/utils"
	"encoding/json"
	"io"
	"net"
	"strconv"
)

//writeProto 写入管道协议
type writeProto struct {
	sid string
	msg []byte
}

//写入通道 writeChan
type writeChan chan writeProto

//读取通道 receiveChan
type receiveChan chan []byte

//控制信号
type controlSignal int
const (
	CloseSig controlSignal = iota
)
//controlProto 控制管道协议
type controlProto struct {
	sid string
	signal controlSignal
}
//控制通道 controlChan
type controlChan chan controlProto

//Client socket客户端会话
type Client struct{
	sess map[string]*Session
	conn net.Conn
	wch writeChan
	sysch controlChan
}

//创建一则会话
func (c *Client)NewSession()*Session{
	s := &Session{}
	s.sid = utils.CreateUniqueCode().String()
	s.wch = c.wch
	s.rch = make(receiveChan)
	s.sysch = c.sysch

	c.sess[s.sid] = s

	return s
}

//创建连接
func (c *Client)Connect(serverAddress string){
	//解析socket地址
	addr,err := net.ResolveTCPAddr("tcp4",serverAddress)
	utils.CheckErr(err)

	//建立tcp连接
	for true {
		conn,err := net.DialTCP("tcp4",nil,addr)
		if utils.CheckErr(err,"connect to world server faild!") {
			glog.Info("try again ...")
			continue
		}

		c.conn = conn
		utils.CheckNil(conn)
		glog.Info("connect to world server success!")
		break
	}

	c.sess = make(map[string]*Session)
	c.wch = make(writeChan)
	c.sysch = make(controlChan)

	go c.write()
	go c.receive()
}

//发送信息
func (c *Client) write() {
	for {
		select {
		case data := <-c.wch:
			_, err := c.conn.Write(data.msg)
			utils.CheckErr(err)
		case ctrl := <-c.sysch:
			if ctrl.signal == CloseSig {
				delete(c.sess, ctrl.sid)
			}
		}
	}
}

//回复消息的数据结构
type sidProto struct {
	Sid string `json:"sid"`
}
//监听信息
func (c *Client)receive() {
	len := 0
	response := make([]byte,1024)
	index := 0

	for{
		n,err := c.conn.Read(response[index:])

		if err != nil {
			if err != io.EOF {
				//TODO 容错
				glog.Error(err)
			}
		}

		if len > 0 {
			response = response[index:index+len]
			glog.Debug("接收信息....:",string(response))

			sp := sidProto{}
			err := json.Unmarshal(response,&sp)
			//TODO 容错
			utils.CheckErr(err)
			utils.CheckNilMsg(sp.Sid)

			c.sess[sp.Sid].rch <- response

			len = 0
			response = make([]byte,1024)
		}

		if n == 5 {

			ls := string(response[index+1:index+n])
			if l, err := strconv.ParseInt(ls, 16, 32); err == nil {
				glog.Debug("将要接收的数据长度为:",l)
				len = int(l)
			}


			index += n
		}
	}
}

//关闭连接
func (c *Client)Close(){
	c.conn.Close()
}