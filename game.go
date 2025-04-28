package main

import (
	"sync"

	"github.com/gliderlabs/ssh"
)

type Board [][]rune

type Game struct {
	BoardWidth  int
	BoardHeight int
	Snakes      map[string]*Snake
	Mutex       sync.Mutex
}

type Snake struct {
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
	for _, v := range g.Snakes {
		if !v.IsAlive {
			continue
		}
		g.Mutex.Lock()
		defer g.Mutex.Unlock()
		switch v.Direction {
		case 'a':
			v.Body[0].X--
		case 'd':
			v.Body[0].X++
		case 'w':
			v.Body[0].Y--
		case 's':
			v.Body[0].Y++
		default:
			continue
		}
		if v.Body[0].X < 0 || v.Body[0].X >= g.BoardWidth || v.Body[0].Y < 0 || v.Body[0].Y >= g.BoardHeight {
			v.IsAlive = false
		}

		g.Render(s)
	}
}
