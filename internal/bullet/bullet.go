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
	"github.com/Dobefu/topdown-adventure-game/internal/game_object"
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
	interfaces.Bullet
	game_object.CollidableGameObject

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

	bullet.SetOnCollision(func(self, other interfaces.GameObject) {
		// Skip collision with the owner of the bullet.
		if other.GetID() == bullet.owner.GetID() {
			return
		}

		if hurtable, ok := other.(interfaces.HurtableGameObject); ok {
			hurtable.Damage(2, self)
		}

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

// GetCollisionRect gets the four points of the collision rectangle.
func (b *Bullet) GetCollisionRect() (x1, y1, x2, y2 float64) {
	return 12, 12, 19, 19
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
	x1, y1, x2, y2 := b.GetCollisionRect()
	newVelocity, hasCollided = b.MoveWithCollisionRect(velocity, x1, y1, x2, y2)

	if hasCollided {
		b.SetIsActive(false)
	}

	return newVelocity, hasCollided
}

// Draw draws the bullet.
func (b *Bullet) Draw(screen *ebiten.Image) {
	pos := b.GetPosition()

	scene := (*b.GetScene())
	camera := scene.GetCamera()

	b.imgOptions.GeoM.Reset()
	b.imgOptions.GeoM.Translate(pos.X, pos.Y)

	camera.Draw(
		bulletSubImgs[int(b.animationState)*NumFrames+b.frameIndex],
		b.imgOptions,
		screen,
	)
}

// DrawAbove draws above the bullet.
func (b *Bullet) DrawAbove(screen *ebiten.Image) {
	x1, y1, x2, y2 := b.GetCollisionRect()
	b.DrawDebugCollision(screen, x1, y1, x2, y2)
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

	b.CheckCollision(*b.GetScene(), *b.GetPosition())
	b.MoveWithCollision(b.velocity)

	position := b.GetPosition()

	for _, particle := range b.trailParticles {
		if particle.GetIsActive() {
			continue
		}

		particle.SetPosition(vectors.Vector3{
			X: position.X + 16 + float64(int(fastrand.Rand.Next()>>28)-8),
			Y: position.Y + 16 + float64(int(fastrand.Rand.Next()>>28)-8),
			Z: 0,
		})
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
