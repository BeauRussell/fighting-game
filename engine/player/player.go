package player

import (
	"github.com/BeauRussell/fighting-game/engine/physics"
	"github.com/BeauRussell/fighting-game/engine/scenery"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Player struct {
	Sprite        *pixel.Sprite
	Position      pixel.Vec
	Keybinds      Keybinds
	MovementRules MovementRules
}

type Keybinds struct {
	Up, Left, Right pixelgl.Button
}

type MovementRules struct {
	Left, Right, Up float64
	Gravity         physics.Gravity
}

func NewPlayer(sprite *pixel.Sprite, position pixel.Vec, keybinds Keybinds, rules MovementRules, gravity float64) Player {
	g := physics.NewGravity(gravity)
	rules.Gravity = g
	return Player{
		Sprite:        sprite,
		Position:      position,
		Keybinds:      keybinds,
		MovementRules: rules,
	}
}

func (p *Player) Draw(win *pixelgl.Window) {
	p.Sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.0).Moved(p.Position))
}

func (p *Player) HandleMovement(win *pixelgl.Window, bounds scenery.Bounds) {
	if win.Pressed(p.Keybinds.Left) && p.Position.X > bounds.Left {
		p.Position.X += p.MovementRules.Left
	}
	if win.Pressed(p.Keybinds.Right) && p.Position.X < bounds.Right {
		p.Position.X += p.MovementRules.Right
	}
	if win.Pressed(p.Keybinds.Up) {
		p.Position.Y += p.MovementRules.Up
	}
	if p.Position.Y > bounds.Bottom {
		p.Position.Y += float64(p.MovementRules.Gravity.CalculateVelocity())
	} else {
		p.MovementRules.Gravity.ResetVelocity()
		p.Position.Y = bounds.Bottom
	}
}
