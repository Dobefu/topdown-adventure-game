// Package enemy provides functionality for enemies.
package enemy

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/Dobefu/topdown-adventure-game/internal/animation"
	"github.com/Dobefu/topdown-adventure-game/internal/bullet"
	"github.com/Dobefu/topdown-adventure-game/internal/gameobject"
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/topdown-adventure-game/internal/state"
	"github.com/Dobefu/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed img/enemy.png
	enemyImgBytes []byte
	enemyImg      *ebiten.Image
	enemySubImgs  []*ebiten.Image

	//go:embed img/shadow.png
	enemyShadowImgBytes []byte
	enemyShadowImg      *ebiten.Image
)

const (
	// FrameWidth defines the width of an animation frame.
	FrameWidth = 32
	// FrameHeight defines the height of an animation frame.
	FrameHeight = 32
	// NumFrames defines number of animation frames in the spritesheet.
	NumFrames = 16
	// Knockback defines the amount of knockback the enemy gets when hit.
	Knockback = 10

	// MaxHealth defines the base max health that the enemy has.
	MaxHealth = 10
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(enemyImgBytes))

	if err != nil {
		log.Fatal(err)
	}

	enemyImg = ebiten.NewImageFromImage(img)
	enemySubImgs = make(
		[]*ebiten.Image,
		enemyImg.Bounds().Dy()/FrameHeight*NumFrames,
	)

	for state := range enemyImg.Bounds().Dy() / FrameHeight {
		for frame := range NumFrames {
			key := state*NumFrames + frame

			enemySubImgs[key] = ebiten.NewImageFromImage(
				enemyImg.SubImage(
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

	img, _, err = image.Decode(bytes.NewReader(enemyShadowImgBytes))

	if err != nil {
		log.Fatal(err)
	}

	enemyShadowImg = ebiten.NewImageFromImage(img)
}

// Enemy struct defines a base enemy type.
type Enemy struct {
	gameobject.CollidableGameObject
	gameobject.HostileGameObject
	gameobject.HurtableGameObject

	movementConfig gameobject.MovementConfig
	velocity       vectors.Vector3

	imgOptions *ebiten.DrawImageOptions

	frameCount     int
	frameIndex     int
	animationState animation.State
	state          state.State
}

// NewEnemy creates a new enemy.
func NewEnemy(position vectors.Vector3) (enemy *Enemy) {
	enemy = &Enemy{}

	enemy.movementConfig = gameobject.DefaultMovementConfig()

	enemy.imgOptions = &ebiten.DrawImageOptions{}

	enemy.SetMaxHealth(MaxHealth)
	enemy.SetHealth(MaxHealth)

	enemy.SetIsActive(true)
	enemy.SetPosition(position)

	enemy.CollisionRect = gameobject.CollisionRect{
		X1: 4,
		Y1: 19,
		X2: 27,
		Y2: 31,
	}

	enemy.animationState = animation.StateIdleDown
	enemy.state = state.StateDefault

	enemy.SetOnCollision(func(self, other interfaces.GameObject) {
		other.Damage(2, self)
	})

	enemy.SetDeathCallback(enemy.Die)

	return enemy
}

// Init initializes a new enemy.
func (e *Enemy) Init() {
	e.GameObject.Init()
	e.CollidableGameObject.Init()

	scene := *e.GetScene()

	for range 10 {
		b := bullet.NewBullet()

		scene.AddGameObject(b)
	}
}

// MoveWithCollision moves the enemy, and checks for collision.
func (e *Enemy) MoveWithCollision(
	velocity vectors.Vector3,
) (newVelocity vectors.Vector3, hasCollided bool, collidedTiles []int) {
	return e.MoveWithCollisionRect(
		velocity,
		e.CollisionRect.X1,
		e.CollisionRect.Y1,
		e.CollisionRect.X2,
		e.CollisionRect.Y2,
	)
}

// Draw draws the enemy.
func (e *Enemy) Draw(screen *ebiten.Image) {
	pos := e.Position

	scene := (*e.GetScene())
	camera := scene.GetCamera()

	e.imgOptions.GeoM.Reset()
	e.imgOptions.GeoM.Translate(pos.X, pos.Y-pos.Z)

	camera.Draw(
		enemySubImgs[int(e.animationState)*NumFrames+e.frameIndex],
		e.imgOptions,
		screen,
	)
}

// DrawBelow draws a shadow beneath the enemy.
func (e *Enemy) DrawBelow(screen *ebiten.Image) {
	scene := (*e.GetScene())
	camera := scene.GetCamera()

	e.imgOptions.GeoM.Reset()
	e.imgOptions.GeoM.Translate(
		e.Position.X,
		e.Position.Y+FrameHeight*.75,
	)

	camera.Draw(
		enemyShadowImg,
		e.imgOptions,
		screen,
	)
}

// DrawAbove draws above the enemy.
func (e *Enemy) DrawAbove(screen *ebiten.Image) {
	e.DrawDebugCollision(screen)
}

// Update runs during the update function of the game.
func (e *Enemy) Update() (err error) {
	gameobject.HandleMovement(
		&e.CollidableGameObject,
		&e.velocity,
		nil,
		nil,
		&e.state,
		e.movementConfig,
	)
	e.handleAnimations()
	e.CheckCollision(*e.GetScene(), e.Position)

	return nil
}

// Damage handles the damaging of the enemy.
func (e *Enemy) Damage(amount int, source interfaces.GameObject) {
	if e.state != state.StateDefault {
		return
	}

	e.HurtableGameObject.Damage(amount, source)
	e.state = state.StateHurt

	pos := e.Position
	pos.Z = 0

	srcPosition := *source.GetPosition()
	srcPosition.Z = 0
	srcPosition.Sub(pos)
	srcPosition.ClampMagnitude(1)
	srcPosition.Bounce()
	srcPosition.Mul(vectors.Vector3{X: Knockback, Y: Knockback, Z: 1})
	e.velocity.Add(srcPosition)
	e.velocity.Z += Knockback / 2
}

// Die handles the death of the enemy.
func (e *Enemy) Die() {
	e.SetIsActive(false)
}
