package main

import (
	"fmt"

	"github.com/nsqio/go-nsq"
)

func init() {
	//var q *nsq.Consumer

	//q.IsStarved()
	//q.Stats()
	//q.Stop()
	//
	//q.SetLoggerLevel()
	//q.SetLogger()
	//q.SetLoggerForLevel()
	//
	//q.AddConcurrentHandlers()
	//q.AddHandler()
	//
	//q.ChangeMaxInFlight()
	//
	//q.ConnectToNSQD()
	//q.ConnectToNSQDs()
	//
	//q.DisconnectFromNSQD()
	//q.DisconnectFromNSQLookupd()
	//
	//q.SetBehaviorDelegate()
	//q.SetLookupdHttpClient()
}

type MessageHandler struct {
	channel string
}

func (h *MessageHandler) HandleMessage(message *nsq.Message) error {
	fmt.Printf("channel:%s, Id:%s, Msg:%s \n", h.channel, message.ID, string(message.Body))
	message.Finish()
	return nil
}

func runConsuemr(topic string, channel string) {
	messageHandler := &MessageHandler{channel: channel}
	q, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	q.AddHandler(messageHandler)
	if err = q.ConnectToNSQD("127.0.0.1:4150"); err != nil {
		panic(err)
	}

	<-q.StopChan
	fmt.Println("stop")
	return
}

func main() {
	//go runConsuemr("test", "lc")
	//go runConsuemr("test", "chx")
	go runConsuemr("pang", "ch1")
	select {}
}
