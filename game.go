package main

import (
	"math/rand"
	"sync"
	"time"

	"github.com/gliderlabs/ssh"
)

type Board [][]rune

type SymbolToColor map[rune]string

type Game struct {
	BoardWidth  int
	BoardHeight int
	Snakes      map[string]*Snake
	Food        map[Position]*Food
	Mutex       sync.Mutex
}

type Snake struct {
	Symbol    rune
	Color     string
	Body      []Position
	Direction rune
	IsAlive   bool
}

type Food struct {
	CreatedAt time.Time
	LifeTime  time.Duration
	DickSize  int
	Symbol    rune
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
		Food:        make(map[Position]*Food),
	}
}

func (g *Game) spawnApple() {
	p := Position{
		X: rand.Intn(g.BoardWidth),
		Y: rand.Intn(g.BoardHeight),
	}

	g.Food[p] = &Food{
		CreatedAt: time.Now(),
		LifeTime:  time.Second * 10,
		DickSize:  1,
		Symbol:    'ϖ',
	}
}

func (g *Game) HandleCollision(snake *Snake, s *ssh.Session) {
	for pos, _ := range g.Food {
		if snake.Body[0].X != pos.X || snake.Body[0].Y != pos.Y {
			continue
		}

		snake.IsAlive = true
		delete(g.Food, pos)

		// newBodyPos := snake.Body[len(snake.Body)-1]
		// snake.Body = append(snake.Body, newBodyPos)
		// for i := range snake.Body {
		// 	snake.Body[i-1] = snake.Body[i]
		// }
	}
}

func (g *Game) Tick(s *ssh.Session) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	now := time.Now()
	for pos, food := range g.Food {
		if now.Sub(food.CreatedAt) > food.LifeTime {
			delete(g.Food, pos)
		}
	}

	if len(g.Food) == 0 || rand.Float64() < 0.1 {
		g.spawnApple()
	}

	for _, v := range g.Snakes {
		if !v.IsAlive {
			continue
		}
		switch v.Direction {
		case 'a':
			if v.Body[0].X > 0 {
				v.Body[0].X--
			}
		case 'd':
			if v.Body[0].X < g.BoardWidth-1 {
				v.Body[0].X++
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

		for i := len(v.Body) - 1; i > 0; i-- {
			v.Body[i] = v.Body[i-1]
		}

		g.Render(s)
		g.HandleCollision(v, s)
	}
}
