package main

import (
	"jonnyarnold/game/engine"
)

func main() {
	engine.Init()
	defer engine.Free()

	disp := engine.Display(engine.Size{W: 640, H: 480}, "Hello, World!")

	bg := engine.Image("bg.jpg")
	bg.PaintTo(engine.Pos{X: 0, Y: 0}, disp)

	icon := engine.Image("email.png")
	icon.PaintTo(engine.Pos{X: 100, Y: 100}, disp)

	disp.Refresh()
	engine.EnterEventLoop()
}
