package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	input "github.com/quasilyte/ebitengine-input"
	"image/color"
)

type Cursor struct {
	x float32
	y float32

	width  float32
	height float32

	Selected bool
}

func NewCursor(width, height float32) *Cursor {
	return &Cursor{
		width:  width,
		height: height,
	}
}

func (c *Cursor) Draw(screen *ebiten.Image) {
	clr := &color.RGBA{A: 255}
	if c.Selected {
		clr = &color.RGBA{R: 255, A: 255}
	}
	vector.StrokeRect(screen, c.x, c.y, c.width, c.height, 5, clr, true)
}

func (c *Cursor) Move(action input.Action) {
	switch action {
	case ActionMoveDown:
		c.y += c.height
	case ActionMoveUp:
		c.y -= c.height
	case ActionMoveLeft:
		c.x -= c.width
	case ActionMoveRight:
		c.x += c.width
	}
}

func (c *Cursor) Select() {
	c.Selected = true
}

func (c *Cursor) UnSelect() {
	c.Selected = false
}

func (c *Cursor) IsSelected() bool {
	return c.Selected

}
