// Package bullet provides functionality for bullets.
package bullet

import (
	"bytes"
	_ "embed"
	"image"
	"log"
	"math"

	"github.com/Dobefu/topdown-adventure-game/internal/animation"
	"github.com/Dobefu/topdown-adventure-game/internal/fastrand"
	"github.com/Dobefu/topdown-adventure-game/internal/gameobject"
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/topdown-adventure-game/internal/particles/pixel"
	"github.com/Dobefu/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed img/bullet.png
	bulletImgBytes []byte
	bulletImg      *ebiten.Image
	bulletSubImgs  []*ebiten.Image
)

const (
	// FrameWidth defines the width of an animation frame.
	FrameWidth = 32
	// FrameHeight defines the height of an animation frame.
	FrameHeight = 32
	// NumFrames defines number of animation frames in the spritesheet.
	NumFrames = 16
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(bulletImgBytes))

	if err != nil {
		log.Fatal(err)
	}

	bulletImg = ebiten.NewImageFromImage(img)
	bulletSubImgs = make(
		[]*ebiten.Image,
		bulletImg.Bounds().Dy()/FrameHeight*NumFrames,
	)

	for state := range bulletImg.Bounds().Dy() / FrameHeight {
		for frame := range NumFrames {
			key := state*NumFrames + frame

			bulletSubImgs[key] = ebiten.NewImageFromImage(
				bulletImg.SubImage(
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
}

// Bullet struct provides a single bullet.
type Bullet struct {
	gameobject.CollidableGameObject

	owner interfaces.GameObject

	velocity vectors.Vector3

	imgOptions     *ebiten.DrawImageOptions
	frameCount     int
	frameIndex     int
	animationState animation.State

	trailParticles []*pixel.Pixel
}

// NewBullet creates a new bullet.
func NewBullet() (bullet *Bullet) {
	bullet = &Bullet{
		velocity: vectors.Vector3{},
	}

	bullet.imgOptions = &ebiten.DrawImageOptions{}

	bullet.SetIsActive(false)

	bullet.CollisionRect = gameobject.CollisionRect{
		X1: 12,
		Y1: 12,
		X2: 19,
		Y2: 19,
	}

	bullet.SetOnCollision(func(self, other interfaces.GameObject) {
		// Skip collision with the owner of the bullet.
		if other.GetID() == bullet.owner.GetID() {
			return
		}

		other.Damage(2, self)
		bullet.SetIsActive(false)
	})

	return bullet
}

// Init initializes a bullet instance.
func (b *Bullet) Init() {
	b.GameObject.Init()
	b.CollidableGameObject.Init()

	scene := *b.GetScene()

	for range 100 {
		p := pixel.NewPixel(vectors.Vector3{X: 10, Y: 10, Z: 0})

		b.trailParticles = append(b.trailParticles, p)
		scene.AddGameObject(p)
	}
}

// GetOwner gets the owner of a bullet.
func (b *Bullet) GetOwner() (owner interfaces.GameObject) {
	return b.owner
}

// SetOwner sets the owner of a bullet.
// The owner of a bullet cannot get hurt by it.
func (b *Bullet) SetOwner(owner interfaces.GameObject) {
	b.owner = owner
}

// Fire fires a single bullet.
func (b *Bullet) Fire(
	from vectors.Vector3,
	angle float64,
	velocity vectors.Vector3,
) {
	b.SetIsActive(true)
	b.SetPosition(from)

	velocity.Add(vectors.Vector3{
		X: -math.Cos(angle+math.Pi/2) * 15,
		Y: -math.Sin(angle+math.Pi/2) * 15,
		Z: 0,
	})

	b.velocity = velocity
}

// SetVelocity sets the current velocity of a bullet.
func (b *Bullet) SetVelocity(velocity vectors.Vector3) {
	b.velocity = velocity
}

// MoveWithCollision moves the bullet, and checks for collision.
func (b *Bullet) MoveWithCollision(
	velocity vectors.Vector3,
) (newVelocity vectors.Vector3, hasCollided bool) {
	newVelocity, hasCollided, _ = b.MoveWithCollisionRect(
		velocity,
		b.CollisionRect.X1,
		b.CollisionRect.Y1,
		b.CollisionRect.X2,
		b.CollisionRect.Y2,
	)

	if hasCollided {
		b.SetIsActive(false)
	}

	return newVelocity, hasCollided
}

// Draw draws the bullet.
func (b *Bullet) Draw(screen *ebiten.Image) {
	scene := (*b.GetScene())
	camera := scene.GetCamera()

	b.imgOptions.GeoM.Reset()
	b.imgOptions.GeoM.Translate(b.Position.X, b.Position.Y)

	camera.Draw(
		bulletSubImgs[int(b.animationState)*NumFrames+b.frameIndex],
		b.imgOptions,
		screen,
	)
}

// DrawAbove draws above the bullet.
func (b *Bullet) DrawAbove(screen *ebiten.Image) {
	b.DrawDebugCollision(screen)
}

// Update runs during the update function of the game.
func (b *Bullet) Update() (err error) {
	b.frameCount++

	// Change the animation frame every 3 game ticks.
	if (b.frameCount % 3) == 0 {
		b.frameIndex++

		if b.frameIndex >= NumFrames {
			b.frameIndex = 0
		}
	}

	angle := b.velocity.AngleDegrees()

	b.animationState = animation.State(
		int((angle+22.5)/45) % 8,
	)

	b.CheckCollision(*b.GetScene(), b.Position)
	b.MoveWithCollision(b.velocity)

	for _, particle := range b.trailParticles {
		if particle.GetIsActive() {
			continue
		}

		particle.Position = vectors.Vector3{
			X: b.Position.X + 16 + float64(int(fastrand.Rand.Next()>>28)-8),
			Y: b.Position.Y + 16 + float64(int(fastrand.Rand.Next()>>28)-8),
			Z: 0,
		}
		particle.SetVelocity(vectors.Vector3{
			X: float64(int(fastrand.Rand.Next()>>28)-8) / 10,
			Y: float64(int(fastrand.Rand.Next()>>28)-8) / 10,
			Z: 0,
		})
		particle.SetLifetime(20)
		particle.SetIsActive(true)

		return
	}

	return nil
}
