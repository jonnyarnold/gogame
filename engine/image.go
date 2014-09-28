package engine

import (
	"github.com/banthar/Go-SDL/sdl"
)

type image struct {
	surface  *sdl.Surface
	position Pos
	redraw   bool
	visible  bool
}

// Image constructs an image.
func Image(filePath string) *image {
	loadedImage := sdl.Load(filePath)
	if loadedImage == nil {
		panic("Image could not be loaded: " + filePath)
	}

	displayImage := *sdl.DisplayFormatAlpha(loadedImage)
	img := image{surface: &displayImage, visible: true, redraw: true}

	loadedImage.Free()
	return &img
}

func (img *image) Position() Pos {
	return img.position
}

func (img *image) SetPosition(pos Pos) {
	img.position = pos
	img.redraw = true
}

func (img *image) Size() Size {
	return Size{W: uint32(img.surface.W), H: uint32(img.surface.H)}
}

func (img *image) Rect() Rect {
	return PosSizeRect(img.position, img.Size())
}

func (img *image) Visible() bool {
	return img.visible
}

func (img *image) SetVisible(visible bool) {
	img.visible = visible
	img.redraw = true
}

// Sets the position of the image such that
// the center of the image is at the given position.
func (img *image) SetCenterPos(centerPos Pos) {
	imageDimensions := Size{W: uint32(img.surface.W), H: uint32(img.surface.H)}
	img.position = Pos{X: centerPos.X - int32(float64(imageDimensions.W)/2), Y: centerPos.Y - int32(float64(imageDimensions.H)/2)}
	img.redraw = true
}

// Implement paintSrc
func (img *image) PaintTo(dest paintDest) {
	if img.visible {
		posRect := img.position.asRect()
		dest.Surface().Blit(&posRect, img.surface, nil)
	}
}

func (img *image) RequiresRedraw() bool {
	return img.redraw
}
