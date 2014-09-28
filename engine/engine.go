/*
A game engine using SDL.
*/

package engine

import (
	"github.com/scottferg/Go-SDL2/sdl"
	"github.com/scottferg/Go-SDL2/ttf"
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

type gameState struct {
	exiting bool
}

// EnterGameLoop passes control to the game engine.
func EnterGameLoop(win *window) {
	state := gameState{}

	for !state.exiting {
		HandleEvents(&state)
		(*win).Refresh()
	}
}
