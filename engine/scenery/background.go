package scenery

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Background struct {
	Sprite   *pixel.Sprite
	Position pixel.Vec
	Ground   float64
}

type Keybinds struct {
	Up, Down, Left, Right pixelgl.Button
}

func NewBackground(sprite *pixel.Sprite, position pixel.Vec, ground float64) Background {
	return Background{
		Sprite:   sprite,
		Position: position,
		Ground:   ground,
	}
}

func (b *Background) Draw(win *pixelgl.Window) {
	b.Sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 1.0).Moved(b.Position))
}
