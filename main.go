package main

import (
	"jonnyarnold/game/engine"
)

func main() {
	engine.Init()
	defer engine.Free()

	disp := engine.Display(engine.Size{W: 640, H: 480}, "Hello, World!")

	bg := engine.Image("bg.jpg")
	bg.Position = engine.Pos{X: 0, Y: 0}
	bg.PaintTo(disp)

	icon := engine.Sprite("email.png", engine.Size{W: 25, H: 25})
	icon.Position = engine.Pos{X: 0, Y: 0}
	icon.PaintTo(disp)

	icon.NextFrame()
	icon.Position = engine.Pos{X: 50, Y: 0}
	icon.PaintTo(disp)

	icon.NextFrame()
	icon.Position = engine.Pos{X: 0, Y: 50}
	icon.PaintTo(disp)

	icon.NextFrame()
	icon.Position = engine.Pos{X: 50, Y: 50}
	icon.PaintTo(disp)

	disp.Refresh()
	engine.EnterEventLoop()
}
