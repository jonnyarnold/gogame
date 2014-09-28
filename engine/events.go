/*
Event handling framework for game engine.
*/

package engine

import (
	"github.com/scottferg/Go-SDL2/sdl"
)

// KeyCode ...
type KeyCode int

// KeyHandler responds to keyboard events.
type KeyHandler struct {
	RespondsTo  func(key KeyCode) bool
	WhenKeyDown func(key KeyCode)
	WhenKeyUp   func(key KeyCode)
}

// Array of subscribers to send keyboard events.
var KeyHandlers []*KeyHandler

// HandleEvents checks for events and dispatches them.
func HandleEvents(state *gameState) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			state.exiting = true
		}
	}

	// As well as the events above, we also check the KeyState and dispatch events.
	// TODO: Nasty nest
	keys := sdl.GetKeyState()
	for key, value := range keys {
		for _, handler := range KeyHandlers {
			keyCode := KeyCode(key)

			if (*handler).RespondsTo(keyCode) {
				if value != 0 {
					(*handler).WhenKeyDown(keyCode)
				} else {
					(*handler).WhenKeyUp(keyCode)
				}
			}
		}
	}
}
