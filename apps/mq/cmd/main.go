package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/hauson/nsq-study/apps/mq/protocol"
	"github.com/hauson/nsq-study/apps/mq/pump"
)

func main() {
	tcpListener, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		panic(err)
	}
	log.Println("TCP: listening on %s", tcpListener.Addr())

	addrMsgCh := make(chan protocol.AddrMsg, 1000)
	msgPumg := pump.NewMsgPump(addrMsgCh)
	_ = msgPumg

	for {
		conn, err := tcpListener.Accept()
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				log.Printf("WARN temporary Accept() failure - %s", err)
				continue
			}

			// theres no direct way to detect this error because it is not exposed
			if !strings.Contains(err.Error(), "use of closed network connection") {
				break
			}
			break
		}

		go func() {
			c := NewConn(conn, addrMsgCh)
			_ = c
		}()
	}

	//todo:要阻塞，等所有conns 的协程结束
	select {}
	msgPumg.Close()

	log.Printf("INFO TCP: closing %s", tcpListener.Addr())
	fmt.Println("exit")
}
