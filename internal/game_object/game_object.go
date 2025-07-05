package game_object

import (
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameObject struct {
	interfaces.GameObject

	scene    *interfaces.Scene
	position vectors.Vector3
	isActive bool
}

func (g *GameObject) Init() {
	// noop
}

func (g *GameObject) Draw(screen *ebiten.Image) {
	// noop
}

func (g *GameObject) DrawBelow(screen *ebiten.Image) {
	// noop
}

func (g *GameObject) Update() (err error) {
	return nil
}

func (g *GameObject) GetCollisionRect() (x1, y1, x2, y2 float64) {
	return 0, 0, 31, 31
}
