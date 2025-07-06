package game_object

import (
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/vectors"
)

type CollidableGameObject struct {
	GameObject
	interfaces.CollidableGameObject
}

func (c *CollidableGameObject) MoveWithCollision(
	velocity vectors.Vector3,
) (newVelocity vectors.Vector3, hasCollided bool) {
	x1, y1, x2, y2 := c.GetCollisionRect()

	return c.MoveWithCollisionRect(velocity, x1, y1, x2, y2)
}

func (c *CollidableGameObject) GetCollisionRect() (x1, y1, x2, y2 float64) {
	return 0, 0, 31, 31
}

func (c *CollidableGameObject) CheckCollision() {
	x1, y1, x2, y2 := c.GetCollisionRect()

	c.CheckCollisionWithCollisionRect(x1, y1, x2, y2)
}

func (c *CollidableGameObject) CheckCollisionWithCollisionRect(
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
) {

}
