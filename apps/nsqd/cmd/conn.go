package main

import (
	"encoding/json"
	"github.com/hauson/nsq-study/apps/nsqd/protocol"
	"log"
	"net"
	"sync"
)

type Conn struct {
	client *protocol.Client

	writeMsgCh chan protocol.Msg
	exitSig    chan interface{}
	wg         sync.WaitGroup
}

func NewConn(conn net.Conn, addrMsgCh chan<- protocol.AddrMsg) *Conn {
	c := &Conn{
		client:     protocol.NewClient(conn),
		writeMsgCh: make(chan protocol.Msg, 1024),
		exitSig:    make(chan interface{}, 1),
	}

	c.wg.Add(2)
	go c.readLoop(addrMsgCh)
	go c.writeLoop()

	return c
}

func (c *Conn) writeLoop() {
	defer c.wg.Done()

	for {
		select {
		case msg := <-c.writeMsgCh:
			bytes, err := json.Marshal(&msg)
			if err != nil {
				continue
			}
			c.client.Write(bytes)
		case <-c.exitSig:
			return
		}
	}
}

func (c *Conn) readLoop(msgCh chan<- protocol.AddrMsg) {
	defer c.wg.Done()

	for {
		select {
		case <-c.exitSig:
			return
		default:
			data, err := c.client.Read()
			if err != nil {
				continue
			}

			log.Println("receive msg", string(data))

			originMsg := &protocol.Msg{}
			if err := json.Unmarshal(data, originMsg); err != nil {
				continue
			}

			msgCh <- protocol.AddrMsg{
				Msg:  originMsg,
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

// sync block
func (s *Conn) Close() {
	close(s.exitSig)
	s.wg.Wait()
}
