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
		game.Snakes[user] = &Snake{
			Body:      []Position{{X: 10, Y: 10}}, // Initialize the snake's body with a starting position
			Direction: 'd',                        // Set the initial direction of the snake
			IsAlive:   true,                       // Set the snake's initial state to alive
		}
		inputCh := make(chan rune)

		go consumeInput(game, user, inputCh, &s)
		produceInput(s, inputCh)
	})

	log.Printf("listening on :%s …\n", port)
	log.Fatal(ssh.ListenAndServe(":"+port, nil))
}
