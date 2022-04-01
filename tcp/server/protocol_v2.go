package main

import (
	"github.com/hauson/nsq-study/tcp/protocol"
	"net"
)

type protocolV2 struct {
}

func (v2 *protocolV2) NewClient(net.Conn) protocol.Client {
	return nil
}

func (v2 *protocolV2) IOLoop(protocol.Client) error {
	return nil
}
