package assets

import (
	resource "github.com/quasilyte/ebitengine-resource"

	_ "image/png"
)

const (
	ImageNone resource.ImageID = iota
	ImageGopher
)

func registerImageResources(loader *resource.Loader) {
	imageResources := map[resource.ImageID]resource.ImageInfo{
		ImageGopher: {Path: "images/gopher.png"},
	}

	for id, res := range imageResources {
		loader.ImageRegistry.Set(id, res)
	}
}
