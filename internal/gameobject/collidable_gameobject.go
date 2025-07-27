package gameobject

import (
	"image/color"

	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/vectors"
	"github.com/hajimehoshi/ebiten/v2"
)

// CollidableGameObject defines a game object that can collide.
type CollidableGameObject struct {
	GameObject

	OnCollision func(self interfaces.GameObject, other interfaces.GameObject)

	debugCollisionImage        *ebiten.Image
	debugCollisionImageOptions *ebiten.DrawImageOptions
}

// Init initializes the game object.
func (c *CollidableGameObject) Init() {
	c.GameObject.Init()

	scene := *c.GetScene()

	if scene.GetGame().GetIsDebugEnabled() {
		c.debugCollisionImageOptions = &ebiten.DrawImageOptions{}
		c.debugCollisionImageOptions.Blend = ebiten.Blend{
			BlendFactorSourceRGB: ebiten.BlendFactorSourceAlpha,
		}
	}
}

// DrawAbove handles drawing above the game object.
func (c *CollidableGameObject) DrawAbove(screen *ebiten.Image) {
	c.GameObject.DrawAbove(screen)
}

// DrawDebugCollision handles drawing of the collision debug overlay.
func (c *CollidableGameObject) DrawDebugCollision(
	screen *ebiten.Image,
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
) {
	scene := *c.GetScene()

	if !scene.GetGame().GetIsDebugActive() {
		return
	}

	if c.debugCollisionImage == nil {
		c.debugCollisionImage = ebiten.NewImage(int(x2-x1), int(y2-y1))
		c.debugCollisionImage.Fill(color.RGBA{R: 255, G: 0, B: 0, A: 128})
	}

	camera := scene.GetCamera()

	c.debugCollisionImageOptions.GeoM.Reset()
	c.debugCollisionImageOptions.GeoM.Translate(
		c.Position.X+x1,
		c.Position.Y+y1,
	)

	camera.Draw(
		c.debugCollisionImage,
		c.debugCollisionImageOptions,
		screen,
	)
}

// GetOnCollision gets the current OnCollision callback function.
func (c *CollidableGameObject) GetOnCollision() (
	callback func(self interfaces.GameObject, other interfaces.GameObject),
) {
	return c.OnCollision
}

// SetOnCollision sets the current OnCollision callback function.
func (c *CollidableGameObject) SetOnCollision(
	callback func(self interfaces.GameObject, other interfaces.GameObject),
) {
	c.OnCollision = callback
}

// MoveWithCollision moves the game object with collision checks.
func (c *CollidableGameObject) MoveWithCollision(
	velocity vectors.Vector3,
) (newVelocity vectors.Vector3, hasCollided bool, collidedTiles []int) {
	x1, y1, x2, y2 := c.GetCollisionRect()

	return c.MoveWithCollisionRect(velocity, x1, y1, x2, y2)
}

// GetCollisionRect gets the four points of the collision rectangle.
func (c *CollidableGameObject) GetCollisionRect() (x1, y1, x2, y2 float64) {
	return 0, 0, 31, 31
}

// CheckCollision checks if the game object collides with another game object.
func (c *CollidableGameObject) CheckCollision(
	scene interfaces.Scene,
	position vectors.Vector3,
) {
	if scene == nil {
		return
	}

	x1, y1, x2, y2 := c.GetCollisionRect()
	activeGameObjects := scene.GetActiveGameObjects()

	c.CheckCollisionWithCollisionRect(x1, y1, x2, y2, activeGameObjects, position)
}

// CheckCollisionWithCollisionRect checks for collision with a bounding box.
func (c *CollidableGameObject) CheckCollisionWithCollisionRect(
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
	activeGameObjects []interfaces.GameObject,
	position vectors.Vector3,
) {
	if position.IsZero() {
		return
	}

	for _, activeGameObject := range activeGameObjects {
		type iCollidable interface {
			GetCollisionRect() (x1, y1, x2, y2 float64)
		}

		// Skip non-collidable gameObjects.
		if _, ok := activeGameObject.(iCollidable); !ok {
			continue
		}

		// Skip GameObjects with the same ID.
		if c.ID == activeGameObject.GetID() {
			continue
		}

		collidable := activeGameObject.(iCollidable)

		otherX1, otherY1, otherX2, otherY2 := collidable.GetCollisionRect()
		otherPosition := activeGameObject.GetPosition()

		if otherPosition.IsZero() {
			continue
		}

		rect1X1 := position.X + x1
		rect1Y1 := position.Y + y1
		rect1X2 := position.X + x2
		rect1Y2 := position.Y + y2

		rect2X1 := otherPosition.X + otherX1
		rect2Y1 := otherPosition.Y + otherY1
		rect2X2 := otherPosition.X + otherX2
		rect2Y2 := otherPosition.Y + otherY2

		if rect1X1 < rect2X2 &&
			rect1X2 > rect2X1 &&
			rect1Y1 < rect2Y2 &&
			rect1Y2 > rect2Y1 {

			if c.OnCollision != nil {
				c.OnCollision(&c.GameObject, activeGameObject)
			}
		}
	}
}
