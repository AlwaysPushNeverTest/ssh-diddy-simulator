package main

import (
	"io"
	"log"
	"os"

	"github.com/gliderlabs/ssh"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ssh.Handle(func(s ssh.Session) {

		io.WriteString(s, "Type away (Ctrl-D to exit):\n")

		inputCh := make(chan rune)

		go consumeInput("user", inputCh)
		produceInput(s, inputCh)
	})

	log.Printf("listening on :%s …\n", port)
	log.Fatal(ssh.ListenAndServe(":"+port, nil))
}
