package pump

import (
	"log"

	"github.com/hauson/nsq-study/apps/frame/exitsig"
	"github.com/hauson/nsq-study/apps/mq/protocol"
)

// 最主要的功能， 收到消息，处理消息

type MsgPump struct {
	addrMsgCh <-chan protocol.AddrMsg

	*exitsig.ExitSig
}

func NewMsgPump(addrMsgCh <-chan protocol.AddrMsg) *MsgPump {
	p := &MsgPump{
		addrMsgCh: addrMsgCh,
		ExitSig:   exitsig.New(nil),
	}

	p.GoFunc(p.runLoop)
	return p
}

func (p *MsgPump) runLoop(sig exitsig.Sig) {
	for {
		select {
		case msg := <-p.addrMsgCh:
			log.Println("msg pump receive msg", msg.Addr, msg.Data)
			//todo: 谁订阅了， 然后将消息做一个分发
		case <-sig:
			return
		}
	}
}
