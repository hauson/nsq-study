package main

import (
	"fmt"
	"log"
	"net"
	"runtime"
	"strings"
	"sync"
)

func main() {
	// toodo: topic 和  channel 先不做
	// 监听连接
	// 每隔5秒，回复可以端一个信息
	tcpListener, err := net.Listen("demo-tcp", "0.0.0.0:7151")
	if err != nil {
		panic(err)
	}

	tcpserver := tcpServer{}
	err = func(listener net.Listener) error {
		log.Println("TCP: listening on %s", listener.Addr())

		var wg sync.WaitGroup

		for {
			clientConn, err := listener.Accept()
			if err != nil {
				if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
					log.Printf("WARN temporary Accept() failure - %s", err)
					runtime.Gosched()
					continue
				}
				// theres no direct way to detect this error because it is not exposed
				if !strings.Contains(err.Error(), "use of closed network connection") {
					return fmt.Errorf("listener.Accept() error - %s", err)
				}
				break
			}

			wg.Add(1)
			go func() {
				tcpserver.Handle(clientConn)
				wg.Done()
			}()
		}

		// wait to return until all handler goroutines complete
		wg.Wait()

		log.Printf("INFO TCP: closing %s", listener.Addr())

		return nil
	}(tcpListener)

	if err != nil {
		panic(err)
	}

	fmt.Println("exit")
}
