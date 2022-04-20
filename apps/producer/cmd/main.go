package main

import (
	"encoding/json"
	"fmt"
	"github.com/hauson/nsq-study/apps/nsqd/protocol"
	"log"
	"net"
	"time"
)

func main() {
	dialer := &net.Dialer{LocalAddr: nil, Timeout: 2 * time.Second}
	conn, err := dialer.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatal(err)
	}

	client := protocol.NewClient(conn)
	for i := 0; i < 10; i++ {
		msg := &protocol.Msg{
			Type: "unkonwn",
			Data: "hello, world",
		}
		data, err := json.Marshal(msg)
		if err != nil {
			log.Fatal(err)
		}

		client.Write(data)
		fmt.Println("send a message")
		time.Sleep(5 * time.Second)

	}
}
