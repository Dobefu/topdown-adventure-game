package particles

import (
	"github.com/Dobefu/topdown-adventure-game/internal/gameobject"
	"github.com/Dobefu/vectors"
)

// Particle struct provides a base particle, which is meant to be embedded.
// Do not use this struct on its own.
type Particle struct {
	gameobject.GameObject

	lifetime int
	velocity vectors.Vector3
}

// Update runs during the game Update method.
func (p *Particle) Update() {
	p.lifetime--

	if p.lifetime <= 0 {
		p.SetIsActive(false)
	}

	p.Position.Add(p.velocity)
}

// GetLifetime gets the lifetime that a particle has left.
func (p *Particle) GetLifetime() (lifetime int) {
	return p.lifetime
}

// SetLifetime sets the lifetime that a particle has left.
func (p *Particle) SetLifetime(lifetime int) {
	p.lifetime = lifetime
}

// GetVelocity gets the current velocity of a particle.
func (p *Particle) GetVelocity() (velocity vectors.Vector3) {
	return p.velocity
}

// SetVelocity sets the velocity of a particle.
func (p *Particle) SetVelocity(velocity vectors.Vector3) {
	p.velocity = velocity
}
