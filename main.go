package main

import (
	"jonnyarnold/game/engine"
)

func main() {
	engine.Init()
	defer engine.Free()
	engine.FontPath = "/usr/share/fonts/truetype"

	// Setup
	disp := engine.Display(engine.Size{W: 640, H: 480}, "Hello, World!")
	bg := engine.Image("bg.jpg")
	icon := engine.Sprite("email.png", engine.Size{W: 25, H: 25})

	font := engine.LoadFont("ubuntu-font-family/Ubuntu-R.ttf", 16)
	//defer font.Close()

	message := font.Text("Hello World!", engine.Color{R: 255, G: 255, B: 255})

	// Painting
	bg.Position = engine.Pos{X: 0, Y: 0}
	bg.PaintTo(disp)

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

	message.SetCenterPos(engine.Pos{X: 37, Y: 37})
	message.PaintTo(disp)

	disp.Refresh()
	engine.EnterEventLoop()
}
