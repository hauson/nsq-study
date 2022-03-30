package main

import (
	"fmt"
	svc "github.com/judwhite/go-svc"
	"log"
	"syscall"
)

type program struct {

}

func (p *program) Init(env svc.Environment) error {
	fmt.Println("hit p.Init")
	return nil
}

func (p *program) Start() error {
	fmt.Println("hit p.Start")
	return nil
}

func (p *program) Stop() error {
	fmt.Println("hit p.Stop")
	return nil
}

func main() {
	prg := &program{}
	if err := svc.Run(prg, syscall.SIGINT, syscall.SIGTERM); err != nil {
		log.Fatal(err)
	}
}
