package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func main() {
	msg := "msg1\nmsg2\nmsg3"
	r := bytes.NewReader([]byte(msg))
	br := bufio.NewReader(r)
	exit := false
	for !exit {
		bs, err := br.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			exit = true
		}

		fmt.Println(string(bs))
	}
}
