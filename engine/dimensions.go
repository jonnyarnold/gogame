/*
Structures to help deal with dimensions.
SDL provides a generic Rect; this splits
the Rect into its Size and Pos components
*/

package engine

import (
	"github.com/banthar/Go-SDL/sdl"
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

func (pos Pos) In(rect Rect) bool {
	return (pos.X > int32(rect.X) &&
		pos.X < int32(rect.X+int16(rect.W))) &&
		(pos.Y > int32(rect.Y) &&
			pos.Y < int32(rect.Y+int16(rect.H)))
}

type Rect sdl.Rect

// PosSizeRect takes a Pos and Size object
// and returns an SDL Rect
func PosSizeRect(pos Pos, size Size) Rect {
	return Rect{X: int16(pos.X), Y: int16(pos.Y), W: uint16(size.W), H: uint16(size.H)}
}
