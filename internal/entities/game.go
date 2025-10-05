package entities

import "github.com/walterclementsjr/gosnakegame/internal/constants"

type Game struct {
	Snake         Snake
	CurrentMotion constants.Direction
	Food          []Tuple2
}

func (game *Game) Move(dir constants.Direction) {
	x, y := 0, 0

	if dir == constants.MOVE_UP && game.CurrentMotion != constants.MOVE_DOWN {
		game.CurrentMotion = constants.MOVE_UP
	} else if dir == constants.MOVE_UP && game.CurrentMotion != constants.MOVE_UP {
		game.CurrentMotion = constants.MOVE_DOWN
	} else if dir == constants.MOVE_RIGHT && game.CurrentMotion != constants.MOVE_LEFT {
		game.CurrentMotion = constants.MOVE_RIGHT
	} else if dir == constants.MOVE_LEFT && game.CurrentMotion != constants.MOVE_RIGHT {
		game.CurrentMotion = constants.MOVE_LEFT
	}
	switch dir {
	case constants.MOVE_UP:
		y = -1
	case constants.MOVE_DOWN:
		y = 1
	case constants.MOVE_LEFT:
		x = -1
	case constants.MOVE_RIGHT:
		x = 1
	}
	var prev Tuple2

	snake := &game.Snake

	var grow bool
	for i := 0; i < len(snake.Location); i++ {
		loc := &snake.Location[i]
		temp := *loc

		// head move
		if i == 0 {
			prev = *loc

			loc.A = (loc.A + x + constants.Width/constants.PixelSize) % (constants.Width / constants.PixelSize)
			loc.B = (loc.B + y + constants.Height/constants.PixelSize) % (constants.Height / constants.PixelSize)

			for _, food := range game.Food {
				if loc.A == food.A && loc.B == food.B {
					// head touches food
					grow = true
				}
			}
		} else {
			// tail pieces move
			loc.A, loc.B = prev.A, prev.B

			// track the previous snake location
			prev = temp
			// grow snake
			if grow && i == len(snake.Location)-1 {
				snake.Location = append(snake.Location, prev)
				break
			}
		}
	}
}

// TODO PlaceFood where the snake is not
func (game *Game) PlaceFood() {
	game.Food = []Tuple2{
		{20, 10},
		{25, 10},
		{30, 10},
	}
}

func (game *Game) Init() {
	game.Snake = Snake{}
	game.Snake.init()
	game.PlaceFood()

	game.CurrentMotion = constants.MOVE_RIGHT
}

type Snake struct {
	Location []Tuple2
}

func (s *Snake) init() {
	s.Location = []Tuple2{
		{5, 10},
		{5, 11},
		{5, 12},
	}
}

type Tuple2 struct {
	A, B int
}
