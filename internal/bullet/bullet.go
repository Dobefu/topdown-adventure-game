package bullet

import (
	"bytes"
	_ "embed"
	"image"
	"log"
	"math"

	"github.com/Dobefu/topdown-adventure-game/internal/animation"
	"github.com/Dobefu/topdown-adventure-game/internal/game_object"
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
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
	FRAME_WIDTH  = 32
	FRAME_HEIGHT = 32
	NUM_FRAMES   = 16
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(bulletImgBytes))

	if err != nil {
		log.Fatal(err)
	}

	bulletImg = ebiten.NewImageFromImage(img)
	bulletSubImgs = make(
		[]*ebiten.Image,
		bulletImg.Bounds().Dy()/FRAME_HEIGHT*NUM_FRAMES,
	)

	for state := range bulletImg.Bounds().Dy() / FRAME_HEIGHT {
		for frame := range NUM_FRAMES {
			key := state*NUM_FRAMES + frame

			bulletSubImgs[key] = ebiten.NewImageFromImage(
				bulletImg.SubImage(
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
}

type Bullet struct {
	interfaces.Bullet
	game_object.GameObject

	velocity vectors.Vector3

	imgOptions     *ebiten.DrawImageOptions
	frameCount     int
	frameIndex     int
	animationState animation.AnimationState
}

func NewBullet() (bullet *Bullet) {
	bullet = &Bullet{
		velocity: vectors.Vector3{},
	}

	bullet.imgOptions = &ebiten.DrawImageOptions{}

	bullet.SetIsActive(false)

	return bullet
}

func (b *Bullet) GetCollisionRect() (x1, y1, x2, y2 float64) {
	return 12, 12, 19, 19
}

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

func (b *Bullet) SetVelocity(velocity vectors.Vector3) {
	b.velocity = velocity
}

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

func (b *Bullet) Draw(screen *ebiten.Image) {
	pos := b.GetPosition()

	scene := (*b.GetScene())
	camera := scene.GetCamera()

	b.imgOptions.GeoM.Reset()
	b.imgOptions.GeoM.Translate(math.Round(pos.X), math.Round(pos.Y))

	camera.Draw(
		bulletSubImgs[int(b.animationState)*NUM_FRAMES+b.frameIndex],
		b.imgOptions,
		screen,
	)
}

func (b *Bullet) Update() (err error) {
	b.frameCount += 1

	// Change the animation frame every 4 game ticks.
	if (b.frameCount % 4) == 0 {
		b.frameIndex += 1

		if b.frameIndex >= NUM_FRAMES {
			b.frameIndex = 0
		}
	}

	angle := b.velocity.AngleDegrees()

	b.animationState = animation.AnimationState(
		int((angle+22.5)/45) % 8,
	)

	b.Move(b.velocity)

	return nil
}
