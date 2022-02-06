package main

import (
	"embed"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"github.com/BeauRussell/fighting-game/engine/asset"
	"github.com/BeauRussell/fighting-game/engine/player"
	"github.com/BeauRussell/fighting-game/engine/scenery"
	"github.com/BeauRussell/fighting-game/game/settings"
)

//go:embed images/*
//go:embed settings.yml
var content embed.FS

func main() {
	pixelgl.Run(runGame)
}

func runGame() {
	gameSettings := settings.RetrieveSettings(content)
	moveSettings := gameSettings.Movement
	windowWidth := gameSettings.Window.Width
	widowHeight := gameSettings.Window.Height

	cfg := pixelgl.WindowConfig{
		Title:     "Fighting Game",
		Bounds:    pixel.R(0, 0, windowWidth, widowHeight),
		VSync:     true,
		Resizable: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(false)

	loadAssets := asset.NewLoad(content)

	bgImage, err := loadAssets.Sprite("images/test-stage.png")
	if err != nil {
		panic(err)
	}

	background := scenery.NewBackground(bgImage, win.Bounds().Center(), gameSettings.Stage.Ground, windowWidth)

	p1Sprite, err := loadAssets.Sprite("images/pengu.png")
	if err != nil {
		panic(err)
	}

	player1Position := pixel.Vec{X: background.Bounds.Left + 100, Y: background.Bounds.Bottom}
	player2Position := pixel.Vec{X: background.Bounds.Right - 100, Y: background.Bounds.Bottom}
	player1Keybinds := player.Keybinds{
		Up:    pixelgl.KeyUp,
		Left:  pixelgl.KeyLeft,
		Right: pixelgl.KeyRight,
	}
	player2Keybinds := player.Keybinds{
		Up:    pixelgl.KeyW,
		Left:  pixelgl.KeyA,
		Right: pixelgl.KeyD,
	}

	p1Settings := gameSettings.Player.Player1
	p2Settings := gameSettings.Player.Player2

	var player1 player.Player = player.NewPlayer(p1Sprite, player1Position, player1Keybinds, player.Stats{
		Health: p1Settings.Health,
		Damage: p1Settings.Damage,
		Speed:  p1Settings.Speed,
		Jump:   10,
	}, moveSettings.Gravity)

	var player2 player.Player = player.NewPlayer(p1Sprite, player2Position, player2Keybinds, player.Stats{
		Health: p2Settings.Health,
		Damage: p2Settings.Damage,
		Speed:  p2Settings.Speed,
		Jump:   10,
	}, moveSettings.Gravity)

	win.Clear(pixel.RGB(0, 0, 0))

	for !win.JustPressed(pixelgl.KeyEscape) {
		background.Draw(win)
		player1.HandleMovement(win, background.Bounds, player2)
		player2.HandleMovement(win, background.Bounds, player1)

		player1.Draw(win)
		player2.Draw(win)

		win.Update()
	}
}
