package player

import (
	"github.com/Dobefu/topdown-adventure-game/internal/animation"
	"github.com/Dobefu/topdown-adventure-game/internal/input"
)

func (p *Player) handleMovement() {
	p.velocity.Clear()

	if p.input.ActionIsPressed(input.ActionMoveLeft) {
		p.animationState = animation.AnimationStateWalkingLeft

		p.velocity.X = -2
	}

	if p.input.ActionIsPressed(input.ActionMoveRight) {
		p.animationState = animation.AnimationStateWalkingRight

		p.velocity.X = 2
	}

	if p.input.ActionIsPressed(input.ActionMoveUp) {
		p.animationState = animation.AnimationStateWalkingUp

		p.velocity.Y = -2
	}

	if p.input.ActionIsPressed(input.ActionMoveDown) {
		p.animationState = animation.AnimationStateWalkingDown

		p.velocity.Y = 2
	}

	pos := p.GetPosition()
	pos.Add(p.velocity)
}
