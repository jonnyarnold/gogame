package engine

import (
	"github.com/banthar/Go-SDL/sdl"
)

type image sdl.Surface

// Image constructs an image.
func Image(filePath string) *image {
	loadedImage := sdl.Load(filePath)
	if loadedImage == nil {
		panic("Image could not be loaded: " + filePath)
	}

	displayImage := image(*sdl.DisplayFormatAlpha(loadedImage))

	loadedImage.Free()
	return &displayImage
}

func (img *image) Surface() *sdl.Surface {
	surface := sdl.Surface(*img)
	return &surface
}

// Implement paintable
func (img *image) PaintTo(pos Pos, dest paintable) {
	posRect := pos.asRect()
	dest.Surface().Blit(&posRect, img.Surface(), nil)
}
