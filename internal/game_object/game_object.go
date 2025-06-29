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

func (g *GameObject) Draw(screen *ebiten.Image) {
	// noop
}

func (g *GameObject) DrawShadow(screen *ebiten.Image) {
	// noop
}

func (g *GameObject) Update() (err error) {
	return nil
}
