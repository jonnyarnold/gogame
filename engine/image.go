package engine

import (
	"github.com/banthar/Go-SDL/sdl"
)

type image struct {
	surface  *sdl.Surface
	Position Pos
}

// Image constructs an image.
func Image(filePath string) *image {
	loadedImage := sdl.Load(filePath)
	if loadedImage == nil {
		panic("Image could not be loaded: " + filePath)
	}

	displayImage := *sdl.DisplayFormatAlpha(loadedImage)
	img := image{surface: &displayImage}

	loadedImage.Free()
	return &img
}

func (img *image) Surface() *sdl.Surface {
	return img.surface
}

// Implement paintable
func (img *image) PaintTo(dest paintable) {
	posRect := img.Position.asRect()
	dest.Surface().Blit(&posRect, img.Surface(), nil)
}
