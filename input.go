package main

import (
	"io"
	"log"

	"github.com/gliderlabs/ssh"
)

func produceInput(s ssh.Session, inputCh chan<- rune) {
	buf := make([]byte, 1)

	for {
		n, err := s.Read(buf)
		if err != nil || n == 0 {
			close(inputCh)
			return
		}
		r := rune(buf[0])
		inputCh <- r
		io.Writer.Write(s, buf)
	}
}

func consumeInput(remote string, inputCh <-chan rune) {
	for r := range inputCh {
		log.Printf("[%s] user typed: %q\n", remote, r)
	}
}
