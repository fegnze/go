package ktsocket

import (
	"ktnet/ktcore/ktlog"
	"net"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		//read from the connection
		//write to then connection
	}
}

func connect(addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		ktlog.Info("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			ktlog.Info("accept error:", err)
			break
		}

		//start a new goroutine to handle
		//the new connection
		go handleConn(c)
	}
}
