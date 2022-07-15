package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/travelerspb/go_rogue/components"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"log"
)

var (
	mplusNormalFont font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type Render struct{}

func (g *Render) Draw(w engine.World, screen *ebiten.Image) {
	view := w.View(components.Pos{}, components.Renderable{})
	view.Each(func(entity engine.Entity) {
		var pos *components.Pos
		var red *components.Renderable
		entity.Get(&pos, &red)

		text.Draw(screen, string(red.Glyph.Rune), mplusNormalFont, pos.X, pos.Y, red.Fg)
	})
}
