package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/travelerspb/go_rogue/components"
	"github.com/travelerspb/go_rogue/entities"
	"github.com/travelerspb/go_rogue/systems"
	"image/color"
	"log"

	"github.com/sedyh/mizu/pkg/engine"
)

type Game struct{}

func (g *Game) Setup(w engine.World) {
	w.AddComponents(components.Pos{}, components.Renderable{})
	w.AddEntities(&entities.Player{
		Pos:        components.NewPosI(100, 100),
		Renderable: components.NewRenderableI(color.White, color.White, text.Glyph{Rune: '@'}),
	})
	w.AddSystems(&systems.Render{})
}

func (g *Game) Update() error {
	return nil
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	g := engine.NewGame(&Game{})
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
