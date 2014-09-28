/*
A game engine using SDL.
*/

package engine

import (
	"github.com/banthar/Go-SDL/mixer"
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
	mixer.OpenAudio(22050, mixer.DEFAULT_FORMAT, 2, 4096)
}

// Free performs cleanup functions for underlying dependencies.
// Deferring a call to Free is recommended after an Init call.
func Free() {
	ttf.Quit()
	sdl.Quit()
	mixer.CloseAudio()
}

type gameState struct {
	exiting bool
}

// EnterGameLoop passes control to the game engine.
func EnterGameLoop(disp *display) {
	state := gameState{}

	for !state.exiting {
		HandleEvents(&state)
		(*disp).Refresh()
	}
}
