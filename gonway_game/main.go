package main

import (
	"github.com/SterbenVD/gonway_game/api"
	"github.com/SterbenVD/gonway_game/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	config.Init()
	height := int32(config.Settings.Height)
	width := int32(config.Settings.Width)
	title := config.Settings.Title
	FPS := int32(config.Settings.FPS)
	FG_col := api.String_To_Color(config.Settings.FG_Col)
	BG_col := api.String_To_Color(config.Settings.BG_Col)
	cell_size := int32(config.Settings.Cell_Size)

	window := api.NewWindow(height, width, FPS, title)
	board := api.NewBoard(height, width, cell_size, FG_col, BG_col)
	board.Init()
	window.Run()
	for !window.Close() {
		// Reset board
		if rl.IsKeyPressed(rl.KeySpace) {
			board.Init()
		}
		// Increase FPS
		if rl.IsKeyPressedRepeat(rl.KeyRight) {
			window.ChangeFPS(true)
		}
		// Decrease FPS
		if rl.IsKeyPressedRepeat(rl.KeyLeft) {
			window.ChangeFPS(false)
		}
		board.Run()
	}
	rl.CloseWindow()
}
