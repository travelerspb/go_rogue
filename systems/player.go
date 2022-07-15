package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/travelerspb/go_rogue/components"
)

const speed = 10

type Player struct {
}

func (g *Player) Draw(w engine.World, _ *ebiten.Image) {
	player, ok := w.View(components.Control{}, components.Pos{}).Get()
	if !ok {
		return
	}
	var pos *components.Pos
	player.Get(&pos)
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		pos.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		pos.X += speed
	} else if ebiten.IsKeyPressed(ebiten.KeyW) {
		pos.Y += speed
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		pos.Y -= speed
	}
}
