package player

import (
	"github.com/Dobefu/topdown-adventure-game/internal/animation"
)

func (p *Player) handleAnimations() {
	frameCount += 1
	prevAnimationState := p.animationState

	// Change the animation frame every 5 game ticks.
	if (frameCount % 5) == 0 {
		p.frameIndex += 1

		if p.frameIndex >= numFrames {
			p.frameIndex = 0
		}
	}

	// TODO: Implement a vector method to check this.
	if p.velocity.X == 0 && p.velocity.Y == 0 {
		p.animationState = animation.AnimationState(int(p.animationState) % 4)
	}

	if p.animationState != prevAnimationState {
		p.frameIndex = 0
	}
}
