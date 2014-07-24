package engine

import (
	"github.com/banthar/Go-SDL/sdl"
)

// EnterEventLoop passes control to the event loop.
// The event loop will continue until a quit event is sent.
func EnterEventLoop() {
	quit := false
	for !quit {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				quit = true
			}
		}
	}
}
