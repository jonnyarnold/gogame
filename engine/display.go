package engine

import (
	"github.com/scottferg/Go-SDL2/sdl"
)

type display struct {
	size    Size
	caption string
	screen  *sdl.Surface
	Objects PaintStack
}

// Display constructs a window.
func Display(size Size, caption string) *display {
	disp := display{size: size}
	disp.initSurface()
	sdl.WM_SetCaption(caption, "")
	return &disp
}

func (disp *display) Refresh() {
	// Redraw the whole screen
	dispRect := sdl.Rect{X: 0, Y: 0, W: uint16(disp.screen.W), H: uint16(disp.screen.H)}
	disp.screen.FillRect(&dispRect, 0)

	disp.Objects.PaintTo(disp)
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

func (disp *display) RequiresRedraw() bool {
	return disp.Objects.RequiresRedraw()
}
