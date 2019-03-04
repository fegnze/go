package main

import (
	"KiteNet/tcp"
	"fmt"
	"time"
)

//Player -
type Player struct {
	Session *tcp.Session
}

//SendMsgToSession -
func (player *Player) SendMsgToSession(data []byte) {
	player.Session.SendMessage(data)
}

//Setup -
func (player *Player) Setup(data []byte) {
	fmt.Println("Player setup ...", string(data))
	time.Sleep(time.Second * 5)
	player.SendMsgToSession([]byte("你好,我是服务器"))
}
