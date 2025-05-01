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

func (g *Game) DeleteSnake(id string) {
	g.Mutex.Lock()
	delete(g.Snakes, id)
	g.Mutex.Unlock()
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

func (g *Game) HandleFoodCollision(snake *Snake, s *ssh.Session) {
	for pos, _ := range g.Food {
		if snake.Body[0].X != pos.X || snake.Body[0].Y != pos.Y {
			continue
		}

		snake.IsAlive = true
		delete(g.Food, pos)

		snake.Body = append(snake.Body, pos)
		for i := 1; i < len(snake.Body); i++ {
			snake.Body[i-1] = snake.Body[i]
		}
	}
}

func (g *Game) HandleSnakeCollision(snake *Snake, s *ssh.Session) {
	for id, v := range g.Snakes {
		if v == snake {
			for i := 1; i < len(v.Body); i++ {
				if snake.Body[0].X == v.Body[i].X && snake.Body[0].Y == v.Body[i].Y {
					g.DeleteSnake(id)
					return
				}
			}
			continue
		}
		for _, pos := range v.Body {
			if snake.Body[0].X == pos.X && snake.Body[0].Y == pos.Y {
				if len(snake.Body) > len(v.Body) {
					g.DeleteSnake(id)
				} else {
					g.DeleteSnake((*s).RemoteAddr().String())
				}
				return
			}
		}
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

		g.HandleFoodCollision(v, s)
		for i := len(v.Body) - 1; i > 0; i-- {
			v.Body[i] = v.Body[i-1]
		}
		g.HandleSnakeCollision(v, s)

		g.Render(s)
	}
}
