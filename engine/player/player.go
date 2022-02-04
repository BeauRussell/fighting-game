package player

import (
	"github.com/BeauRussell/fighting-game/engine/physics"
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

func (p *Player) HandleMovement(win *pixelgl.Window, ground float64) {
	if win.Pressed(p.Keybinds.Left) {
		p.Position.X += p.MovementRules.Left
	}
	if win.Pressed(p.Keybinds.Right) {
		p.Position.X += p.MovementRules.Right
	}
	if win.Pressed(p.Keybinds.Up) {
		p.Position.Y += p.MovementRules.Up
	}
	if p.Position.Y > ground {
		p.Position.Y += float64(p.MovementRules.Gravity.CalculateVelocity())
	} else {
		p.MovementRules.Gravity.ResetVelocity()
		p.Position.Y = ground
	}
}
