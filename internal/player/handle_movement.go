package player

import (
	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/topdown-adventure-game/internal/state"
	"github.com/Dobefu/topdown-adventure-game/internal/tiledata"
	"github.com/Dobefu/vectors"
)

const (
	// VelocityDamping defines the amount of friction that the player has.
	VelocityDamping float64 = .9
	// StopThreshold defines the minimum velocity at which the player is
	// considered "stopped".
	StopThreshold float64 = .15
	// Acceleration defines the acceleration of the velocity.
	Acceleration float64 = .6
	// MaxSpeed defines the maximum speed that the enemy can have.
	MaxSpeed float64 = 4
)

func (p *Player) handleMovement() {
	// Dampen the velocity.
	p.velocity.Mul(vectors.Vector3{X: VelocityDamping, Y: VelocityDamping, Z: VelocityDamping})

	// If the velocity magnitude is very low, set it to zero.
	// This allows the idle animations to work.
	if p.velocity.Magnitude() < StopThreshold {
		p.velocity.Mul(vectors.Vector3{X: 0, Y: 0, Z: 1})
	}

	p.rawInputVelocity.Clear()

	p.handleInput()

	if p.rawInputVelocity.Magnitude() > 0 {
		p.rawInputVelocity.Normalize()
		p.rawInputVelocity.Mul(vectors.Vector3{X: Acceleration, Y: Acceleration, Z: 1})
	}

	p.velocity.Add(p.rawInputVelocity)

	pos := *p.GetPosition()

	// Apply gravity.
	if pos.Z > 0 {
		p.velocity.Z -= Acceleration / 2
	}

	// Reset the jump state.
	if p.state == state.StateJump && pos.Z <= 0 {
		p.state = state.StateDefault
	}

	if p.state != state.StateHurt {
		if _, ok := p.input.PressedActionInfo(input.ActionAimAnalog); ok ||
			p.input.ActionIsPressed(input.ActionAimMouse) {

			p.velocity.ClampMagnitude(RunningThreshold)
		} else if p.state == state.StateDefault {
			p.velocity.ClampMagnitude(MaxSpeed)
		}
	}

	x1, y1, x2, y2 := p.GetCollisionRect()
	newVelocity, _, collidedTile := p.MoveWithCollisionRect(p.velocity, x1, y1, x2, y2)
	p.velocity = newVelocity

	if collidedTile == tiledata.TileCollisionLedgeVertical &&
		p.state != state.StateJump &&
		p.velocity.Y != 0 {
		p.state = state.StateJump

		p.velocity.X = 0
		p.velocity.Normalize()

		p.velocity.Y *= 8
		p.velocity.Z = 4
	}

	if collidedTile == tiledata.TileCollisionLedgeHorizontal &&
		p.state != state.StateJump &&
		p.velocity.X != 0 {
		p.state = state.StateJump

		p.velocity.Y = 0
		p.velocity.Normalize()

		p.velocity.X *= 8
		p.velocity.Z = 4
	}
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
			p.rawInputVelocity.X--
		}

		if p.input.ActionIsPressed(input.ActionMoveRight) {
			p.rawInputVelocity.X++
		}

		if p.input.ActionIsPressed(input.ActionMoveUp) {
			p.rawInputVelocity.Y--
		}

		if p.input.ActionIsPressed(input.ActionMoveDown) {
			p.rawInputVelocity.Y++
		}
	}
}
