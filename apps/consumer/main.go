package main

import (
	"fmt"
	"time"

	"github.com/hauson/nsq-study/apps/frame"
)

func main() {
	exitSig := frame.New(func() error {
		fmt.Println("hello,world")
		return nil
	})

	fmt.Println("start")

	exitSig.RunGo(func(exitSig <-chan int) {
		for {
			select {
			case <-exitSig:
				fmt.Println("close goroutine 1")
				return
			default:
				time.Sleep(3 * time.Second)
				fmt.Println("work 1")
			}
		}
	})

	exitSig.RunGo(func(exitSig <-chan int) {
		<-exitSig
		fmt.Println("close goroutine 2")
	})

	exitSig.RunGo(func(exitSig <-chan int) {
		<-exitSig
		fmt.Println("close goroutine 3")
	})

	time.Sleep(3 * time.Second)
	exitSig.Close()
}
