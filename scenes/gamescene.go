package scenes

import (
	"github.com/hajimehoshi/ebiten"
)

type GameScene struct {
}

func (g *GameScene) Startup() {
}

func (g *GameScene) Update(screen *ebiten.Image) error {

	return nil
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
