package main

import (
	"sync"

	"github.com/gliderlabs/ssh"
)

type Board [][]rune

type SymbolToColor map[rune]string

type Game struct {
	BoardWidth  int
	BoardHeight int
	Snakes      map[string]*Snake
	Mutex       sync.Mutex
}

type Snake struct {
	Symbol    rune
	Color     string
	Body      []Position
	Direction rune
	IsAlive   bool
}

type Position struct {
	X int
	Y int
}

func NewGame(width, height int) *Game {
	return &Game{
		BoardWidth:  width,
		BoardHeight: height,
		Snakes:      make(map[string]*Snake),
	}
}

func (g *Game) Tick(s *ssh.Session) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	for _, v := range g.Snakes {
		if !v.IsAlive {
			continue
		}
		switch v.Direction {
		case 'a':
			if v.Body[0].X > 1 {
				v.Body[0].X -= 2
			} else {
				v.Body[0].X = 0
				// And u dead asf
			}
		case 'd':
			if v.Body[0].X < g.BoardWidth-2 {
				v.Body[0].X += 2
			} else {
				v.Body[0].X = g.BoardWidth - 1
			}
		case 'w':
			if v.Body[0].Y > 0 {
				v.Body[0].Y--
			}
		case 's':
			if v.Body[0].Y < g.BoardHeight-1 {
				v.Body[0].Y++
			}
		default:
			continue
		}

		g.Render(s)
	}
}
