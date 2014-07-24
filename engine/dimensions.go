package engine

import (
	"github.com/banthar/Go-SDL/sdl"
)

// Size denotes the dimensions (not position) of a rectangle
type Size struct {
	W int
	H int
}

// Pos denotes the position (not dimensions) of an object
type Pos struct {
	X int16
	Y int16
}

func (pos Pos) asRect() sdl.Rect {
	return sdl.Rect{X: pos.X, Y: pos.Y}
}
