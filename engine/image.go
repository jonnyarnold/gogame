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

// Sets the position of the image such that
// the center of the image is at the given position.
func (img *image) SetCenterPos(centerPos Pos) {
	imageDimensions := Size{W: uint32(img.surface.W), H: uint32(img.surface.H)}
	img.Position = Pos{X: centerPos.X - int32(float64(imageDimensions.W)/2), Y: centerPos.Y - int32(float64(imageDimensions.H)/2)}
}
