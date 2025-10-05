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
	MoveUp Direction = iota
	MoveDown
	MoveLeft
	MoveRight
)

// String implement the stringer interface
func (d Direction) String() string {
	switch d {
	case MoveUp:
		return "MoveUp"
	case MoveDown:
		return "MoveDown"
	case MoveLeft:
		return "MoveLeft"
	case MoveRight:
		return "MoveRight"
	default:
		return fmt.Sprintf("Direction(%d)", int(d))
	}
}
