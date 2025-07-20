package interfaces

import "github.com/Dobefu/vectors"

// Particle defines the interface for a particle.
type Particle interface {
	Update()
	GetLifetime() (lifetime int)
	SetLifetime(lifetime int)
	GetVelocity() (velocity vectors.Vector3)
	SetVelocity(velocity vectors.Vector3)
}
