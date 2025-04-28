package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gliderlabs/ssh"
)

var port string

func main() {
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, "Hello world\n")
	})
	fmt.Println(os.Getenv("PORT"))
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	log.Fatal(ssh.ListenAndServe(":"+port, nil))
}
