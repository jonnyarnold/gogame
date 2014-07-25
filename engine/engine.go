/*
A game engine using SDL.
*/

package engine

import (
	"github.com/banthar/Go-SDL/sdl"
	"github.com/banthar/Go-SDL/ttf"
	"runtime"
)

// Init initisalises underlying dependencies.
// This should be called before any other calls in this package.
func Init() {
	runtime.LockOSThread() // SDL calls must be on a single thread
	sdl.Init(sdl.INIT_EVERYTHING)
	ttf.Init()
}

// Free performs cleanup functions for underlying dependencies.
// Deferring a call to Free is recommended after an Init call.
func Free() {
	ttf.Quit()
	sdl.Quit()
}
