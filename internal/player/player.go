package player

import (
	"bytes"
	_ "embed"
	"image"
	"log"
	"math"

	"github.com/Dobefu/topdown-adventure-game/internal/animation"
	"github.com/Dobefu/topdown-adventure-game/internal/bullet"
	"github.com/Dobefu/topdown-adventure-game/internal/game_object"
	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
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
	NUM_FRAMES       = 16
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
	interfaces.Player
	game_object.GameObject
	game_object.HurtableGameObject

	bulletPool []*bullet.Bullet

	velocity         vectors.Vector3
	rawInputVelocity vectors.Vector3

	input      *ebitengine_input.Handler
	imgOptions *ebiten.DrawImageOptions

	frameCount     int
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

func (p *Player) Init() {
	scene := *p.GetScene()

	for range 10 {
		b := bullet.NewBullet()

		p.bulletPool = append(p.bulletPool, b)
		scene.AddGameObject(b)
	}
}

func (p *Player) GetCollisionRect() (x1, y1, x2, y2 float64) {
	return 4, 23, 27, 31
}

func (p *Player) Move(
	velocity vectors.Vector3,
) (newVelocity vectors.Vector3, hasCollided bool) {
	x1, y1, x2, y2 := p.GetCollisionRect()

	return p.MoveWithCollisionRect(velocity, x1, y1, x2, y2)
}

func (p *Player) Draw(screen *ebiten.Image) {
	pos := p.GetPosition()

	scene := (*p.GetScene())
	camera := scene.GetCamera()

	p.imgOptions.GeoM.Reset()
	p.imgOptions.GeoM.Translate(math.Round(pos.X), math.Round(pos.Y))

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
	p.imgOptions.GeoM.Translate(
		math.Round(pos.X),
		math.Round(pos.Y+FRAME_HEIGHT*.75),
	)

	camera.Draw(
		playerShadowImg,
		p.imgOptions,
		screen,
	)
}

func (p *Player) Update() (err error) {
	p.handleMovement()
	p.handleAnimations()

	if p.input.ActionIsJustPressed(input.ActionShoot) {
		p.Shoot()
	}

	return nil
}
