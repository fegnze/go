package gate

import (
	"github.com/rs/xid"
)

type CloseCode int

const (
	CloseNormal  CloseCode = iota
	CloseTimeOut CloseCode = iota
)

//Session
type Session struct {
	ID xid.ID

	readChan  chan []byte
	writeChan chan []byte
	sysChan   chan CloseCode
}

//Read
func (ss *Session) Read(msg []byte) {
	proxy.DispatchMsg(ss.ID, msg)
}

//Write
func (ss *Session) Write(msg []byte) {
	ss.writeChan <- msg
}

//Close
func (ss *Session) Close(code CloseCode) {
	ss.sysChan <- code
}
