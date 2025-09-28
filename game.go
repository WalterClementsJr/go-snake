package main

import (
	"log"
)

type Direction int

const (
	MOVE_UP Direction = iota
	MOVE_DOWN
	MOVE_LEFT
	MOVE_RIGHT
)

type Game struct {
	snake         Snake
	currentMotion Direction
	food          []Tuple2
}

func (game *Game) move(dir Direction) {
	x, y := 0, 0

	switch dir {
	case MOVE_UP:
		y = -1
	case MOVE_DOWN:
		y = 1
	case MOVE_LEFT:
		x = -1
	case MOVE_RIGHT:
		x = 1
	}
	var prev Tuple2

	snake := &game.snake

	var grow bool
	for i := 0; i < len(snake.location); i++ {
		loc := &snake.location[i]
		temp := *loc

		// head move
		if i == 0 {
			prev = *loc

			loc.a = (loc.a + x + Width/PixelSize) % (Width / PixelSize)
			loc.b = (loc.b + y + Height/PixelSize) % (Height / PixelSize)

			for _, food := range game.food {
				if loc.a == food.a && loc.b == food.b {
					// head eat food, grow snake
					grow = true
					log.Println("EAT")
				}
			}
		} else {
			// tail pieces move
			loc.a, loc.b = prev.a, prev.b

			// track the previous snake location
			prev = temp

			if grow && i == len(snake.location)-1 {
				snake.location = append(snake.location, prev)
				break
			}
		}
	}
}

// TODO place food where the snake is not
func (game *Game) placeFood() {
	game.food = []Tuple2{
		{20, 10},
		{25, 10},
		{30, 10},
	}
}

func (game *Game) init() {
	game.snake = Snake{}
	game.snake.init()
	game.placeFood()

	game.currentMotion = MOVE_RIGHT
}

type Snake struct {
	location []Tuple2
}

func (s *Snake) init() {
	s.location = make([]Tuple2, 0)

	s.location = append(s.location, Tuple2{5, 10})
	s.location = append(s.location, Tuple2{5, 11})
	s.location = append(s.location, Tuple2{5, 12})
}

type Tuple2 struct {
	a, b int
}

// func main() {
// 	s := Snake{}
// 	s.init()
//
// 	log.Printf("%+v\n\n", s.location)
// 	s.move(MOVE_LEFT)
// 	s.move(MOVE_UP)
// 	s.move(MOVE_RIGHT)
// 	s.move(MOVE_DOWN)
// }
