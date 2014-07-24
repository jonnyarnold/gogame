package engine

import (
	"github.com/banthar/Go-SDL/sdl"
)

type display struct {
	size    Size
	caption string
	screen  *sdl.Surface
}

// Display constructs a window.
func Display(size Size, caption string) *display {
	disp := display{size: size}
	disp.initSurface()
	return &disp
}

func (disp *display) Refresh() {
	disp.screen.Flip()
}

func (disp *display) initSurface() {
	disp.screen = sdl.SetVideoMode(disp.size.W, disp.size.H, 32, sdl.SWSURFACE)
	if disp.screen == nil {
		panic("sdl error")
	}
}

// Implement paintable
func (disp *display) Surface() *sdl.Surface {
	return disp.screen
}

func (disp *display) PaintTo(pos Pos, dest paintable) {
	posRect := pos.asRect()
	dest.Surface().Blit(&posRect, disp.Surface(), nil)
}
