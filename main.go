package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/gliderlabs/ssh"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	symbolToColor = make(SymbolToColor)
	game := NewGame(50, 20)
	game.CreateBoard()
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, "\033[33m")
		user := s.RemoteAddr().String()
		game.Mutex.Lock()
		game.Snakes[user] = &Snake{
			Symbol:    RandomRuneGen(),
			Color:     RandomColorGen(),
			Body:      []Position{{X: 25, Y: 10}},
			Direction: 'd',
			IsAlive:   true,
		}
		symbolToColor[game.Snakes[user].Symbol] = game.Snakes[user].Color
		game.Mutex.Unlock()

		inputCh := make(chan rune)

		go consumeInput(game, user, inputCh, &s)
		go func() {
			for {
				time.Sleep(time.Millisecond * 100)
				game.Tick(&s)
			}
		}()
		produceInput(s, inputCh)

		// TO BE TESTED
		game.Mutex.Lock()
		delete(game.Snakes, user)
		game.Mutex.Unlock()
	})

	log.Printf("listening on :%s …\n", port)
	log.Fatal(ssh.ListenAndServe(":"+port, nil))
}
