package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var flag int32
	go func() {
		atomic.StoreInt32(&flag, 1)
		fmt.Println("1", flag)
		time.Sleep(20*time.Second)
		atomic.StoreInt32(&flag, 0)
		wg.Done()
		fmt.Println("1", flag)
	}()
	go func() {
		atomic.StoreInt32(&flag, 1)
		fmt.Println("2",flag)
		atomic.StoreInt32(&flag, 0)
		wg.Done()
		fmt.Println("2",flag)
	}()

	wg.Wait()
	fmt.Println(flag)
}
