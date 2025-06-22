package interfaces

import (
	"github.com/Dobefu/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameObject interface {
	Update() (err error)
	Draw(screen *ebiten.Image)
	SetPosition(position vectors.Vector3)
	GetPosition() (position *vectors.Vector3)
	SetIsActive(isActive bool)
	GetIsActive() (isActive bool)
	SetScene(scene Scene)
	GetScene() *Scene
}
