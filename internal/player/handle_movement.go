package player

import (
	"github.com/Dobefu/topdown-adventure-game/internal/input"
)

func (p *Player) handleMovement() {
	// Clear the velocity at the start of every update tick.
	// This ensures that we can use the velocity for animation states later.
	p.velocity.Clear()

	if p.input.ActionIsPressed(input.ActionMoveLeft) {
		p.velocity.X = -2
	}

	if p.input.ActionIsPressed(input.ActionMoveRight) {
		p.velocity.X = 2
	}

	if p.input.ActionIsPressed(input.ActionMoveUp) {
		p.velocity.Y = -2
	}

	if p.input.ActionIsPressed(input.ActionMoveDown) {
		p.velocity.Y = 2
	}

	pos := p.GetPosition()
	pos.Add(p.velocity)
}
