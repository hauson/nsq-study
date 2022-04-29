package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/hauson/nsq-study/demo-tcp/protocol"
)

const defaultBufferSize int = 16 * 1024

type clientV2 struct {
	conn   net.Conn
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func (c *clientV2) Close() error {
	return c.conn.Close()
}

type protocolV2 struct {
}

func (v2 *protocolV2) NewClient(conn net.Conn) protocol.Client {
	return &clientV2{
		conn:   conn,
		Reader: bufio.NewReaderSize(conn, defaultBufferSize),
		Writer: bufio.NewWriterSize(conn, defaultBufferSize),
	}
}

func (v2 *protocolV2) IOLoop(c protocol.Client) error {
	// 接受消息，然后回复
	client := c.(*clientV2)
	for {
		buf := make([]byte, 1024)
		client.conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		n, err := client.conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				err = nil
			} else {
				err = fmt.Errorf("failed to read command - %s", err)
			}
			break
		}

		fmt.Println(string(buf[:n]), time.Now().String())
		//protocol.SendResponse(client.conn, []byte("response:, hello, world"))
		protocol.SendResponse(client.Writer, []byte("response:, hello, world"))
		client.Writer.Flush()
	}

	goto exit
exit:
	return nil
}
