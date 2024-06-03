package assets

import (
	resource "github.com/quasilyte/ebitengine-resource"
)

const (
	FontNone resource.FontID = iota
	FontNormal
	FontBig
)

func registerFontResources(loader *resource.Loader) {
	fontResources := map[resource.FontID]resource.FontInfo{
		FontNormal: {Path: "fonts/DejavuSansMono.ttf", Size: 10},
		FontBig:    {Path: "fonts/DejavuSansMono.ttf", Size: 14},
	}

	for id, res := range fontResources {
		loader.FontRegistry.Set(id, res)
		loader.LoadFont(id)
	}
}
