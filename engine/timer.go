package engine

import (
	"github.com/banthar/Go-SDL/sdl"
)

type timer struct {
	startTicks uint32
	running    bool
	stopTicks  uint32
}

// Timer creates a timer.
func Timer() *timer {
	t := timer{}
	return &t
}

func (t *timer) Start() {
	t.startTicks = sdl.GetTicks()
	t.running = true
}

func (t *timer) Stop() {
	t.stopTicks = sdl.GetTicks()
	t.running = false
}

func (t *timer) GetTicks() uint32 {
	if t.running {
		return sdl.GetTicks() - t.startTicks
	} else {
		return t.stopTicks - t.startTicks
	}
}
