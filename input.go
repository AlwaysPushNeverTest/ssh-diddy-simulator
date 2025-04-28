package main

import (
	"io"

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

func consumeInput(game *Game, remote string, inputCh <-chan rune, s *ssh.Session) {
	for r := range inputCh {
		switch r {
		case 'a':
			fallthrough
		case 'd':
			fallthrough
		case 'w':
			fallthrough
		case 's':
			game.Mutex.Lock()
			game.Snakes[remote].Direction = r
			game.Mutex.Unlock()
		case 'l':
			io.WriteString(*s, "\033[H\033[2J")
			game.Mutex.Lock()
			delete(game.Snakes, remote)
			game.Mutex.Unlock()
			(*s).Exit(0)
			return
		default:
			continue
		}
	}
}
