package particles

import (
	"github.com/Dobefu/topdown-adventure-game/internal/game_object"
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
)

type Particle struct {
	game_object.GameObject
	interfaces.Particle

	lifetime int
}

func (p *Particle) Update() {
	p.lifetime -= 1

	if p.lifetime <= 0 {
		p.SetIsActive(false)
	}
}

func (p *Particle) GetLifetime() (lifetime int) {
	return p.lifetime
}

func (p *Particle) SetLifetime(lifetime int) {
	p.lifetime = lifetime
}
