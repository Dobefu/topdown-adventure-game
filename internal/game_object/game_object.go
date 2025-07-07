package game_object

import (
	"sync/atomic"

	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	nextGameObjectID uint64
)

type GameObject struct {
	interfaces.GameObject

	id uint64

	scene    *interfaces.Scene
	position vectors.Vector3
	isActive bool
}

func (g *GameObject) Init() {
	if g.id == 0 {
		g.id = atomic.AddUint64(&nextGameObjectID, 1)
	}
}

func (g *GameObject) GetID() (id uint64) {
	return g.id
}

func (g *GameObject) SetID(id uint64) {
	g.id = id
}

func (g *GameObject) Draw(screen *ebiten.Image) {
	// noop
}

func (g *GameObject) DrawBelow(screen *ebiten.Image) {
	// noop
}

func (g *GameObject) DrawUI(screen *ebiten.Image) {
	// noop
}

func (g *GameObject) Update() (err error) {
	return nil
}
