package interfaces

import (
	"github.com/Dobefu/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

// GameObject defines the interface for a base game object.
type GameObject interface {
	Init()
	Update() (err error)
	Draw(screen *ebiten.Image)
	DrawBelow(screen *ebiten.Image)
	DrawAbove(screen *ebiten.Image)
	DrawUI(screen *ebiten.Image)
	GetID() (id uint64)
	GetPosition() (position *vectors.Vector3)
	GetCameraPosition() (position *vectors.Vector3)
	GetIsActive() (isActive bool)
	SetScene(scene Scene)
	Damage(amount int, source GameObject)
}
