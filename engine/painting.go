/*
Abstract classes to deal with painting textures to a renderer.
*/

package engine

import (
	"github.com/scottferg/Go-SDL2/sdl"
)

type paintDest interface {
	GetRenderer() *sdl.Renderer
}

type paintSrc interface {
	PaintTo(renderer *sdl.Renderer)
	RequiresRedraw() bool
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
		p.PaintTo(dest.GetRenderer())
	}
}

// Add an object to the stack (at the front)
func (ps *PaintStack) Add(p paintSrc) {
	updatedPs := append(*ps, p)
	ps = &updatedPs
}
