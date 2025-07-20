package interfaces

import "github.com/Dobefu/vectors"

// Bullet defines the interface for bullets.
type Bullet interface {
	Fire(from vectors.Vector3, angle float64, velocity vectors.Vector3)
	SetVelocity(vectors.Vector3)
	GetOwner() (owner GameObject)
	SetOwner(owner GameObject)
}
