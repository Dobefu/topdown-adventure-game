package interfaces

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

// Game defines the interface for the game.
type Game interface {
	ebiten.Game

	UpdateUIInput()

	GetIsDebugEnabled() (isDebugEnabled bool)
	GetIsDebugActive() (isDebugActive bool)
	GetScale() (scale float64)
	GetScene() (scene Scene)
	SetScene(scene Scene)
	GetAudioContext() (audioContext *audio.Context)
}
