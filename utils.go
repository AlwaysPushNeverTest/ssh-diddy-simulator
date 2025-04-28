package main

import (
	"math/rand"
)

func RandomRuneGen() rune {
	s := "abcdefghijklmnopqrstuvwxyz"
	idx := rand.Intn(len(s))
	return rune(s[idx])
}
