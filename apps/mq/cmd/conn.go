package main

import (
	"encoding/json"
	"log"
	"net"

	"github.com/hauson/nsq-study/apps/exitsig"
	"github.com/hauson/nsq-study/apps/mq/protocol"
)

type Conn struct {
	*exitsig.ExitSig
	client     *protocol.Conn
	writeMsgCh chan protocol.Msg
}

func NewConn(conn net.Conn, addrMsgCh chan<- protocol.AddrMsg) *Conn {
	c := &Conn{
		client:     protocol.NewConn(conn),
		writeMsgCh: make(chan protocol.Msg, 1024),
		ExitSig:    exitsig.New(nil),
	}

	c.GoFunc(c.writeLoop)

	c.GoFunc(func(sig exitsig.Sig) {
		c.readLoop(sig, addrMsgCh)
	})

	return c
}

func (c *Conn) writeLoop(exitSig exitsig.Sig) {
	for {
		select {
		case msg := <-c.writeMsgCh:
			bytes, err := json.Marshal(&msg)
			if err != nil {
				continue
			}
			c.client.Write(bytes)
		case <-exitSig:
			return
		}
	}
}

func (c *Conn) readLoop(exitSig exitsig.Sig, msgCh chan<- protocol.AddrMsg) {
	for {
		select {
		case <-exitSig:
			return
		default:
			data, err := c.client.Read()
			if err != nil {
				continue
			}

			log.Println("conn receive msg", string(data))

			msg := &protocol.Msg{}
			if err := json.Unmarshal(data, msg); err != nil {
				continue
			}

			if msg.Type == protocol.HeartBeat {
				c.SendMsg(*msg)
			}

			msgCh <- protocol.AddrMsg{
				Msg:  msg,
				Addr: c.client.Addr(),
			}
		}
	}
}

// if channel full, abandon msg
func (s *Conn) SendMsg(msg protocol.Msg) {
	select {
	case s.writeMsgCh <- msg:
	default:
	}
}
