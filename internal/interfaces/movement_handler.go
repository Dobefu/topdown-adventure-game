package interfaces

import (
	"github.com/Dobefu/vectors"
)

// MovementHandler defines methods for objects that can handle movement.
type MovementHandler interface {
	GetPosition() *vectors.Vector3
	GetCollisionRect() (x1, y1, x2, y2 float64)
	MoveWithCollisionRect(velocity vectors.Vector3, x1, y1, x2, y2 float64) (newVelocity vectors.Vector3, hasCollided bool, collidedTiles []int)
}
