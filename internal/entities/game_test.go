package entities

import (
	"reflect"
	"testing"

	"github.com/walterclementsjr/gosnakegame/internal/constants"
)

func TestSnakeMove(t *testing.T) {
	originalLocation := []Tuple2{
		{5, 10},
		{5, 11},
		{5, 12},
	}

	// test table
	tests := map[string]struct {
		innitDirection    constants.Direction
		inputDirection    constants.Direction
		expectedDirection constants.Direction
		expectedLocation  []Tuple2
	}{
		"expect snake continue moving 1 tile up": {
			inputDirection:    constants.MoveDown,
			expectedDirection: constants.MoveUp,
			expectedLocation: []Tuple2{
				{5, 11}, {5, 12}, {5, 13},
			},
		},

		"expect snake move 1 tile left": {
			inputDirection:    constants.MoveLeft,
			expectedDirection: constants.MoveLeft,
			expectedLocation: []Tuple2{
				{4, 10}, {5, 11}, {5, 13},
			},
		},

		"expect snake move 1 tile right": {
			inputDirection:    constants.MoveRight,
			expectedDirection: constants.MoveRight,
			expectedLocation: []Tuple2{
				{6, 10}, {5, 12}, {5, 13},
			},
		},

		"expect snake continue 1 tile up": {
			inputDirection:    constants.MoveDown,
			expectedDirection: constants.MoveUp,
			expectedLocation: []Tuple2{
				{6, 10}, {5, 12}, {5, 13},
			},
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			game := Game{}
			game.Init()
			game.Snake.Location = originalLocation
			game.CurrentMotion = constants.MoveUp

			game.Move(testCase.inputDirection)
			snakeLocation := game.Snake.Location

			if game.CurrentMotion != testCase.expectedDirection {
				t.Errorf("Wrong move, got %v, expected %v", game.CurrentMotion, testCase.expectedDirection)
			}

			if reflect.DeepEqual(snakeLocation, testCase.expectedLocation) {
				t.Errorf("Wrong location, got %v, expected %v", snakeLocation, testCase.expectedLocation)
			}
		})
	}
}
