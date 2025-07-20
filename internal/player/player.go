// Package player provides a player.
package player

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"log"

	"github.com/Dobefu/topdown-adventure-game/internal/animation"
	"github.com/Dobefu/topdown-adventure-game/internal/audioplayer"
	"github.com/Dobefu/topdown-adventure-game/internal/bullet"
	"github.com/Dobefu/topdown-adventure-game/internal/gameobject"
	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/topdown-adventure-game/internal/state"
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
	// FrameWidth defines the width of an animation frame.
	FrameWidth = 32
	// FrameHeight defines the height of an animation frame.
	FrameHeight = 32
	// NumFrames defines number of animation frames in the spritesheet.
	NumFrames = 16
	// GamepadDeadzone defines the dead zone for gamepad control sticks.
	GamepadDeadzone = .1
	// Knockback defines the amount of knockback the player gets when hit.
	Knockback = 10

	// MaxHealth defines the base max health that the player has.
	MaxHealth = 20
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(playerImgBytes))

	if err != nil {
		log.Fatal(err)
	}

	playerImg = ebiten.NewImageFromImage(img)
	playerSubImgs = make(
		[]*ebiten.Image,
		playerImg.Bounds().Dy()/FrameHeight*NumFrames,
	)

	for state := range playerImg.Bounds().Dy() / FrameHeight {
		for frame := range NumFrames {
			key := state*NumFrames + frame

			playerSubImgs[key] = ebiten.NewImageFromImage(
				playerImg.SubImage(
					image.Rect(
						frame*FrameWidth,
						state*FrameHeight,
						frame*FrameWidth+FrameWidth,
						state*FrameHeight+FrameHeight,
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

// Player struct defines the player.
type Player struct {
	interfaces.Player
	gameobject.HurtableGameObject

	audioPlayer *audioplayer.AudioPlayer

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
	animationState animation.State
	state          state.State
}

// NewPlayer creates a new player.
func NewPlayer(position vectors.Vector3) (player *Player) {
	player = &Player{}

	player.aimOverlayImg = ebiten.NewImage(MaxCursorDistance*2, MaxCursorDistance*2)
	player.imgOptions = &ebiten.DrawImageOptions{}

	player.SetMaxHealth(MaxHealth)
	player.SetHealth(MaxHealth)

	player.SetIsActive(true)
	player.SetPosition(position)

	player.input = input.Input.NewHandler(0, input.PlayerKeymap)
	player.input.GamepadDeadzone = GamepadDeadzone

	player.shootCooldownMax = 20

	player.animationState = animation.StateIdleDown

	player.SetOnCollision(func(_, other interfaces.GameObject) {
		// Skip the collision callback if the player hits a bullet they own.
		if bullet, ok := other.(*bullet.Bullet); ok {
			if bullet.GetOwner().GetID() == player.GetID() {
				return
			}
		}
	})

	player.SetDeathCallback(player.Die)

	return player
}

// Init initializes the player.
func (p *Player) Init() {
	p.GameObject.Init()
	p.CollidableGameObject.Init()

	scene := *p.GetScene()

	for range 10 {
		b := bullet.NewBullet()

		p.bulletPool = append(p.bulletPool, b)
		scene.AddGameObject(b)
	}

	audioContext := scene.GetGame().GetAudioContext()
	p.audioPlayer = audioplayer.NewAudioPlayerFromBytes(audioContext, playerShootSound)
}

// GetCollisionRect gets the four points of the collision rectangle.
func (p *Player) GetCollisionRect() (x1, y1, x2, y2 float64) {
	return 4, 19, 27, 31
}

// MoveWithCollision moves the player with collision checks.
func (p *Player) MoveWithCollision(
	velocity vectors.Vector3,
) (newVelocity vectors.Vector3, hasCollided bool) {
	x1, y1, x2, y2 := p.GetCollisionRect()

	return p.MoveWithCollisionRect(velocity, x1, y1, x2, y2)
}

// Draw runs during the game's Draw function.
func (p *Player) Draw(screen *ebiten.Image) {
	pos := p.GetPosition()

	scene := (*p.GetScene())
	camera := scene.GetCamera()

	cameraPos := *p.GetCameraPosition()
	currentPos := *p.GetPosition()
	currentPos.Z = 0

	currentPos.Add(vectors.Vector3{
		X: FrameWidth / 2,
		Y: FrameHeight / 2,
		Z: 0,
	})

	cameraPos.Sub(currentPos)
	cameraPos.ClampMagnitude(MaxCameraOffset)

	p.aimOverlayImg.Clear()

	vector.StrokeLine(
		p.aimOverlayImg,
		MaxCursorDistance,
		MaxCursorDistance,
		MaxCursorDistance+float32(cameraPos.X*2),
		MaxCursorDistance+float32(cameraPos.Y*2),
		1,
		color.Alpha{A: uint8(cameraPos.Magnitude() / MaxCameraOffset * 255)},
		false,
	)

	p.imgOptions.GeoM.Reset()
	p.imgOptions.GeoM.Translate(
		(pos.X + (FrameWidth / 2) - MaxCursorDistance),
		(pos.Y + (FrameHeight / 2) - MaxCursorDistance),
	)
	camera.Draw(p.aimOverlayImg, p.imgOptions, screen)

	p.imgOptions.GeoM.Reset()
	p.imgOptions.GeoM.Translate(pos.X, pos.Y-pos.Z)

	camera.Draw(
		playerSubImgs[int(p.animationState)*NumFrames+p.frameIndex],
		p.imgOptions,
		screen,
	)
}

// DrawBelow draws a shadow below the player.
func (p *Player) DrawBelow(screen *ebiten.Image) {
	pos := p.GetPosition()

	scene := (*p.GetScene())
	camera := scene.GetCamera()

	p.imgOptions.GeoM.Reset()
	p.imgOptions.GeoM.Translate(
		pos.X,
		pos.Y+FrameHeight*.75,
	)

	camera.Draw(
		playerShadowImg,
		p.imgOptions,
		screen,
	)
}

// DrawAbove draws a debug overlay if debugging is enabled and active.
func (p *Player) DrawAbove(screen *ebiten.Image) {
	x1, y1, x2, y2 := p.GetCollisionRect()
	p.DrawDebugCollision(screen, x1, y1, x2, y2)
}

// DrawUI draws the player's health bar.
func (p *Player) DrawUI(screen *ebiten.Image) {
	ui.DrawHealthBar(
		screen,
		vectors.Vector2{X: 5, Y: 5},
		p.GetHealth(),
		p.GetMaxHealth(),
	)
}

// Update runs during the game's Update function.
func (p *Player) Update() (err error) {
	p.handleMovement()
	p.handleAnimations()
	p.CheckCollision(*p.GetScene(), *p.GetPosition())

	if p.shootCooldown > 0 {
		p.shootCooldown--
	}

	if p.input.ActionIsPressed(input.ActionShoot) && p.shootCooldown <= 0 {
		p.Shoot()
	}

	return nil
}

// Damage handles damaging the player.
func (p *Player) Damage(amount int, source interfaces.GameObject) {
	if p.state != state.StateDefault {
		return
	}

	scene := (*p.GetScene())
	camera := scene.GetCamera()
	camera.AddTrauma(.5)

	p.HurtableGameObject.Damage(amount, source)
	p.state = state.StateHurt

	pos := *p.GetPosition()
	pos.Z = 0

	srcPosition := (*source.GetPosition())
	srcPosition.Z = 0
	srcPosition.Sub(pos)
	srcPosition.ClampMagnitude(1)
	srcPosition.Bounce()
	srcPosition.Mul(vectors.Vector3{X: Knockback, Y: Knockback, Z: 1})
	p.velocity.Add(srcPosition)
	p.velocity.Z += Knockback / 2
}

// Die handles the player death.
func (p *Player) Die() {
	scene := *p.GetScene()
	scene.RemoveGameObject(p)

	player := NewPlayer(vectors.Vector3{})
	scene.AddGameObject(player)

	scene.SetCameraTarget(player)
}
