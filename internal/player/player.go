package player

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"log"
	"math"

	"github.com/Dobefu/topdown-adventure-game/internal/animation"
	"github.com/Dobefu/topdown-adventure-game/internal/bullet"
	"github.com/Dobefu/topdown-adventure-game/internal/game_object"
	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/topdown-adventure-game/internal/ui"
	"github.com/Dobefu/vectors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

	MAX_HEALTH = 20
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
	game_object.HurtableGameObject

	bulletPool       []*bullet.Bullet
	shootCooldown    int
	shootCooldownMax int

	velocity         vectors.Vector3
	rawInputVelocity vectors.Vector3

	input         *ebitengine_input.Handler
	aimOverlayImg *ebiten.Image
	imgOptions    *ebiten.DrawImageOptions

	frameCount     int
	frameIndex     int
	animationState animation.AnimationState
}

func NewPlayer(position vectors.Vector3) (player *Player) {
	player = &Player{}

	player.aimOverlayImg = ebiten.NewImage(MAX_CURSOR_DISTANCE*2, MAX_CURSOR_DISTANCE*2)
	player.imgOptions = &ebiten.DrawImageOptions{}

	player.SetMaxHealth(MAX_HEALTH)
	player.SetHealth(MAX_HEALTH)

	player.SetIsActive(true)
	player.SetPosition(position)

	player.input = input.Input.NewHandler(0, input.Keymap)
	player.input.GamepadDeadzone = GAMEPAD_DEADZONE

	player.shootCooldownMax = 20

	player.animationState = animation.AnimationStateIdleDown

	return player
}

func (p *Player) Init() {
	p.GameObject.Init()

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

func (p *Player) MoveWithCollision(
	velocity vectors.Vector3,
) (newVelocity vectors.Vector3, hasCollided bool) {
	x1, y1, x2, y2 := p.GetCollisionRect()

	return p.MoveWithCollisionRect(velocity, x1, y1, x2, y2)
}

func (p *Player) Draw(screen *ebiten.Image) {
	pos := p.GetPosition()

	scene := (*p.GetScene())
	camera := scene.GetCamera()

	cameraPos := *p.GetCameraPosition()
	currentPos := *p.GetPosition()
	currentPos.Add(vectors.Vector3{
		X: FRAME_WIDTH / 2,
		Y: FRAME_HEIGHT / 2,
		Z: 0,
	})

	cameraPos.Sub(currentPos)
	cameraPos.ClampMagnitude(MAX_CAMERA_OFFSET)

	p.aimOverlayImg.Clear()
	vector.StrokeCircle(
		p.aimOverlayImg,
		MAX_CURSOR_DISTANCE,
		MAX_CURSOR_DISTANCE,
		MAX_CURSOR_DISTANCE,
		1,
		color.Alpha{A: uint8(cameraPos.Magnitude() / MAX_CAMERA_OFFSET * 127)},
		false,
	)

	vector.StrokeLine(
		p.aimOverlayImg,
		MAX_CURSOR_DISTANCE,
		MAX_CURSOR_DISTANCE,
		MAX_CURSOR_DISTANCE+float32(cameraPos.X*2),
		MAX_CURSOR_DISTANCE+float32(cameraPos.Y*2),
		1,
		color.Alpha{A: uint8(cameraPos.Magnitude() / MAX_CAMERA_OFFSET * 255)},
		false,
	)

	p.imgOptions.GeoM.Reset()
	p.imgOptions.GeoM.Translate(
		math.Round(pos.X+(FRAME_WIDTH/2)-MAX_CURSOR_DISTANCE),
		math.Round(pos.Y+(FRAME_HEIGHT/2)-MAX_CURSOR_DISTANCE),
	)
	camera.Draw(p.aimOverlayImg, p.imgOptions, screen)

	p.imgOptions.GeoM.Reset()
	p.imgOptions.GeoM.Translate(math.Round(pos.X), math.Round(pos.Y))

	camera.Draw(
		playerSubImgs[int(p.animationState)*NUM_FRAMES+p.frameIndex],
		p.imgOptions,
		screen,
	)
}

func (p *Player) DrawBelow(screen *ebiten.Image) {
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

func (p *Player) DrawUI(screen *ebiten.Image) {
	ui.DrawHealthBar(
		screen,
		vectors.Vector2{X: 5, Y: 5},
		p.GetHealth(),
		p.GetMaxHealth(),
	)
}

func (p *Player) Update() (err error) {
	p.handleMovement()
	p.handleAnimations()

	if p.shootCooldown > 0 {
		p.shootCooldown -= 1
	}

	if p.input.ActionIsPressed(input.ActionShoot) && p.shootCooldown <= 0 {
		p.Shoot()
	}

	return nil
}
