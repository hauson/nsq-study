package protocol

import (
	"bufio"
	"net"
	"time"
)

const DefaultBufferSize int = 100 * 1024

type Client struct {
	conn   net.Conn
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn:   conn,
		Reader: bufio.NewReaderSize(conn, DefaultBufferSize),
		Writer: bufio.NewWriterSize(conn, DefaultBufferSize),
	}
}

func (c *Client) Write(bytes []byte) (int, error) {
	defer c.Writer.Flush()
	c.conn.SetWriteDeadline(time.Now().Add(30 * time.Second))
	return SendResponse(c.Writer, bytes)
}

func (c *Client) Read() ([]byte, error) {
	c.conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	return ReadResponse(c.Reader)
}

func (c *Client) Addr() string {
	return c.conn.RemoteAddr().String()
}

func (c *Client) Close() error {
	return c.conn.Close()
}
