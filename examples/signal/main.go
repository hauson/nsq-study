package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
usage:

go run main.go
## --> listen sig
## --> PID: 4604

kill -s HUP 4604
# --> Hangup: 1
*/

func main() {
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	sig := <-termChan
	fmt.Println(sig.String())
}
