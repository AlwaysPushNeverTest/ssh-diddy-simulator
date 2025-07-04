package main

import (
	"io"

	"github.com/gliderlabs/ssh"
)

var board Board
var symbolToColor SymbolToColor

func (g *Game) CreateBoard() {
	board = make(Board, g.BoardHeight)
	for i := range g.BoardHeight {
		board[i] = make([]rune, g.BoardWidth)
		for j := range g.BoardWidth {
			board[i][j] = ' '
		}
	}
}

func (g *Game) Render(s ssh.Session) {
	io.WriteString(s, "\033[H\033[2J")

	for i := range g.BoardHeight {
		for j := range g.BoardWidth {
			board[i][j] = ' '
		}
	}

	for pos, food := range g.Food {
		board[pos.Y][pos.X] = food.Symbol
	}

	for _, v := range g.Snakes {
		for _, pos := range v.Body {
			board[pos.Y][pos.X] = v.Symbol
		}
	}

	for range g.BoardWidth + 2 {
		io.WriteString(s, "#")
	}

	io.WriteString(s, "\n")

	for i := range g.BoardHeight {
		io.WriteString(s, "#")
		for j := range g.BoardWidth {
			cell := board[i][j]
			if cell == ' ' {
				io.WriteString(s, " ")
			} else if color, ok := symbolToColor[cell]; ok {
				io.WriteString(s, color+string(cell)+"\033[0m")
			} else {
				io.WriteString(s, string(cell))
			}
		}
		io.WriteString(s, "#\n")
	}

	for _ = range g.BoardWidth + 2 {
		io.WriteString(s, "#")
	}

	io.WriteString(s, "\n")
}
