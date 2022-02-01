package main

import (
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"gitlab.com/yourknightmares/fighting-game/engine/asset"
	"gitlab.com/yourknightmares/fighting-game/engine/player"
	"gitlab.com/yourknightmares/fighting-game/engine/scenery"
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

	gravityConst := float64(-0.5)

	players := make([]player.Player, 0)
	player1Position := win.Bounds().Center()
	players = append(players, player.NewPlayer(p1Sprite, player1Position, player.Keybinds{
		Up:    pixelgl.KeyUp,
		Left:  pixelgl.KeyLeft,
		Right: pixelgl.KeyRight,
	}, player.MovementRules{
		Left:  -2.0,
		Right: +2.0,
		Up:    +10.0,
	}, gravityConst))

	bgImage, err := loadAssets.Sprite("test-stage.png")
	if err != nil {
		panic(err)
	}

	background := scenery.NewBackground(bgImage, win.Bounds().Center(), 180)

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
