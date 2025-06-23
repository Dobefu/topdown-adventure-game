package player

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/Dobefu/topdown-adventure-game/internal/animation"
	"github.com/Dobefu/topdown-adventure-game/internal/game_object"
	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	ebitengine_input "github.com/quasilyte/ebitengine-input"
)

var (
	//go:embed img/player.png
	playerImgBytes []byte
	playerImg      *ebiten.Image

	frameCount int
)

const (
	frameWidth  = 32
	frameHeight = 32
	numFrames   = 4
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(playerImgBytes))

	if err != nil {
		log.Fatal(err)
	}

	playerImg = ebiten.NewImageFromImage(img)
}

type Player struct {
	game_object.GameObject

	velocity vectors.Vector3

	input      *ebitengine_input.Handler
	imgOptions *ebiten.DrawImageOptions

	frameIndex     int
	animationState animation.AnimationState
}

func NewPlayer(position vectors.Vector3) (player *Player) {
	player = &Player{}

	player.imgOptions = &ebiten.DrawImageOptions{}

	player.SetIsActive(true)
	player.SetPosition(position)

	player.input = input.Input.NewHandler(0, input.Keymap)

	player.animationState = animation.AnimationStateIdleDown

	return player
}

func (p *Player) Draw(screen *ebiten.Image) {
	pos := p.GetPosition()

	p.imgOptions.GeoM.Reset()
	p.imgOptions.GeoM.Translate(pos.X, pos.Y)

	(*p.GetScene()).GetCamera().Draw(
		ebiten.NewImageFromImage(
			playerImg.SubImage(
				image.Rect(
					p.frameIndex*frameWidth,
					int(p.animationState)*frameHeight,
					p.frameIndex*frameWidth+frameWidth,
					int(p.animationState)*frameHeight+frameHeight,
				),
			),
		),
		p.imgOptions,
		screen,
	)
}

func (p *Player) Update() (err error) {
	p.handleMovement()
	p.handleAnimations()

	return nil
}
