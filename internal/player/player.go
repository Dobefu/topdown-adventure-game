package player

import (
	"image/color"

	"github.com/Dobefu/topdown-adventure-game/internal/game_object"
	"github.com/Dobefu/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	game_object.GameObject
}

func NewPlayer(position vectors.Vector3) (player *Player) {
	player = &Player{}

	player.SetIsActive(true)
	player.SetPosition(position)

	return player
}

func (p *Player) Draw(screen *ebiten.Image) {
	pos := p.GetPosition()

	vector.DrawFilledCircle(
		screen,
		float32(pos.X),
		float32(pos.Y),
		10,
		color.White,
		true,
	)
}
