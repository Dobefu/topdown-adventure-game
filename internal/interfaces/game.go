package interfaces

import "github.com/hajimehoshi/ebiten/v2"

type Game interface {
	ebiten.Game

	GetScale() (scale float64)
	GetScene() (scene Scene)
	SetScene(scene Scene)
}
