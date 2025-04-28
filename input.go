package main

import (
	"github.com/gliderlabs/ssh"
	"io"
)

func handleInput(s ssh.Session, inputChan chan<- rune) {
	buf := make([]byte, 1)
	for {
		n, err := s.Read(buf)
		if err != nil || n == 0 {
			close(inputChan)
			return
		}
		inputChan <- rune(buf[0])
		io.WriteString(s, string(buf))
	}
}
