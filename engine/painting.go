package engine

import (
	"github.com/banthar/Go-SDL/sdl"
)

type paintable interface {
	Surface() *sdl.Surface
	RequiresRedraw() bool
	PaintTo(paintable)
}

// PaintStack is an ordered array
// of paintable objects.
type PaintStack []paintable

// Surface returns nil. TODO
func Surface() *sdl.Surface {
	return nil
}

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
func (ps *PaintStack) PaintTo(dest paintable) {
	for _, p := range *ps {
		p.PaintTo(dest)
	}
}

// Add an object to the stack (at the front)
func (ps *PaintStack) Add(p paintable) {
	updatedPs := append(*ps, p)
	ps = &updatedPs
}
