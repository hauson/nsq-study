package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/hauson/nsq-study/tcp/protocol"
)

//inputs: 127.0.0.1:4151, topic, channel
func main() {
	dialer := &net.Dialer{
		LocalAddr: nil,
		Timeout:   1 * time.Second,
	}

	conn, err := dialer.Dial("tcp", "127.0.0.1:7151")
	if err != nil {
		panic(err)
	}

	c := &Conn{
		conn: conn.(*net.TCPConn),
	}

	if _, err := c.Write(MagicV2); err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			time.Sleep(10 * time.Second)
			if _, err := c.Write([]byte("request: hello")); err != nil {
				panic(err)
			}
		}
	}()
	// 2. 发送一个版本消息
	// 开通一个 readLoop()
	go func() {
		defer wg.Done()
		for {
			n, bytes, err := protocol.ReadUnpackedResponse(c)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(n, string(bytes))
			}
		}
	}()

	// 开通一个 writeLoop()
	// 等待处理
	wg.Wait()
	fmt.Println("sucess")
}

// 订阅主题和channel
