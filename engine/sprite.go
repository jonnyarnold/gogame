package engine

import (
	"github.com/banthar/Go-SDL/sdl"
)

type sprite struct {
	*image
	sheet            *sdl.Surface
	FrameSize        Size
	currentFrameRect sdl.Rect
	position         Pos
	redraw           bool
	visible          bool
}

// Sprite constructs a sprite.
// The sprite sheet will be scanned from top-left, reading left-to-right.
func Sprite(sheetPath string, frameSize Size) *sprite {
	loadedSheet := sdl.Load(sheetPath)
	if loadedSheet == nil {
		panic("Image could not be loaded: " + sheetPath)
	}

	sheet := sdl.DisplayFormatAlpha(loadedSheet)
	s := sprite{sheet: sheet, FrameSize: frameSize, currentFrameRect: frameSize.asRect(), redraw: true}

	loadedSheet.Free()
	return &s
}

// Advance the sprite to the next frame.
func (s *sprite) NextFrame() {
	// TODO: Far too much casting here!!
	sheetSize := Size{W: uint32(s.sheet.W), H: uint32(s.sheet.H)}

	// Initially move to the right.
	framePos := Pos{X: int32(s.currentFrameRect.X) + int32(s.FrameSize.W), Y: int32(s.currentFrameRect.Y)}

	// Move to the next row if we reach the end of the row
	if framePos.X >= int32(sheetSize.W) {
		framePos = Pos{X: 0, Y: int32(s.currentFrameRect.Y) + int32(s.FrameSize.H)}

		// Move to the top left if we reach the end of the sheet
		if framePos.Y >= int32(sheetSize.H) {
			framePos = Pos{X: 0, Y: 0}
		}
	}

	s.currentFrameRect = sdl.Rect{X: int16(framePos.X), Y: int16(framePos.Y), W: uint16(s.FrameSize.W), H: uint16(s.FrameSize.H)}
	s.redraw = true
}

func (img *sprite) Position() Pos {
	return img.position
}

func (img *sprite) SetPosition(pos Pos) {
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

// Implement paintSrc
func (s *sprite) PaintTo(dest paintDest) {
	if img.visible {
		posRect := s.position.asRect()
		dest.Surface().Blit(&posRect, s.sheet, &s.currentFrameRect)
	}
}

func (img *sprite) RequiresRedraw() bool {
	return img.redraw
}
