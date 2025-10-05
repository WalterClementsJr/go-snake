package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/walterclementsjr/gosnakegame/internal/constants"
	"github.com/walterclementsjr/gosnakegame/internal/entities"
)

func draw(game *entities.Game) {
	rl.BeginDrawing()

	rl.ClearBackground(constants.BoardColor)

	for _, food := range game.Food {
		rl.DrawRectangle(int32(food.A)*constants.PixelSize, int32(food.B)*constants.PixelSize, constants.PixelSize, constants.PixelSize, rl.Red)
	}
	for _, loc := range game.Snake.Location {
		rl.DrawRectangle(int32(loc.A)*constants.PixelSize, int32(loc.B)*constants.PixelSize, constants.PixelSize, constants.PixelSize, constants.SnakeColor)
	}
	rl.EndDrawing()
}

func stateUpdate(game *entities.Game) {
	game.Move(game.CurrentMotion)

	if rl.IsKeyPressed(rl.KeyUp) {
		game.CurrentMotion = constants.MOVE_UP
	} else if rl.IsKeyPressed(rl.KeyDown) {
		game.CurrentMotion = constants.MOVE_DOWN
	} else if rl.IsKeyPressed(rl.KeyRight) {
		game.CurrentMotion = constants.MOVE_RIGHT
	} else if rl.IsKeyPressed(rl.KeyLeft) {
		game.CurrentMotion = constants.MOVE_LEFT
	}
}

func main() {
	game := entities.Game{}
	game.Init()

	rl.InitWindow(constants.Width, constants.Height, "Snake game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(10)

	for !rl.WindowShouldClose() {
		draw(&game)
		stateUpdate(&game)
	}
}
