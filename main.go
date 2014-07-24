package main

import (
	"jonnyarnold/game/engine"
)

func main() {
	engine.Init()
	defer engine.Free()

	disp := engine.Display(engine.Size{W: 640, H: 480}, "Hello, World!")

	icon := engine.Image("email.png")
	defer icon.Free()

	bg := engine.Image("bg.jpg")
	defer bg.Free()

	bg.PaintTo(engine.Pos{X: 0, Y: 0}, disp)
	icon.PaintTo(engine.Pos{X: 100, Y: 100}, disp)

	disp.Refresh()
}
