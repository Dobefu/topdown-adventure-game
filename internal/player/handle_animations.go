package player

import (
	"github.com/Dobefu/topdown-adventure-game/internal/animation"
)

var (
	frameCount int
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

	// Reset to the idle state of the current direction when not moving.
	if p.velocity.IsZero() {
		p.animationState = animation.AnimationState(int(p.animationState) % 8)
	} else {
		angle := p.velocity.AngleDegrees()
		p.animationState = animation.AnimationState(int((angle+22.5)/45)%8 + 8)
	}

	// After changing the animation state, reset the animation index.
	if p.animationState != prevAnimationState {
		p.frameIndex = 0
	}
}
