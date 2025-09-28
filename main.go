package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	SnakeColor = rl.Color{245, 150, 60, 99}
	BoardColor = rl.Color{20, 20, 20, 99}
)

const (
	PixelSize = 15
	Width     = 800
	Height    = 500
)

func draw(game *Game) {
	rl.BeginDrawing()

	rl.ClearBackground(BoardColor)

	for _, food := range game.food {
		rl.DrawRectangle(int32(food.a)*PixelSize, int32(food.b)*PixelSize, PixelSize, PixelSize, rl.Red)
	}
	for _, loc := range game.snake.location {
		rl.DrawRectangle(int32(loc.a)*PixelSize, int32(loc.b)*PixelSize, PixelSize, PixelSize, SnakeColor)
	}
	rl.EndDrawing()
}

func stateUpdate(game *Game) {
	game.move(game.currentMotion)

	if rl.IsKeyPressed(rl.KeyUp) {
		game.currentMotion = MOVE_UP
	} else if rl.IsKeyPressed(rl.KeyDown) {
		game.currentMotion = MOVE_DOWN
	} else if rl.IsKeyPressed(rl.KeyRight) {
		game.currentMotion = MOVE_RIGHT
	} else if rl.IsKeyPressed(rl.KeyLeft) {
		game.currentMotion = MOVE_LEFT
	}
}

func main() {
	game := Game{}
	game.init()

	rl.InitWindow(Width, Height, "Snake game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(10)

	for !rl.WindowShouldClose() {
		draw(&game)
		stateUpdate(&game)
	}
}
