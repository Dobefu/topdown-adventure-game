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
	playerSubImgs  []*ebiten.Image

	//go:embed img/shadow.png
	playerShadowImgBytes []byte
	playerShadowImg      *ebiten.Image
)

const (
	FRAME_WIDTH      = 32
	FRAME_HEIGHT     = 32
	GAMEPAD_DEADZONE = .1
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(playerImgBytes))

	if err != nil {
		log.Fatal(err)
	}

	playerImg = ebiten.NewImageFromImage(img)
	playerSubImgs = make(
		[]*ebiten.Image,
		playerImg.Bounds().Dy()/FRAME_HEIGHT*NUM_FRAMES,
	)

	for state := range playerImg.Bounds().Dy() / FRAME_HEIGHT {
		for frame := range NUM_FRAMES {
			key := state*NUM_FRAMES + frame

			playerSubImgs[key] = ebiten.NewImageFromImage(
				playerImg.SubImage(
					image.Rect(
						frame*FRAME_WIDTH,
						state*FRAME_HEIGHT,
						frame*FRAME_WIDTH+FRAME_WIDTH,
						state*FRAME_HEIGHT+FRAME_HEIGHT,
					),
				),
			)
		}
	}

	img, _, err = image.Decode(bytes.NewReader(playerShadowImgBytes))

	if err != nil {
		log.Fatal(err)
	}

	playerShadowImg = ebiten.NewImageFromImage(img)
}

type Player struct {
	game_object.GameObject

	velocity         vectors.Vector3
	rawInputVelocity vectors.Vector3

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
	player.input.GamepadDeadzone = GAMEPAD_DEADZONE

	player.animationState = animation.AnimationStateIdleDown

	return player
}

func (p *Player) Draw(screen *ebiten.Image) {
	pos := p.GetPosition()

	scene := (*p.GetScene())
	camera := scene.GetCamera()

	p.imgOptions.GeoM.Reset()
	p.imgOptions.GeoM.Translate(pos.X, pos.Y)

	camera.Draw(
		playerSubImgs[int(p.animationState)*NUM_FRAMES+p.frameIndex],
		p.imgOptions,
		screen,
	)
}

func (p *Player) DrawShadow(screen *ebiten.Image) {
	pos := p.GetPosition()

	scene := (*p.GetScene())
	camera := scene.GetCamera()

	p.imgOptions.GeoM.Reset()
	p.imgOptions.GeoM.Translate(pos.X, pos.Y+FRAME_HEIGHT*.75)

	camera.Draw(
		playerShadowImg,
		p.imgOptions,
		screen,
	)
}

func (p *Player) Update() (err error) {
	p.handleMovement()
	p.handleAnimations()

	return nil
}
