package player

import (
	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/vectors"
)

func (p *Player) handleMovement() {
	// Dampen the X and Y velocity.
	p.velocity.Mul(vectors.Vector3{X: .9, Y: .9, Z: 1})

	// If the velocity magnitude is very low, set it to zero.
	// This allows the idle animations to work.
	if p.velocity.Magnitude() < .15 {
		p.velocity.Mul(vectors.Vector3{X: 0, Y: 0, Z: 1})
	}

	p.rawInputVelocity.Clear()

	if info, ok := p.input.PressedActionInfo(input.ActionMoveAnalog); ok {
		p.rawInputVelocity.Add(vectors.Vector3{
			X: info.Pos.X,
			Y: info.Pos.Y,
			Z: 0,
		})
	} else {
		if p.input.ActionIsPressed(input.ActionMoveLeft) {
			p.rawInputVelocity.X -= 1
		}

		if p.input.ActionIsPressed(input.ActionMoveRight) {
			p.rawInputVelocity.X += 1
		}

		if p.input.ActionIsPressed(input.ActionMoveUp) {
			p.rawInputVelocity.Y -= 1
		}

		if p.input.ActionIsPressed(input.ActionMoveDown) {
			p.rawInputVelocity.Y += 1
		}
	}

	if p.rawInputVelocity.Magnitude() > 0 {
		p.rawInputVelocity.Normalize()
		p.rawInputVelocity.Mul(vectors.Vector3{X: .6, Y: .6, Z: 1})
	}

	p.velocity.Add(p.rawInputVelocity)
	p.velocity.ClampMagnitude(4)

	pos := p.GetPosition()
	pos.Add(p.velocity)
}
