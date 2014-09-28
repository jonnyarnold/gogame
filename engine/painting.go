package engine

import (
	"github.com/banthar/Go-SDL/sdl"
)

// A paintSrc is an object
// that can be painted onto a
// paintDest.
type paintSrc interface {
	RequiresRedraw() bool
	PaintTo(paintDest)
}

// A paintDest exposes an
// SDL surface.
type paintDest interface {
	Surface() *sdl.Surface
}

// PaintStack is an ordered array
// of paintable objects.
type PaintStack []paintSrc

// RequiresRedraw returns true if any object
// on the stack requires a redraw
func (ps *PaintStack) RequiresRedraw() bool {
	for _, p := range *ps {
		if p.RequiresRedraw() {
			return true
		}
	}

	return false
}

// PaintTo paints each object on the stack in order
func (ps *PaintStack) PaintTo(dest paintDest) {
	for _, p := range *ps {
		p.PaintTo(dest)
	}
}

// Add an object to the stack (at the front)
func (ps *PaintStack) Add(p paintSrc) {
	updatedPs := append(*ps, p)
	ps = &updatedPs
}
