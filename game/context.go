package game

import (
	input "github.com/quasilyte/ebitengine-input"
	"thegame/game/image"
)

type Context struct {
	Rows int
	Cols int

	Width  uint
	Height uint

	Cursor     *Cursor
	Matrix     [][]*image.Image
	MatrixHead *matrixCoordinate

	Image *image.Image

	InputSystem input.System
	Input       *input.Handler
}

type matrixCoordinate struct {
	x int
	y int
}

func NewContext(img *image.Image, rows, cols int, width, height uint) *Context {
	return &Context{
		Image:  img,
		Rows:   rows,
		Cols:   cols,
		Width:  width,
		Height: height,
	}
}

func (c *Context) Init() error {
	c.InputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})

	c.Input = c.InputSystem.NewHandler(0, keymap)

	images := c.Image.Cut(c.Rows, c.Cols)
	image.Shuffle(images)

	c.MatrixHead = &matrixCoordinate{0, 0}
	c.setMatrix(images)

	c.setCursor()

	return nil
}

func (c *Context) MoveCursor(action input.Action) {
	if !c.canMove(action) {
		return
	}

	if c.Cursor.Selected {
		c.Swap(action)
		c.UnSelectCursor()
	}

	switch action {
	case ActionMoveDown:
		c.MatrixHead.x += 1
		c.Cursor.Move(action)
	case ActionMoveUp:
		c.MatrixHead.x -= 1
		c.Cursor.Move(action)
	case ActionMoveLeft:
		if c.MatrixHead.y == 0 {
			return
		}
		c.MatrixHead.y -= 1
		c.Cursor.Move(action)
	case ActionMoveRight:
		if c.MatrixHead.y == c.Cols-1 {
			return
		}
		c.MatrixHead.y += 1
		c.Cursor.Move(action)
	}

}

func (c *Context) canMove(action input.Action) bool {
	switch action {
	case ActionMoveDown:
		if c.MatrixHead.x == c.Rows-1 {
			return false
		}
		return true
	case ActionMoveUp:
		if c.MatrixHead.x == 0 {
			return false
		}
		return true
	case ActionMoveLeft:
		if c.MatrixHead.y == 0 {
			return false
		}
		return true
	case ActionMoveRight:
		if c.MatrixHead.y == c.Cols-1 {
			return false
		}
		return true
	default:
		return false
	}
}

func (c *Context) Swap(action input.Action) {
	x := c.MatrixHead.x
	y := c.MatrixHead.y

	switch action {
	case ActionMoveDown:
		c.Matrix[x][y], c.Matrix[x+1][y] = c.Matrix[x+1][y], c.Matrix[x][y]
	case ActionMoveUp:
		c.Matrix[x][y], c.Matrix[x-1][y] = c.Matrix[x-1][y], c.Matrix[x][y]
	case ActionMoveLeft:
		c.Matrix[x][y], c.Matrix[x][y-1] = c.Matrix[x][y-1], c.Matrix[x][y]
	case ActionMoveRight:
		c.Matrix[x][y], c.Matrix[x][y+1] = c.Matrix[x][y+1], c.Matrix[x][y]
	}
}

func (c *Context) SelectCursor() {
	c.Cursor.Select()
}
func (c *Context) UnSelectCursor() {
	c.Cursor.UnSelect()
}

func (c *Context) setMatrix(imageParts []*image.Image) {
	matrix := make([][]*image.Image, c.Rows)

	for y := 0; y < c.Rows; y++ {
		matrix[y] = make([]*image.Image, c.Cols)
	}

	for y := 0; y < c.Rows; y++ {
		for x := 0; x < c.Cols; x++ {
			matrix[y][x] = imageParts[y*c.Cols+x]
		}
	}

	c.Matrix = matrix

}

func (c *Context) setCursor() {
	size := c.Matrix[0][0].Bounds().Size()
	c.Cursor = NewCursor(float32(size.X), float32(size.Y))
}

func (c *Context) setImageParts() {
	images := c.Image.Cut(c.Rows, c.Cols)
	image.Shuffle(images)
	//c.ImageParts = images
}

func (c *Context) CheckWin() bool {
	counter := 0
	final := c.Rows * c.Cols
	for x := 0; x < len(c.Matrix); x++ {
		for y := 0; y < len(c.Matrix[x]); y++ {
			img := c.Matrix[x][y]
			if x == img.StartX && y == img.StartY {
				counter++
			}
		}
	}
	return counter == final
}
