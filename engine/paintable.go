package engine

import (
	"github.com/banthar/Go-SDL/sdl"
)

type paintable interface {
	Surface() *sdl.Surface
	PaintTo(Pos, paintable)
}
