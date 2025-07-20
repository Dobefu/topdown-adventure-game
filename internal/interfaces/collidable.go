package interfaces

// Collidable defines the interface for a collidable.
type Collidable interface {
	GetCollisionRect() (x1, y1, x2, y2 float64)
}
