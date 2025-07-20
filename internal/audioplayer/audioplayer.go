// Package audioplayer provides a custom audio player with some extra features.
package audioplayer

import (
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/topdown-adventure-game/internal/storage"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

// AudioPlayer struct provides a wrapper for the Ebiten audio player.
type AudioPlayer struct {
	interfaces.AudioPlayer

	player *audio.Player
}

// NewAudioPlayerFromBytes creates a new audio player from audio bytes.
func NewAudioPlayerFromBytes(ctx *audio.Context, src []byte) (audioPlayer *AudioPlayer) {
	audioPlayer = &AudioPlayer{
		player: ctx.NewPlayerFromBytes(src),
	}

	return audioPlayer
}

// Play plays an audio fragment with the volume from the storage.
func (a *AudioPlayer) Play() (err error) {
	currentVolume, err := storage.GetOption("volume", 100)

	if err != nil {
		return err
	}

	a.player.SetVolume(float64(currentVolume) / 100)

	err = a.player.Rewind()

	if err != nil {
		return err
	}

	a.player.Play()

	return nil
}
