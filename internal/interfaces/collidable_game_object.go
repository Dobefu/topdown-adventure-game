package interfaces

import (
	"github.com/Dobefu/vectors"
)

type CollidableGameObject interface {
	MoveWithCollision(velocity vectors.Vector3)
	GetCollisionRect() (x1, y1, x2, y2 float64)
	CheckCollision()
	CheckCollisionWithCollisionRect(x1 float64, y1 float64, x2 float64, y2 float64)
}
