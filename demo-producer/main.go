package main

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	defer producer.Stop()

	//todo: 12 个函数, 按类别分一下
	fmt.Println(producer.String())

	if false {
		if err := producer.Publish("test", []byte("hello")); err != nil {
			panic(err)
		}
	}

	if false {
		if err := producer.DeferredPublish("test", 20*time.Second, []byte("这是一条延迟消息?")); err != nil {
			panic(err)
		}
	}

	fmt.Println("sucess")
}
