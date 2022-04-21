package protocol

import (
	"bufio"
	"net"
	"time"
)

const DefaultBufferSize int = 100 * 1024

type Conn struct {
	conn   net.Conn
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		conn:   conn,
		Reader: bufio.NewReaderSize(conn, DefaultBufferSize),
		Writer: bufio.NewWriterSize(conn, DefaultBufferSize),
	}
}

func (c *Conn) Write(bytes []byte) (int, error) {
	defer c.Writer.Flush()
	c.conn.SetWriteDeadline(time.Now().Add(60 * time.Second))
	return SendResponse(c.Writer, bytes)
}

func (c *Conn) Read() ([]byte, error) {
	c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	return ReadResponse(c.Reader)
}

func (c *Conn) Addr() string {
	return c.conn.RemoteAddr().String()
}

func (c *Conn) Close() error {
	return c.conn.Close()
}
