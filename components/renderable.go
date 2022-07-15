package components

import (
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
)

type Renderable struct {
	Bg    color.Color
	Fg    color.Color
	Glyph text.Glyph
}

func NewRenderableI(bg color.Color, fg color.Color, glyph text.Glyph) Renderable {
	return Renderable{bg, fg, glyph}
}
