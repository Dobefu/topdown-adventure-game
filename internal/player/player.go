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

	input      *ebitengine_input.Handler
	img        *ebiten.Image
	imgOptions *ebiten.DrawImageOptions
}

func NewPlayer(position vectors.Vector3) (player *Player) {
	player = &Player{}

	player.img = ebiten.NewImage(32, 32)
	player.imgOptions = &ebiten.DrawImageOptions{}

	player.SetIsActive(true)
	player.SetPosition(position)

	player.input = input.Input.NewHandler(0, input.Keymap)

	return player
}

func (p *Player) Draw(screen *ebiten.Image) {
	pos := p.GetPosition()

	vector.DrawFilledCircle(p.img, 16, 16, 16, color.White, true)

	p.imgOptions.GeoM.Reset()
	p.imgOptions.GeoM.Translate(pos.X, pos.Y)

	(*p.GetScene()).GetCamera().Draw(p.img, p.imgOptions, screen)

	p.img.Clear()
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
