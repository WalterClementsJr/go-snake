// Package constants holds the main data structure
package constants

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	SnakeColor = rl.Color{R: 245, G: 150, B: 60, A: 99}
	BoardColor = rl.Color{R: 20, G: 20, B: 20, A: 99}
)

const (
	PixelSize = 15
	Width     = 800
	Height    = 500
)

type Direction int

const (
	MOVE_UP Direction = iota
	MOVE_DOWN
	MOVE_LEFT
	MOVE_RIGHT
)

// String implement the stringer interface
func (d Direction) String() string {
	switch d {
	case MOVE_UP:
		return "MOVE_UP"
	case MOVE_DOWN:
		return "MOVE_DOWN"
	case MOVE_LEFT:
		return "MOVE_LEFT"
	case MOVE_RIGHT:
		return "MOVE_RIGHT"
	default:
		return fmt.Sprintf("Direction(%d)", int(d))
	}
}
