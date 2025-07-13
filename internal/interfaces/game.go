package interfaces

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type Game interface {
	ebiten.Game

	UpdateUIInput()

	GetIsDebugActive() (isDebugActive bool)
	GetScale() (scale float64)
	GetScene() (scene Scene)
	SetScene(scene Scene)
	GetAudioContext() (audioContext *audio.Context)
}
