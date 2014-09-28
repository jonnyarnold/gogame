/*
Structures to help deal with dimensions.
SDL provides a generic Rect; this splits
the Rect into its Size and Pos components.
Utility methods are provided to ensure the
Rect values are the right type.
*/

package engine

import (
	"github.com/scottferg/Go-SDL2/sdl"
)

// Size denotes the dimensions (not position) of a rectangle
type Size struct {
	W uint32
	H uint32
}

func (size Size) asRect() sdl.Rect {
	return sdl.Rect{W: uint16(size.W), H: uint16(size.H)}
}

// Pos denotes the position (not dimensions) of an object
type Pos struct {
	X int32
	Y int32
}

func (pos Pos) asRect() sdl.Rect {
	return sdl.Rect{X: int16(pos.X), Y: int16(pos.Y)}
}

// PosSizeRect takes a Pos and Size object
// and returns an SDL Rect
func PosSizeRect(pos Pos, size Size) sdl.Rect {
	return sdl.Rect{X: int16(pos.X), Y: int16(pos.Y), W: uint16(size.W), H: uint16(size.H)}
}
