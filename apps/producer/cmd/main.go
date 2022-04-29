package main

import (
	"encoding/json"
	"fmt"
	"github.com/hauson/nsq-study/apps/frame/client"
	"log"
	"net"
	"time"

	"github.com/hauson/nsq-study/apps/mq/protocol"
)

func main() {
	dialer := &net.Dialer{LocalAddr: nil, Timeout: 2 * time.Second}
	conn, err := dialer.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatal(err)
	}

	c := client.New(conn)
	go func() {
		for {
			bytes, err := c.Read()
			if err != nil {
				log.Fatal(err)
			}
			msg := &protocol.Msg{}
			if err := json.Unmarshal(bytes, msg); err != nil {
				log.Fatal(err)
			}
			fmt.Println(*msg)

		}
	}()

	time.Sleep(300 * time.Second)
}
