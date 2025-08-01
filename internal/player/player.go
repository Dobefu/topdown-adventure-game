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
	gameobject.CollidableGameObject
	gameobject.HurtableGameObject

	audioPlayer *audioplayer.AudioPlayer

	bulletPool       []*bullet.Bullet
	shootCooldown    int
	shootCooldownMax int

	movementConfig   gameobject.MovementConfig
	rawInputVelocity vectors.Vector3

	aimOverlayImg *ebiten.Image
	imgOptions    *ebiten.DrawImageOptions
}

// NewPlayer creates a new player.
func NewPlayer(position vectors.Vector3) (player *Player) {
	player = &Player{}

	player.movementConfig = gameobject.DefaultMovementConfig()

	player.aimOverlayImg = ebiten.NewImage(
		gameobject.MaxCameraCursorDistance*2,
		gameobject.MaxCameraCursorDistance*2,
	)
	player.imgOptions = &ebiten.DrawImageOptions{}

	player.SetMaxHealth(MaxHealth)
	player.SetHealth(MaxHealth)

	player.SetIsActive(true)
	player.SetPosition(position)

	player.CollisionRect = gameobject.CollisionRect{
		X1: 4,
		Y1: 19,
		X2: 27,
		Y2: 31,
	}

	player.Input = input.Input.NewHandler(0, input.PlayerKeymap)
	player.Input.GamepadDeadzone = GamepadDeadzone

	player.shootCooldownMax = 20

	player.AnimationState = animation.StateIdleDown

	player.SetOnCollision(func(_, other interfaces.GameObject) {
		// Skip the collision callback if the player hits a bullet they own.
		if bullet, ok := other.(*bullet.Bullet); ok {
			if bullet.GetOwner().GetID() == player.ID {
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

	p.NumFrames = NumFrames
	p.FrameHeight = FrameHeight
	p.FrameWidth = FrameWidth

	scene := *p.GetScene()

	for range 10 {
		b := bullet.NewBullet()

		p.bulletPool = append(p.bulletPool, b)
		scene.AddGameObject(b)
	}

	audioContext := scene.GetGame().GetAudioContext()
	p.audioPlayer = audioplayer.NewAudioPlayerFromBytes(audioContext, playerShootSound)
}

// MoveWithCollision moves the player with collision checks.
func (p *Player) MoveWithCollision(
	velocity vectors.Vector3,
) (newVelocity vectors.Vector3, hasCollided bool, collidedTiles []int) {
	return p.MoveWithCollisionRect(
		velocity,
		p.CollisionRect.X1,
		p.CollisionRect.Y1,
		p.CollisionRect.X2,
		p.CollisionRect.Y2,
	)
}

// Draw runs during the game's Draw function.
func (p *Player) Draw(screen *ebiten.Image) {
	pos := p.Position

	scene := (*p.GetScene())
	camera := scene.GetCamera()

	cameraPos := *p.GetCameraPosition()
	currentPos := p.Position
	currentPos.Z = 0

	currentPos.Add(vectors.Vector3{
		X: FrameWidth / 2,
		Y: FrameHeight / 2,
		Z: 0,
	})

	cameraPos.Sub(currentPos)
	cameraPos.ClampMagnitude(gameobject.MaxCameraOffset)

	p.aimOverlayImg.Clear()

	vector.StrokeLine(
		p.aimOverlayImg,
		gameobject.MaxCameraCursorDistance,
		gameobject.MaxCameraCursorDistance,
		gameobject.MaxCameraCursorDistance+float32(cameraPos.X*2),
		gameobject.MaxCameraCursorDistance+float32(cameraPos.Y*2),
		1,
		color.Alpha{A: uint8(cameraPos.Magnitude() / gameobject.MaxCameraOffset * 255)},
		false,
	)

	p.imgOptions.GeoM.Reset()
	p.imgOptions.GeoM.Translate(
		(pos.X + (FrameWidth / 2) - gameobject.MaxCameraCursorDistance),
		(pos.Y + (FrameHeight / 2) - gameobject.MaxCameraCursorDistance),
	)
	camera.Draw(p.aimOverlayImg, p.imgOptions, screen)

	p.imgOptions.GeoM.Reset()
	p.imgOptions.GeoM.Translate(pos.X, pos.Y-pos.Z)

	camera.Draw(
		playerSubImgs[int(p.AnimationState)*NumFrames+p.FrameIndex],
		p.imgOptions,
		screen,
	)
}

// DrawBelow draws a shadow below the player.
func (p *Player) DrawBelow(screen *ebiten.Image) {
	pos := p.Position

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
	p.DrawDebugCollision(screen)
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
	gameobject.HandleMovement(
		&p.CollidableGameObject,
		&p.Velocity,
		&p.rawInputVelocity,
		p.Input,
		&p.State,
		p.movementConfig,
	)
	gameobject.HandleAnimations(&p.GameObject)
	p.CheckCollision(*p.GetScene(), p.Position)

	if p.shootCooldown > 0 {
		p.shootCooldown--
	}

	if p.Input.ActionIsPressed(input.ActionShoot) && p.shootCooldown <= 0 {
		p.Shoot()
	}

	return nil
}

// Damage handles damaging the player.
func (p *Player) Damage(amount int, source interfaces.GameObject) {
	if p.State != state.StateDefault {
		return
	}

	scene := *p.GetScene()
	camera := scene.GetCamera()
	camera.AddTrauma(.5)

	p.HurtableGameObject.Damage(amount, source)
	p.State = state.StateHurt

	pos := p.Position
	pos.Z = 0

	srcPosition := *source.GetPosition()
	srcPosition.Z = 0
	srcPosition.Sub(pos)
	srcPosition.ClampMagnitude(1)
	srcPosition.Bounce()
	srcPosition.Mul(vectors.Vector3{X: Knockback, Y: Knockback, Z: 1})
	p.Velocity.Add(srcPosition)
	p.Velocity.Z += Knockback / 2
}

// Die handles the player death.
func (p *Player) Die() {
	scene := *p.GetScene()
	scene.RemoveGameObject(p)

	player := NewPlayer(vectors.Vector3{})
	scene.AddGameObject(player)

	scene.SetCameraTarget(player)
}
