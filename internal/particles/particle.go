package particles

import (
	"github.com/Dobefu/topdown-adventure-game/internal/game_object"
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/vectors"
)

type Particle struct {
	game_object.GameObject
	interfaces.Particle

	lifetime int
	velocity vectors.Vector3
}

func (p *Particle) Update() {
	p.lifetime -= 1

	if p.lifetime <= 0 {
		p.SetIsActive(false)
	}

	position := p.GetPosition()

	position.Add(p.velocity)
}

func (p *Particle) GetLifetime() (lifetime int) {
	return p.lifetime
}

func (p *Particle) SetLifetime(lifetime int) {
	p.lifetime = lifetime
}

func (p *Particle) GetVelocity() (velocity vectors.Vector3) {
	return p.velocity
}

func (p *Particle) SetVelocity(velocity vectors.Vector3) {
	p.velocity = velocity
}
