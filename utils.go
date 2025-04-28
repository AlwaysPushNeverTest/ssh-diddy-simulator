package main

import (
	"math/rand"
)

func RandomRuneGen() rune {
	s := "abcdefghijklmnopqrstuvwxyz"
	idx := rand.Intn(len(s))
	return rune(s[idx])
}

func RandomColorGen() string {
	colors := []string{
		"\033[31m", // Red
		"\033[32m", // Green
		"\033[33m", // Yellow
		"\033[34m", // Blue
		"\033[35m", // Magenta
		"\033[36m", // Cyan
	}
	return colors[rand.Intn(len(colors))]
}
