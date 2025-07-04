package player

import (
	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/vectors"
)

const (
	VELOCITY_DAMPING float64 = .9
	STOP_THRESHOLD   float64 = .15
	ACCELERATION     float64 = .6
	MAX_SPEED        float64 = 4
)

func (p *Player) handleMovement() {
	// Dampen the X and Y velocity.
	p.velocity.Mul(vectors.Vector3{X: VELOCITY_DAMPING, Y: VELOCITY_DAMPING, Z: 1})

	// If the velocity magnitude is very low, set it to zero.
	// This allows the idle animations to work.
	if p.velocity.Magnitude() < STOP_THRESHOLD {
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
		p.rawInputVelocity.Mul(vectors.Vector3{X: ACCELERATION, Y: ACCELERATION, Z: 1})
	}

	p.velocity.Add(p.rawInputVelocity)

	if _, ok := p.input.PressedActionInfo(input.ActionAimAnalog); ok ||
		p.input.ActionIsPressed(input.ActionAimMouse) {

		p.velocity.ClampMagnitude(RUNNING_THRESHOLD)
	} else {
		p.velocity.ClampMagnitude(MAX_SPEED)
	}

	p.velocity, _ = p.Move(p.velocity)
}
