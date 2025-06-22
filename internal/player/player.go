package player

import (
	"image/color"

	"github.com/Dobefu/topdown-adventure-game/internal/game_object"
	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	ebitengine_input "github.com/quasilyte/ebitengine-input"
)

type Player struct {
	game_object.GameObject

	input *ebitengine_input.Handler
}

func NewPlayer(position vectors.Vector3) (player *Player) {
	player = &Player{}

	player.SetIsActive(true)
	player.SetPosition(position)

	player.input = input.Input.NewHandler(0, input.Keymap)

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

func (p *Player) Update() (err error) {
	pos := p.GetPosition()

	if p.input.ActionIsPressed(input.ActionMoveLeft) {
		pos.X -= 1
	}

	if p.input.ActionIsPressed(input.ActionMoveRight) {
		pos.X += 1
	}

	if p.input.ActionIsPressed(input.ActionMoveUp) {
		pos.Y -= 1
	}

	if p.input.ActionIsPressed(input.ActionMoveDown) {
		pos.Y += 1
	}

	return nil
}
