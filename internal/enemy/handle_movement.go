package enemy

import (
	"github.com/Dobefu/vectors"
)

const (
	// VelocityDamping defines the amount of friction that the enemy has.
	VelocityDamping float64 = .9
	// StopThreshold defines the minimum velocity at which the enemy is
	// considered "stopped".
	StopThreshold float64 = .15
	// Acceleration defines the acceleration of the velocity.
	Acceleration float64 = .6
	// MaxSpeed defines the maximum speed that the enemy can have.
	MaxSpeed float64 = 4
)

func (e *Enemy) handleMovement() {
	// Dampen the velocity.
	e.velocity.Mul(vectors.Vector3{X: VelocityDamping, Y: VelocityDamping, Z: VelocityDamping})

	// If the velocity magnitude is very low, set it to zero.
	// This allows the idle animations to work.
	if e.velocity.Magnitude() < StopThreshold {
		e.velocity.Mul(vectors.Vector3{X: 0, Y: 0, Z: 1})
	}

	pos := *e.GetPosition()

	// Apply gravity.
	if pos.Z > 0 {
		e.velocity.Z -= Acceleration / 2
	}

	e.velocity.ClampMagnitude(MaxSpeed)

	x1, y1, x2, y2 := e.GetCollisionRect()
	e.velocity, _, _ = e.MoveWithCollisionRect(e.velocity, x1, y1, x2, y2)
}
