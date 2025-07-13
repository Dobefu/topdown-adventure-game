package audioplayer

import (
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/topdown-adventure-game/internal/storage"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type AudioPlayer struct {
	interfaces.AudioPlayer

	player *audio.Player
}

func NewAudioPlayerFromBytes(ctx *audio.Context, src []byte) (audioPlayer *AudioPlayer) {
	audioPlayer = &AudioPlayer{
		player: ctx.NewPlayerFromBytes(src),
	}

	return audioPlayer
}

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
