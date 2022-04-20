package protocol

import "fmt"

type Msg struct {
	Type string
	Data interface{}
}

type MsgHandler func(*Msg) error

var msgHandlers map[interface{}]MsgHandler

func init() {
	msgHandlers = make(map[interface{}]MsgHandler)
	msgHandlers["default"] = func(msg *Msg) error {
		fmt.Println(msg.Data)
		return nil
	}
}

type AddrMsg struct {
	*Msg
	Addr string
}
