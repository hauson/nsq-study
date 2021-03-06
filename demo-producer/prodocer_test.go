package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/nsqio/go-nsq"
)

func init() {
	//var demo-producer *nsq.Producer

	//1. 打印地址
	//demo-producer.String()

	//2. 链接控制
	//demo-producer.Ping()
	//demo-producer.Stop()

	//3. 发布消息
	//demo-producer.Publish()
	//demo-producer.PublishAsync()
	//demo-producer.DeferredPublish()
	//demo-producer.DeferredPublishAsync()
	//demo-producer.MultiPublish()
	//demo-producer.MultiPublishAsync()

	//4. 设置日志
	//demo-producer.SetLogger()
	//demo-producer.SetLoggerForLevel()
	//demo-producer.SetLoggerLevel()
}

func Test_PublishAsync(t *testing.T) {
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	defer producer.Stop()

	// 异步发送消息
	doneChan := make(chan *nsq.ProducerTransaction, 1)
	if err := producer.PublishAsync("pang", []byte("hi, man"), doneChan, "arg1", 2); err != nil {
		t.Fatal(err)
	}

	for {
		tx := <-doneChan
		if tx.Error != nil {
			t.Fatal(tx.Error)
		}

		bytes, err := json.MarshalIndent(tx.Args, "", "  ")
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(string(bytes))
		return
	}
}

func Test_Publish(t *testing.T) {
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	defer producer.Stop()

	// 同步发送消息
	if err := producer.Publish("pang", []byte("hi, girl")); err != nil {
		t.Fatal(err)
	}
}

func Test_String(t *testing.T) {
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	defer producer.Stop()

	// 打印地址
	fmt.Println(producer.String())
}

func Test_Config(t *testing.T) {
	// 创建配置文件
	cfg := nsq.NewConfig()
	bytes, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	// 默认的配置值
	fmt.Println(string(bytes))
}

func Test_Stop(t *testing.T) {
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	// 停止生产者
	defer producer.Stop()
}

func Test_Ping(t *testing.T) {
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	// 停止生产者
	defer producer.Stop()

	fmt.Println("start ping")
	if err = producer.Ping(); err != nil {
		t.Fatal(err)
	}
}
