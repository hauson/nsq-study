package main

import (
	"io"
	"log"
	"net"
	"sync"

	"github.com/hauson/nsq-study/demo-tcp/protocol"
)

type ClientStats interface {
	String() string
}

type Client interface {
	Type() int
	Stats(string) ClientStats
}

type tcpServer struct {
	// remoteAddr-> Client
	conns sync.Map
}

func (p *tcpServer) Handle(conn net.Conn) {
	buf := make([]byte, 4)
	// 一个阻塞等待的行为
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		conn.Close()
		return
	}

	protocolMagic := string(buf)
	var prot protocol.Protocol
	switch protocolMagic {
	case "  V2":
		prot = &protocolV2{}
	default:
		const frameTypeError = 1
		protocol.SendFramedResponse(conn, frameTypeError, []byte("E_BAD_PROTOCOL"))
		conn.Close()
		return
	}

	client := prot.NewClient(conn)
	p.conns.Store(conn.RemoteAddr(), client)

	err = prot.IOLoop(client)
	if err != nil {
		log.Printf("WARN ")
	}

	p.conns.Delete(conn.RemoteAddr())
	client.Close()
}
