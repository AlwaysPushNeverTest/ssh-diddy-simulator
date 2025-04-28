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

	for _, v := range g.Snakes {
		board[v.Body[0].Y][v.Body[0].X] = 'N'
	}

	for i := range g.BoardHeight {
		for j := range g.BoardWidth {
			io.WriteString(*s, string(board[i][j]))
		}
		io.WriteString(*s, "\n")
	}

	// for i := range(g.BoardWidth + 2) {
	// 	fmt.Printf("#")
	// }
	// for i := range(g.BoardHeight) {
	// 	fmt.Printf("#")

	// 	for j := range(g.BoardWidth) {

	// 	}

	// 	fmt.Printf("#")
	// }
	// for i := range(g.BoardWidth + 2) {
	// 	fmt.Printf("#")
	// }
}
