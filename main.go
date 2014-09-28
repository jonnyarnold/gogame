package main

import (
	"github.com/scottferg/Go-SDL2/sdl"
	"jonnyarnold/game/engine"
)

func main() {
	engine.Init()
	defer engine.Free()
	engine.FontPath = "/usr/share/fonts/truetype"

	// Setup
	disp := engine.Window(engine.Size{W: 640, H: 480}, "Hello, World!")
	bg := engine.Image("bg.jpg")
	bg.SetPosition(engine.Pos{X: 0, Y: 0})
	disp.Objects = append(disp.Objects, bg)

	font := engine.LoadFont("ubuntu-font-family/Ubuntu-R.ttf", 32)
	// TODO: defer font.Close()

	u := font.Text("Up!", engine.Color{R: 255, G: 255, B: 255})
	u.SetCenterPos(engine.Pos{X: 320, Y: 140})
	u.SetVisible(false)
	disp.Objects = append(disp.Objects, u)

	d := font.Text("Down!", engine.Color{R: 255, G: 255, B: 255})
	d.SetCenterPos(engine.Pos{X: 320, Y: 340})
	d.SetVisible(false)
	disp.Objects = append(disp.Objects, d)

	l := font.Text("Left!", engine.Color{R: 255, G: 255, B: 255})
	l.SetCenterPos(engine.Pos{X: 220, Y: 240})
	l.SetVisible(false)
	disp.Objects = append(disp.Objects, l)

	r := font.Text("Right!", engine.Color{R: 255, G: 255, B: 255})
	r.SetCenterPos(engine.Pos{X: 420, Y: 240})
	r.SetVisible(false)
	disp.Objects = append(disp.Objects, r)

	keys := engine.KeyHandler{

		RespondsTo: func(key engine.KeyCode) bool {
			return (key == sdl.K_UP ||
				key == sdl.K_DOWN ||
				key == sdl.K_LEFT ||
				key == sdl.K_RIGHT)
		},

		WhenKeyDown: func(key engine.KeyCode) {
			switch {
			case key == sdl.K_UP:
				u.SetVisible(true)
			case key == sdl.K_DOWN:
				d.SetVisible(true)
			case key == sdl.K_LEFT:
				l.SetVisible(true)
			case key == sdl.K_RIGHT:
				r.SetVisible(true)
			}
		},

		WhenKeyUp: func(key engine.KeyCode) {
			switch {
			case key == sdl.K_UP:
				u.SetVisible(false)
			case key == sdl.K_DOWN:
				d.SetVisible(false)
			case key == sdl.K_LEFT:
				l.SetVisible(false)
			case key == sdl.K_RIGHT:
				r.SetVisible(false)
			}
		},
	}

	engine.KeyHandlers = append(engine.KeyHandlers, &keys)

	engine.EnterGameLoop(disp)
}
