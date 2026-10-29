package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	initScreenW = 1280
	initScreenH = 720
)

func main() {
	fmt.Println("Hello World")
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(initScreenW, initScreenH, "Testing title")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.EndDrawing()
	}
}
