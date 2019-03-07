package agent

import (
	"KiteNet/cmd/game/module"
	"KiteNet/cmd/game/protos"
	"KiteNet/utils"
	"KiteNet/wsocket/gate"
	"github.com/golang/protobuf/proto"
	"github.com/rs/xid"
	"sync"
	"time"
)

//agent
type Agent struct {
	session *gate.Session
	player  *module.Player

	lastReciveTime  time.Time
	tickerLock      sync.Mutex
	tickerCloseChan chan int
}

func (agent *Agent) Recive(data []byte) {
	agent.tickerLock.Lock()
	agent.lastReciveTime = time.Now()
	agent.tickerLock.Unlock()

	len := data[:2]
	protoType := data[2:6]
	msgID := data[6:8]
	msgData := data[8:]

	if agent.pid == xid.NilID() {
		loginData := &protos.Login{}
		err := proto.Unmarshal(data, loginData)
		utils.CheckErr(err)

		agent.Connect()
	}

}

func (agent *Agent) Send(data []byte) {
	agent.session.Write(data)
}

func (agent *Agent) DisConnet(code gate.CloseCode) {
	agent.session.Close(code)
}

func (agent *Agent) Connect() {

	ticker := time.NewTicker(time.Second * 10)
	go agent.HeartBeat(ticker)
}

func (agent *Agent) HeartBeat(ticker *time.Ticker) {
R:
	for {
		select {
		case <-ticker.C:
			var interval int64
			agent.tickerLock.Lock()
			interval = time.Now().Unix() - agent.lastReciveTime.Unix()
			agent.tickerLock.Unlock()

			if interval > 10 {
				agent.DisConnet(gate.CloseTimeOut)
			}
		case <-agent.tickerCloseChan:
			break R
		}
	}
}

func (agent *Agent) Release() {
	agent.session = nil
}
