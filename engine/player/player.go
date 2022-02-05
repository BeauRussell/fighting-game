package player

import (
	"fmt"

	"github.com/BeauRussell/fighting-game/engine/physics"
	"github.com/BeauRussell/fighting-game/engine/scenery"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type Player struct {
	Sprite   *pixel.Sprite
	Position pixel.Vec
	Keybinds Keybinds
	Stats    Stats
}

type Keybinds struct {
	Up, Left, Right pixelgl.Button
}

type Stats struct {
	Health  float64
	Damage  float64
	Speed   float64
	Jump    float64
	Gravity physics.Gravity
}

func NewPlayer(sprite *pixel.Sprite, position pixel.Vec, keybinds Keybinds, stats Stats, gravity float64) Player {
	g := physics.NewGravity(gravity)
	stats.Gravity = g
	return Player{
		Sprite:   sprite,
		Position: position,
		Keybinds: keybinds,
		Stats:    stats,
	}
}

func (p *Player) Draw(win *pixelgl.Window) {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	statsBox := text.New(pixel.V(p.Position.X+50, p.Position.Y+50), atlas)
	p.Sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.0).Moved(p.Position))
	fmt.Fprintf(statsBox, "Health: %s\nSpeed: %s\nDamage: %s", fmt.Sprint(p.Stats.Health), fmt.Sprint(p.Stats.Speed), fmt.Sprint(p.Stats.Damage))
	statsBox.Draw(win, pixel.IM)
}

func (p *Player) HandleMovement(win *pixelgl.Window, bounds scenery.Bounds) {
	if win.Pressed(p.Keybinds.Left) && p.Position.X > bounds.Left {
		p.Position.X -= p.Stats.Speed
	}
	if win.Pressed(p.Keybinds.Right) && p.Position.X < bounds.Right {
		p.Position.X += p.Stats.Speed
	}
	if win.Pressed(p.Keybinds.Up) {
		p.Position.Y += p.Stats.Jump
	}
	if p.Position.Y > bounds.Bottom {
		p.Position.Y += float64(p.Stats.Gravity.CalculateVelocity())
	} else {
		p.Stats.Gravity.ResetVelocity()
		p.Position.Y = bounds.Bottom
	}
}
