package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	sig := <-termChan
	fmt.Println(sig.String())
}
