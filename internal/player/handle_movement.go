package player

import (
	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/topdown-adventure-game/internal/state"
	"github.com/Dobefu/vectors"
)

const (
	VELOCITY_DAMPING float64 = .9
	STOP_THRESHOLD   float64 = .15
	ACCELERATION     float64 = .6
	MAX_SPEED        float64 = 4
)

func (p *Player) handleMovement() {
	// Dampen the velocity.
	p.velocity.Mul(vectors.Vector3{X: VELOCITY_DAMPING, Y: VELOCITY_DAMPING, Z: VELOCITY_DAMPING})

	// If the velocity magnitude is very low, set it to zero.
	// This allows the idle animations to work.
	if p.velocity.Magnitude() < STOP_THRESHOLD {
		p.velocity.Mul(vectors.Vector3{X: 0, Y: 0, Z: 1})
	}

	p.rawInputVelocity.Clear()

	p.handleInput()

	if p.rawInputVelocity.Magnitude() > 0 {
		p.rawInputVelocity.Normalize()
		p.rawInputVelocity.Mul(vectors.Vector3{X: ACCELERATION, Y: ACCELERATION, Z: 1})
	}

	p.velocity.Add(p.rawInputVelocity)

	pos := *p.GetPosition()

	// Apply gravity.
	if pos.Z > 0 {
		p.velocity.Z -= ACCELERATION / 2
	}

	if p.state != state.StateHurt {
		if _, ok := p.input.PressedActionInfo(input.ActionAimAnalog); ok ||
			p.input.ActionIsPressed(input.ActionAimMouse) {

			p.velocity.ClampMagnitude(RUNNING_THRESHOLD)
		} else {
			p.velocity.ClampMagnitude(MAX_SPEED)
		}
	}

	x1, y1, x2, y2 := p.GetCollisionRect()
	p.velocity, _ = p.MoveWithCollisionRect(p.velocity, x1, y1, x2, y2)
}

func (p *Player) handleInput() {
	if p.state != state.StateDefault {
		return
	}

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
}
