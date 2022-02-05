package scenery

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Background struct {
	Sprite   *pixel.Sprite
	Position pixel.Vec
	Bounds   Bounds
}

type Keybinds struct {
	Up, Down, Left, Right pixelgl.Button
}

type Bounds struct {
	Bottom, Left, Right float64
}

func NewBackground(sprite *pixel.Sprite, position pixel.Vec, ground float64, width float64) Background {
	return Background{
		Sprite:   sprite,
		Position: position,
		Bounds: Bounds{
			Bottom: ground,
			Left:   100,
			Right:  width - 100,
		},
	}
}

func (b *Background) Draw(win *pixelgl.Window) {
	b.Sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 1.0).Moved(b.Position))
}
