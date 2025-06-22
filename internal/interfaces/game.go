package interfaces

import "github.com/hajimehoshi/ebiten/v2"

type Game interface {
	ebiten.Game

	GetScene() (scene Scene)
	SetScene(scene Scene)
}
