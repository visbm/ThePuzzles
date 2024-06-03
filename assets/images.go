package assets

import (
	resource "github.com/quasilyte/ebitengine-resource"
	_ "image/png"
)

const (
	ImgNone resource.ImageID = iota
	ImgBMW
)

func registerImageResources(loader *resource.Loader) {
	imageResources := map[resource.ImageID]resource.ImageInfo{
		ImgBMW: {Path: "images/bmw.png"},
	}
	for id, res := range imageResources {
		loader.ImageRegistry.Set(id, res)
		loader.LoadImage(id)
	}
}
