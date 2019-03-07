package gate

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
var sessSysChan chan SessChanMsg
var service ServiceInterface

//wsHandle
func wsHandle(conn *websocket.Conn) {
	request := make([]byte, 128)
	defer conn.Close()

	s := &Session{
		ID: utils.CreateUniqueCode(),
	}
	sess[id] = s
	service.BindSession(s)

	for {
		readlen, err := conn.Read(request)
		if utils.CheckErr(err) {
			break
		}

		if readlen == 0 {
			glog.Info("Client connection close!")
			break
		} else {
			glog.Debug(string(request[:readlen]))
			s.Recive(request[:readlen])
		}

		request = make([]byte, 128)
	}
}

//Listen
func Listen(addr string, path string, s ServiceInterface) {
	if path == "" {
		path = "ws"
	}

	sess = make(map[xid.ID]*Session)
	sessSysChan = make(chan SessChanMsg)
	service = s

	http.Handle("/"+path, websocket.Handler(wsHandle))
	err := http.ListenAndServe(addr, nil)
	if utils.CheckErr(err) {
		panic(err)
	}
}
