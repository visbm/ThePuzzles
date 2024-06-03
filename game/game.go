package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	resource "github.com/quasilyte/ebitengine-resource"
	srcImage "image"
	"thegame/assets"
	"thegame/game/image"
	"thegame/logger"
)

type Game struct {
	ctx    *Context
	log    logger.Logger
	loader *resource.Loader

	rows int
	cols int

	width  uint
	height uint
}

func NewGame(log logger.Logger, rows, cols int, width, height uint) *Game {
	return &Game{
		log:    log,
		rows:   rows,
		cols:   cols,
		width:  width,
		height: height,
	}
}

func (g *Game) Init() error {
	g.loader = createLoader()
	assets.RegisterResources(g.loader)
	img := g.newImage()
	resized := img.Resize(g.width, g.height)

	ctx := NewContext(resized, g.rows, g.cols, g.width, g.height)
	g.ctx = ctx
	err := g.ctx.Init()
	if err != nil {
		g.log.Errorf("init game error: %w", err)
		return err
	}
	return nil
}

func (g *Game) Update() error {
	g.ctx.InputSystem.Update()

	if g.ctx.CheckWin() {
		g.log.Info("win")
		return ebiten.Termination
	}

	if g.ctx.Input.ActionIsJustPressed(ActionMoveLeft) {
		g.log.Info("head left", g.ctx.MatrixHead)
		g.ctx.MoveCursor(ActionMoveLeft)
	}
	if g.ctx.Input.ActionIsJustPressed(ActionMoveRight) {
		g.log.Info("head right", g.ctx.MatrixHead)
		g.ctx.MoveCursor(ActionMoveRight)
	}
	if g.ctx.Input.ActionIsJustPressed(ActionMoveUp) {
		g.log.Info("head up", g.ctx.MatrixHead)
		g.ctx.MoveCursor(ActionMoveUp)
	}
	if g.ctx.Input.ActionIsJustPressed(ActionMoveDown) {
		g.log.Info("head down", g.ctx.MatrixHead)
		g.ctx.MoveCursor(ActionMoveDown)
	}
	if g.ctx.Input.ActionIsJustPressed(ActionConfirm) {
		if g.ctx.Cursor.Selected {
			g.ctx.UnSelectCursor()
		} else {
			g.ctx.SelectCursor()
		}
		g.log.Info("head input", g.ctx.MatrixHead)
	}

	if g.ctx.Input.ActionIsJustPressed(ActionPrintState) {
		g.log.Info("state", g.ctx)
	}

	return nil
}

func (g *Game) Configure(screen *ebiten.Image) {
	// Draw nothing
}
func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawImages(screen)

	g.ctx.Cursor.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(g.width), int(g.height)
}

func (g *Game) DrawImages(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	for y := 0; y < g.rows; y++ {
		for x := 0; x < g.cols; x++ {
			op.GeoM.Reset()
			op.GeoM.Translate(float64(x)*float64(g.width/uint(g.cols)), float64(y)*float64(g.height/uint(g.rows)))
			screen.DrawImage(g.ctx.Matrix[y][x].Image, op)
		}
	}
}

//func (g *Game) win(screen *ebiten.Image) {
//	s := "Dangan Ronpa!"
//	fontFace := ctx.GetFontFace("font.ttf")
//	var opts ebiten.DrawImageOptions
//	opts.ColorM.ScaleWithColor(color.RGBA{A: 255})
//	opts.GeoM.Translate(64, 64)
//	text.DrawWithOptions(screen, s, fontFace, &opts)
//}

func createLoader() *resource.Loader {
	sampleRate := 44100
	audioContext := audio.NewContext(sampleRate)
	loader := resource.NewLoader(audioContext)
	loader.OpenAssetFunc = assets.OpenAsset
	return loader
}

func (g *Game) newImage() *image.Image {
	img := g.loader.LoadImage(assets.ImgBMW)

	f := img.Data

	ifdf, ok := f.(*srcImage.Image)
	if !ok {
		return &image.Image{f, 0, 0}
	}
	fmt.Print(ifdf)

	return &image.Image{img.Data, 0, 0}
}
