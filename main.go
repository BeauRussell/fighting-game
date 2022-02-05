package main

import (
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"github.com/BeauRussell/fighting-game/engine/asset"
	"github.com/BeauRussell/fighting-game/engine/player"
	"github.com/BeauRussell/fighting-game/engine/scenery"
	"github.com/BeauRussell/fighting-game/game/settings"
)

func main() {
	pixelgl.Run(runGame)
}

func runGame() {
	gameSettings := settings.RetrieveSettings()
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

	loadAssets := asset.NewLoad(os.DirFS("./images"))

	bgImage, err := loadAssets.Sprite("test-stage.png")
	if err != nil {
		panic(err)
	}

	background := scenery.NewBackground(bgImage, win.Bounds().Center(), gameSettings.Stage.Ground, windowWidth)

	p1Sprite, err := loadAssets.Sprite("pengu.png")
	if err != nil {
		panic(err)
	}

	players := make([]player.Player, 0)
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

	players = append(players, player.NewPlayer(p1Sprite, player1Position, player1Keybinds, player.Stats{
		Health: p1Settings.Health,
		Damage: p1Settings.Damage,
		Speed:  p1Settings.Speed,
		Jump:   10,
	}, moveSettings.Gravity))

	players = append(players, player.NewPlayer(p1Sprite, player2Position, player2Keybinds, player.Stats{
		Health: p2Settings.Health,
		Damage: p2Settings.Damage,
		Speed:  p2Settings.Speed,
		Jump:   10,
	}, moveSettings.Gravity))

	win.Clear(pixel.RGB(0, 0, 0))

	for !win.JustPressed(pixelgl.KeyEscape) {
		background.Draw(win)
		for i := range players {
			players[i].HandleMovement(win, background.Bounds)
		}

		for i := range players {
			// Collision Detection Here
			players[i].Draw(win)
		}

		win.Update()
	}
}
