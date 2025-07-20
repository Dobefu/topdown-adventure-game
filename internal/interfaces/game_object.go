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
	SetID(id uint64)
	GetPosition() (position *vectors.Vector3)
	SetPosition(position vectors.Vector3)
	GetCameraPosition() (position *vectors.Vector3)
	SetIsActive(isActive bool)
	GetIsActive() (isActive bool)
	SetScene(scene Scene)
	GetScene() *Scene
	Move(velocity vectors.Vector3)
	findMaxMovement(velocity vectors.Vector3) float64
	canMoveTo(velocity vectors.Vector3) bool
}
