package player

import (
	"math"

	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/vectors"
)

func (p *Player) handleMovement() {
	// Dampen the X and Y velocity.
	p.velocity.Mul(vectors.Vector3{X: .9, Y: .9, Z: 1})

	if math.Abs(p.velocity.X) < .25 && math.Abs(p.velocity.Y) < .25 {
		p.velocity.Mul(vectors.Vector3{X: 0, Y: 0, Z: 1})
	}

	if p.input.ActionIsPressed(input.ActionMoveLeft) {
		p.velocity.X -= .3
	}

	if p.input.ActionIsPressed(input.ActionMoveRight) {
		p.velocity.X += .3
	}

	if p.input.ActionIsPressed(input.ActionMoveUp) {
		p.velocity.Y -= .3
	}

	if p.input.ActionIsPressed(input.ActionMoveDown) {
		p.velocity.Y += .3
	}

	pos := p.GetPosition()
	pos.Add(p.velocity)
}
