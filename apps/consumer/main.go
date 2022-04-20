package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	exitSig := make(chan int)

	fmt.Println("hello, world")
	/*
		var wg sync.WaitGroup
		wg.Add(3)
		exitSig := make(chan interface{})
		go func() {
			defer wg.Done()

			<- exitSig
			fmt.Println("close goroutine 1")
		}()

		go func() {
			defer wg.Done()

			<- exitSig
			fmt.Println("close goroutine 2")
		}()

		go func() {
			defer wg.Done()

			<- exitSig
			fmt.Println("close goroutine 3")
		}()
	*/

	go SaveClose(exitSig)
	//go fmt.Println(IsClose(exitSig))

	go SaveClose(exitSig)

	time.Sleep(3 * time.Second)
	fmt.Println("exit success")
}

type ExitSig struct {
	wg       sync.WaitGroup
	Ch       chan int
	callback func() error
}

func New(f func() error) *ExitSig {
	return &ExitSig{
		Ch:       make(chan int),
		callback: f,
	}
}

func (s *ExitSig) IsClose() bool {
	select {
	case <-s.Ch:
		return true
	default:
		return false
	}
}

func (s *ExitSig) Close() error {
	select {
	case <-s.Ch:
		return nil
	default:
		close(s.Ch)
		s.wg.Wait()
		return s.callback()
	}
}
