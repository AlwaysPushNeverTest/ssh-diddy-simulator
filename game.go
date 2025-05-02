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
	delete(g.Snakes, id)
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

		lastPos := snake.Body[len(snake.Body)-1]
		delete(g.Food, pos)
		snake.Body = append(snake.Body, lastPos)

		// for i := 1; i < len(snake.Body); i++ {
		// 	snake.Body[i-1] = snake.Body[i]
		// }
	}
}

func (g *Game) HandleSnakeCollision(snake *Snake, s *ssh.Session) {
	for _, s := range g.Snakes {
		if s == snake {
			continue
		}

		for _, pos := range s.Body {
			if snake.Body[0].X == pos.X && snake.Body[0].Y == pos.Y {
				if len(snake.Body) > len(s.Body) {
					s.IsAlive = false
				} else if len(snake.Body) < len(s.Body) {
					snake.IsAlive = false
				} else {
					s.IsAlive = false
					snake.IsAlive = false
				}
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
	snake, ok := g.Snakes[(*s).RemoteAddr().String()]

	if !ok {
		return
	}

	if !g.Snakes[(*s).RemoteAddr().String()].IsAlive {
		delete(g.Snakes, (*s).RemoteAddr().String())
	}

	switch snake.Direction {
	case 'a':
		if snake.Body[0].X > 0 {
			snake.Body[0].X--
		}
	case 'd':
		if snake.Body[0].X < g.BoardWidth-1 {
			snake.Body[0].X++
		}
	case 'w':
		if snake.Body[0].Y > 0 {
			snake.Body[0].Y--
		}
	case 's':
		if snake.Body[0].Y < g.BoardHeight-1 {
			snake.Body[0].Y++
		}
	default:
		return
	}
	for i := len(snake.Body) - 1; i > 0; i-- {
		snake.Body[i] = snake.Body[i-1]
	}
	g.HandleFoodCollision(snake, s)
	g.HandleSnakeCollision(snake, s)

	g.Render(s)
}
