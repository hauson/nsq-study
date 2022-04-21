package client

import (
	"encoding/json"
	"github.com/hauson/nsq-study/apps/exitsig"
	"github.com/hauson/nsq-study/apps/mq/protocol"
	"net"
	"time"
)

type Client struct {
	exitSig *exitsig.ExitSig
	*protocol.Conn
}

func New(conn net.Conn) *Client {
	c := &Client{
		exitSig: exitsig.New(nil),
		Conn:    protocol.NewConn(conn),
	}

	c.exitSig.GoFunc(c.heartBeatLoop)
	return c
}

func (c *Client) heartBeatLoop(sig exitsig.Sig) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			data, _ := json.Marshal(&protocol.Msg{Type: protocol.HeartBeat})
			c.Write(data)
		case <-sig:
			return
		}
	}
}
