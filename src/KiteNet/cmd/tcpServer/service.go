package main

import (
	"KiteNet/tcp"
	"fmt"

	"github.com/rs/xid"
)

//AgentService -
type AgentService struct {
	svmap  map[xid.ID]*Player
	buffer []byte
}

//BindSession -
func (agent *AgentService) BindSession(session *tcp.Session) {
	if agent.svmap == nil {
		agent.svmap = make(map[xid.ID]*Player, 0)
	}
	if agent.svmap[session.HD] != nil {
		fmt.Println("Error:session已存在!")
	}
	agent.svmap[session.HD] = &Player{
		Session: session,
	}
}

//Dispatch -
func (agent *AgentService) Dispatch(sessionID xid.ID, data []byte) {
	if agent.buffer == nil {
		agent.buffer = make([]byte, 0)
	}
	if len(agent.buffer) > 0 {
		agent.buffer = append(agent.buffer, data...)
	} else {

	}
	agent.svmap[sessionID].Setup(data)
}
