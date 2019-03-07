package wsocket

import (
	"KiteNet/log"
	"KiteNet/utils"
	"github.com/rs/xid"
	"golang.org/x/net/websocket"
	"net/http"
)

//SysType -控制信号类型
type SysType int

//noinspection GoUnusedConst
const (
	SESSIONSYSTYPECLOSE SysType = iota
	SESSIONSYSTYPEERROR
)

type SessChanMsg struct {
	sessID  xid.ID
	sysType SysType
	msg     string
}

var sess map[xid.ID]*Session
var proxy ProxyInterface

//wsHandle
func wsHandle(conn *websocket.Conn) {
	request := make([]byte, 128)
	defer conn.Close()

	id := utils.CreateUniqueCode()
	s := &Session{
		ID: id,
	}
	sess[id] = s
	proxy.BindSession(s)

	go func(session *Session, c *websocket.Conn) {
	R:
		for {
			select {
			case code := <-session.sysChan:
				if code == 1 {
					err := c.Close()
					utils.CheckErr(err)
					break R
				}
				//TODO
			case data := <-session.writeChan:
				n, err := c.Write(data)
				utils.CheckNilAndErr(n, err)
			}
		}
	}(s, conn)

	for {
		readlen, err := conn.Read(request)
		if utils.CheckErr(err) {
			break
		}

		if readlen == 0 {
			glog.Info("Client connection close!")
			s.Close(CloseNormal)
			break
		} else {
			glog.Debug(string(request[:readlen]))
			s.Read(request[:readlen])
		}

		request = make([]byte, 128)
	}

	proxy.BindSession(s)
	delete(sess, id)
}

//Listen
func Listen(addr string, path string, p ProxyInterface) {
	if path == "" {
		path = "ws"
	}

	sess = make(map[xid.ID]*Session)
	proxy = p

	http.Handle("/"+path, websocket.Handler(wsHandle))
	err := http.ListenAndServe(addr, nil)
	if utils.CheckErr(err) {
		panic(err)
	}
}
