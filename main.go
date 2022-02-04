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
	cfg := pixelgl.WindowConfig{
		Title:     "Fighting Game",
		Bounds:    pixel.R(0, 0, 1280, 720),
		VSync:     true,
		Resizable: true,
	}

	gameSettings := settings.RetrieveSettings()
	moveSettings := gameSettings.Movement

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(false)

	loadAssets := asset.NewLoad(os.DirFS("./images"))

	p1Sprite, err := loadAssets.Sprite("pengu.png")
	if err != nil {
		panic(err)
	}

	players := make([]player.Player, 0)
	player1Position := win.Bounds().Center()
	player2Position := pixel.Vec{X: 1000, Y: 800}
	players = append(players, player.NewPlayer(p1Sprite, player1Position, player.Keybinds{
		Up:    pixelgl.KeyUp,
		Left:  pixelgl.KeyLeft,
		Right: pixelgl.KeyRight,
	}, player.MovementRules{
		Left:  moveSettings.Left,
		Right: moveSettings.Right,
		Up:    moveSettings.Up,
	}, moveSettings.Gravity))

	players = append(players, player.NewPlayer(p1Sprite, player2Position, player.Keybinds{
		Up:    pixelgl.KeyW,
		Left:  pixelgl.KeyA,
		Right: pixelgl.KeyD,
	}, player.MovementRules{
		Left:  moveSettings.Left,
		Right: moveSettings.Right,
		Up:    moveSettings.Up,
	}, moveSettings.Gravity))

	bgImage, err := loadAssets.Sprite("test-stage.png")
	if err != nil {
		panic(err)
	}

	background := scenery.NewBackground(bgImage, win.Bounds().Center(), gameSettings.Stage.Ground)

	win.Clear(pixel.RGB(0, 0, 0))

	for !win.JustPressed(pixelgl.KeyEscape) {
		background.Draw(win)
		for i := range players {
			players[i].HandleMovement(win, background.Ground)
		}

		for i := range players {
			// Collision Detection Here
			players[i].Draw(win)
		}

		win.Update()
	}
}
