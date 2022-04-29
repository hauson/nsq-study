package main

import (
	"net"
	"time"
)

// MagicV2 is the initial identifier sent when connecting for V2 clients
var MagicV2 = []byte("  V2")

type Conn struct {
	conn *net.TCPConn
}

func (c *Conn) Read(p []byte) (int, error) {
	c.conn.SetReadDeadline(time.Now().Add(time.Minute))
	return c.conn.Read(p)
}

// Write performs a deadlined write on the underlying TCP connection
func (c *Conn) Write(p []byte) (int, error) {
	c.conn.SetWriteDeadline(time.Now().Add(1 * time.Minute))
	return c.conn.Write(p)
}
