package main

import (
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

	game := NewGame(50, 20)
	game.CreateBoard()
	ssh.Handle(func(s ssh.Session) {
		user := s.RemoteAddr().String()
		game.Mutex.Lock()
		game.Snakes[user] = &Snake{
			Symbol:    RandomRuneGen(),
			Body:      []Position{{X: 25, Y: 10}},
			Direction: 'd',
			IsAlive:   true,
		}
		game.Mutex.Unlock()

		inputCh := make(chan rune)

		go consumeInput(game, user, inputCh, &s)
		go func() {
			ticker := time.NewTicker(100 * time.Millisecond)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					game.Tick(&s)
				case <-s.Context().Done():
					return
				}
			}
		}()
		produceInput(s, inputCh)

		game.Mutex.Lock()
		delete(game.Snakes, user)
		game.Mutex.Unlock()
	})

	log.Printf("listening on :%s …\n", port)
	log.Fatal(ssh.ListenAndServe(":"+port, nil))
}
