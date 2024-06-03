package image

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nfnt/resize"
	"image"
	_ "image/png"
	"math/rand"
	"os"
	"thegame/logger"
)

var log = logger.GetLogger()

type Image struct {
	*ebiten.Image
	StartX int
	StartY int
}

func LoadImage(path string, screenWidth, screenHeight uint) (*Image, error) {

	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Errorf("close file error: %v", err)
		}
	}()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Errorf("decode image error: %v", err)
		return nil, err
	}

	resizedImage := resize.Resize(screenWidth, screenHeight, img, resize.Lanczos3)

	//f:= ebiten.NewImageFromImage(resizedImage)

	//resizedImage2 := resize.Resize(screenWidth, screenHeight, f, resize.Lanczos3)

	return &Image{ebiten.NewImageFromImage(resizedImage), 0, 0}, nil
}

func (i *Image) Resize(screenWidth, screenHeight uint) *Image {
	resizedImage := resize.Resize(screenWidth, screenHeight, i, resize.Lanczos3)

	return &Image{ebiten.NewImageFromImage(resizedImage), 0, 0}
}

// Cut returns an array of images from the original image with the given rows and columns
func (i *Image) Cut(rows, cols int) []*Image {

	images := make([]*Image, 0, rows*cols)
	size := i.Bounds().Size()
	partWidth := float64(size.X) / float64(cols)
	partHeight := float64(size.Y) / float64(rows)

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			rect := image.Rect(y*int(partWidth), x*int(partHeight), (y+1)*int(partWidth), (x+1)*int(partHeight))
			subImg := i.SubImage(rect).(*ebiten.Image)
			images = append(images, &Image{subImg, x, y})
		}
	}
	return images

}

func (i *Image) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	screen.DrawImage(i.Image, op)
}

func Shuffle(images []*Image) {
	for i := range images {
		j := rand.Intn(len(images))
		images[i], images[j] = images[j], images[i]
	}
}

func (i *Image) Bounds() image.Rectangle {
	return i.Image.Bounds()
}
