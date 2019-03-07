package gate

import (
	"KiteNet/utils"
	"github.com/rs/xid"
	"golang.org/x/net/websocket"
)

//Session
type Session struct {
	ID   xid.ID
	Conn *websocket.Conn
}

//Recive
func (ss *Session) Recive(msg []byte) {
	service.DispatchMsg(ss.ID, msg)
}

//Write
func (ss *Session) Write(msg []byte) {
	n, err := ss.Conn.Write(msg)
	utils.CheckNilAndErr(n, err)
}

//Close
func (ss *Session) Close() {
	err := ss.Conn.Close()
	utils.CheckErr(err)
}
