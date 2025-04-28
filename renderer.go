package main

import (
	"io"

	"github.com/gliderlabs/ssh"
)

var board Board

func (g *Game) CreateBoard() {
	board = make(Board, g.BoardHeight)
	for i := range g.BoardHeight {
		board[i] = make([]rune, g.BoardWidth)
		for j := range g.BoardWidth {
			board[i][j] = ' '
		}
	}
}

func (g *Game) Render(s *ssh.Session) {
	io.WriteString(*s, "\033[H\033[2J")

	for i := range g.BoardHeight {
		for j := range g.BoardWidth {
			board[i][j] = ' '
		}
	}

	for _, v := range g.Snakes {
		board[v.Body[0].Y][v.Body[0].X] = v.Symbol
	}

	for range g.BoardWidth + 2 {
		io.WriteString(*s, "#")
	}

	io.WriteString(*s, "\n")

	for i := range g.BoardHeight {
		io.WriteString(*s, "#")
		for j := range g.BoardWidth {
			io.WriteString(*s, string(board[i][j]))
		}
		io.WriteString(*s, "#\n")
	}

	for _ = range g.BoardWidth + 2 {
		io.WriteString(*s, "#")
	}

	io.WriteString(*s, "\n")
}
