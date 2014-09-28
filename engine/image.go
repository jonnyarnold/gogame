/*
Provides an Image struct that holds a static image.
*/

package engine

import (
	"github.com/scottferg/Go-SDL2/sdl"
)

type image struct {
	texture  *sdl.Texture
	size     Size
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

	img := image{visible: true, redraw: true, size: Size{W: uint32(loadedImage.W), H: uint32(loadedImage.H)}}
	img.texture = sdl.CreateTextureFromSurface(nil, loadedImage)
	return &img
}

// Implement paintable
func (img *image) PaintTo(renderer *sdl.Renderer) {
	if img.visible {
		posRect := img.position.asRect()
		renderer.Copy(img.texture, nil, &posRect)
	}
}

func (img *image) RequiresRedraw() bool {
	return img.redraw
}

// Position changes
func (img *image) Position() Pos {
	return img.position
}

func (img *image) SetPosition(pos Pos) {
	img.position = pos
	img.redraw = true
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
	img.position = Pos{X: centerPos.X - int32(float64(img.size.W)/2), Y: centerPos.Y - int32(float64(img.size.H)/2)}
	img.redraw = true
}
