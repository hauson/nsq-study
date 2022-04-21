package pump

import (
	"sync"

	"github.com/hauson/nsq-study/apps/nsqd/protocol"
)

type Pump struct {
	addrMsgCh <-chan protocol.AddrMsg
	exitSig   chan interface{}
	wg        sync.WaitGroup
}

func NewPump(addrMsgCh <-chan protocol.AddrMsg) *Pump {
	p := &Pump{
		addrMsgCh: addrMsgCh,
		exitSig:   make(chan interface{}, 1),
	}

	p.wg.Add(1)
	go p.runLoop()
	return p
}

func (p *Pump) runLoop() {
	defer p.wg.Done()

	for {
		select {
		case msg := <-p.addrMsgCh:
			_ = msg
		case <-p.exitSig:
			return
		}
	}
}

//a blocking process
func (p *Pump) Close() {
	//todo: 可能需要加锁
	//todo: 如果已经关闭了，则直接退出
	close(p.exitSig)
	p.wg.Wait()
}
