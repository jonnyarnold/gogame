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
	disp.screen = sdl.SetVideoMode(int(disp.size.W), int(disp.size.H), 32, sdl.SWSURFACE)
	if disp.screen == nil {
		panic("sdl error")
	}
}

// Implement paintable
func (disp *display) Surface() *sdl.Surface {
	return disp.screen
}

func (disp *display) PaintTo(dest paintable) {
	dest.Surface().Blit(nil, disp.Surface(), nil)
}
