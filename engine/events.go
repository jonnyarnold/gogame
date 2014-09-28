package engine

import (
	"github.com/banthar/Go-SDL/sdl"
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

// MouseButtons ...
type MouseButtons int

// MouseHandler responds to mouse events.
type MouseHandler struct {
	RespondsTo func(pos Pos, mouse MouseButtons) bool
	OnClick    func()
}

// Array of subscribers to send mouse events.
var MouseHandlers []*MouseHandler

// Stores the last mouse state
var previousMouseState = MouseButtons(0)

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
		keyCode := KeyCode(key)

		for _, handler := range KeyHandlers {
			if (*handler).RespondsTo(keyCode) {
				if value != 0 {
					(*handler).WhenKeyDown(keyCode)
				} else {
					(*handler).WhenKeyUp(keyCode)
				}
			}
		}
	}

	x, y := 0, 0
	buttons := MouseButtons(sdl.GetMouseState(&x, &y))
	pos := Pos{X: int32(x), Y: int32(y)}
	for _, handler := range MouseHandlers {
		if (*handler).RespondsTo(pos, buttons) {
			if previousMouseState == 1 && buttons == 0 {
				(*handler).OnClick()
			}
		}
	}

	previousMouseState = buttons
}
