package enemy

import (
	"github.com/Dobefu/vectors"
)

const (
	VELOCITY_DAMPING float64 = .9
	STOP_THRESHOLD   float64 = .15
	ACCELERATION     float64 = .6
	MAX_SPEED        float64 = 4
)

func (e *Enemy) handleMovement() {
	// Dampen the X and Y velocity.
	e.velocity.Mul(vectors.Vector3{X: VELOCITY_DAMPING, Y: VELOCITY_DAMPING, Z: 1})

	// If the velocity magnitude is very low, set it to zero.
	// This allows the idle animations to work.
	if e.velocity.Magnitude() < STOP_THRESHOLD {
		e.velocity.Mul(vectors.Vector3{X: 0, Y: 0, Z: 1})
	}

	e.rawInputVelocity.Clear()

	if e.rawInputVelocity.Magnitude() > 0 {
		e.rawInputVelocity.Normalize()
		e.rawInputVelocity.Mul(vectors.Vector3{X: ACCELERATION, Y: ACCELERATION, Z: 1})
	}

	e.velocity.Add(e.rawInputVelocity)
	e.velocity.ClampMagnitude(MAX_SPEED)

	e.velocity, _ = e.MoveWithCollision(e.velocity)
}
