package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10000; i++ {
		time.Sleep(5 * time.Second)
		fmt.Println("hello, world")
	}
}
