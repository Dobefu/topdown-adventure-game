package enemy

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/Dobefu/topdown-adventure-game/internal/animation"
	"github.com/Dobefu/topdown-adventure-game/internal/bullet"
	"github.com/Dobefu/topdown-adventure-game/internal/game_object"
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
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
	FRAME_WIDTH      = 32
	FRAME_HEIGHT     = 32
	NUM_FRAMES       = 16
	GAMEPAD_DEADZONE = .1

	MAX_HEALTH = 10
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(enemyImgBytes))

	if err != nil {
		log.Fatal(err)
	}

	enemyImg = ebiten.NewImageFromImage(img)
	enemySubImgs = make(
		[]*ebiten.Image,
		enemyImg.Bounds().Dy()/FRAME_HEIGHT*NUM_FRAMES,
	)

	for state := range enemyImg.Bounds().Dy() / FRAME_HEIGHT {
		for frame := range NUM_FRAMES {
			key := state*NUM_FRAMES + frame

			enemySubImgs[key] = ebiten.NewImageFromImage(
				enemyImg.SubImage(
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

	img, _, err = image.Decode(bytes.NewReader(enemyShadowImgBytes))

	if err != nil {
		log.Fatal(err)
	}

	enemyShadowImg = ebiten.NewImageFromImage(img)
}

type Enemy struct {
	game_object.HostileGameObject
	game_object.HurtableGameObject

	velocity         vectors.Vector3
	rawInputVelocity vectors.Vector3

	imgOptions *ebiten.DrawImageOptions

	frameCount     int
	frameIndex     int
	animationState animation.AnimationState
}

func NewEnemy(position vectors.Vector3) (enemy *Enemy) {
	enemy = &Enemy{}

	enemy.imgOptions = &ebiten.DrawImageOptions{}

	enemy.SetMaxHealth(MAX_HEALTH)
	enemy.SetHealth(MAX_HEALTH)

	enemy.SetIsActive(true)
	enemy.SetPosition(position)

	enemy.animationState = animation.AnimationStateIdleDown

	enemy.SetOnCollision(func(self, other interfaces.GameObject) {
		if hurtable, ok := other.(interfaces.HurtableGameObject); ok {
			hurtable.Damage(2, self)
		}
	})

	enemy.SetDeathCallback(enemy.Die)

	return enemy
}

func (e *Enemy) Init() {
	e.GameObject.Init()
	e.CollidableGameObject.Init()

	scene := *e.GetScene()

	for range 10 {
		b := bullet.NewBullet()

		scene.AddGameObject(b)
	}
}

func (e *Enemy) GetCollisionRect() (x1, y1, x2, y2 float64) {
	return 4, 19, 27, 31
}

func (e *Enemy) MoveWithCollision(
	velocity vectors.Vector3,
) (newVelocity vectors.Vector3, hasCollided bool) {
	x1, y1, x2, y2 := e.GetCollisionRect()

	return e.MoveWithCollisionRect(velocity, x1, y1, x2, y2)
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	pos := e.GetPosition()

	scene := (*e.GetScene())
	camera := scene.GetCamera()

	e.imgOptions.GeoM.Reset()
	e.imgOptions.GeoM.Translate(pos.X, pos.Y)

	camera.Draw(
		enemySubImgs[int(e.animationState)*NUM_FRAMES+e.frameIndex],
		e.imgOptions,
		screen,
	)
}

func (e *Enemy) DrawBelow(screen *ebiten.Image) {
	pos := e.GetPosition()

	scene := (*e.GetScene())
	camera := scene.GetCamera()

	e.imgOptions.GeoM.Reset()
	e.imgOptions.GeoM.Translate(
		pos.X,
		pos.Y+FRAME_HEIGHT*.75,
	)

	camera.Draw(
		enemyShadowImg,
		e.imgOptions,
		screen,
	)
}

func (e *Enemy) DrawAbove(screen *ebiten.Image) {
	x1, y1, x2, y2 := e.GetCollisionRect()
	e.DrawDebugCollision(screen, x1, y1, x2, y2)
}

func (e *Enemy) Update() (err error) {
	e.handleMovement()
	e.handleAnimations()
	e.CheckCollision(*e.GetScene(), *e.GetPosition())

	return nil
}

func (e *Enemy) Die() {
	e.SetIsActive(false)
}
