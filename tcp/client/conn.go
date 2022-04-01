package main

import (
	"net"
	"time"
)

type Conn struct {
	conn *net.TCPConn
}

func (c *Conn) Read(p []byte) (int, error) {
	c.conn.SetReadDeadline(time.Now().Add(time.Minute))
	return c.conn.Read(p)
}
