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
	directions := []rune{'w', 'a', 's', 'd'}

	for r := range inputCh {
		switch r {
		case 'a':
			fallthrough
		case 'd':
			fallthrough
		case 'w':
			fallthrough
		case 's':
			if !game.Snakes[remote].IsAlive {
				continue
			}
			indexOfDir := 0
			for i, v := range directions {
				if v == r {
					indexOfDir = i
					break
				}
			}
			if directions[(indexOfDir+2)%4] == game.Snakes[remote].Direction {
				continue
			}
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
