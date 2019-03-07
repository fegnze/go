package login

import (
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
	pid     xid.ID

	lastReciveTime  time.Time
	tickerLock      sync.Mutex
	tickerCloseChan chan int
}

func (agent *Agent) Recive(data []byte) {
	agent.tickerLock.Lock()
	agent.lastReciveTime = time.Now()
	agent.tickerLock.Unlock()

	if agent.pid == xid.NilID() {
		loginData := &Login{}
		err := proto.Unmarshal(data, loginData)
		utils.CheckErr(err)

		agent.Connect()
	}

}

func (agent *Agent) Send(data []byte) {
	agent.session.Write(data)
}

func (agent *Agent) DisConnet() {

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
				agent.DisConnet()
			}
		case <-agent.tickerCloseChan:
			break R
		}
	}
}
