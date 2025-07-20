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
	// Dampen the velocity.
	e.velocity.Mul(vectors.Vector3{X: VELOCITY_DAMPING, Y: VELOCITY_DAMPING, Z: VELOCITY_DAMPING})

	// If the velocity magnitude is very low, set it to zero.
	// This allows the idle animations to work.
	if e.velocity.Magnitude() < STOP_THRESHOLD {
		e.velocity.Mul(vectors.Vector3{X: 0, Y: 0, Z: 1})
	}

	pos := *e.GetPosition()

	// Apply gravity.
	if pos.Z > 0 {
		e.velocity.Z -= ACCELERATION / 2
	}

	e.velocity.ClampMagnitude(MAX_SPEED)

	x1, y1, x2, y2 := e.GetCollisionRect()
	e.velocity, _ = e.MoveWithCollisionRect(e.velocity, x1, y1, x2, y2)
}
