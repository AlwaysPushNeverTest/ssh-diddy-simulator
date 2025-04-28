package main

import (
	"log"
	"os"

	"github.com/gliderlabs/ssh"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	game := NewGame(20, 20)
	game.CreateBoard()
	ssh.Handle(func(s ssh.Session) {
		user := s.RemoteAddr().String()
		game.Mutex.Lock()
		game.Snakes[user] = &Snake{
			Symbol:    RandomRuneGen(),
			Body:      []Position{{X: 10, Y: 10}},
			Direction: 'd',
			IsAlive:   true,
		}
		game.Mutex.Unlock()

		inputCh := make(chan rune)

		go consumeInput(game, user, inputCh, &s)
		produceInput(s, inputCh)
	})

	log.Printf("listening on :%s …\n", port)
	log.Fatal(ssh.ListenAndServe(":"+port, nil))
}
