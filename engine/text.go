package engine

import (
	"github.com/banthar/Go-SDL/sdl"
	"github.com/banthar/Go-SDL/ttf"
	"path"
)

// FontPath is the folder that
// fonts will be pulled from.
var FontPath string

type font ttf.Font

// Color is an analogue to SDL's Color
type Color sdl.Color

// LoadFont loads the font with the given filename
// at the current font path.
func LoadFont(fontName string, fontSize int) *font {
	fontPath := path.Join(FontPath, fontName)
	ttfFont := ttf.OpenFont(fontPath, fontSize)
	localFont := font(*ttfFont)
	return &localFont
}

// Creates text for the given font
func (f *font) Text(text string, color Color) *image {
	unpackedFont := ttf.Font(*f)
	textSurface := ttf.RenderUTF8_Solid(&unpackedFont, text, sdl.Color(color))
	img := image{surface: textSurface}
	return &img
}
