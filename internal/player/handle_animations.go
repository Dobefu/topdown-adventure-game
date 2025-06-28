package player

import (
	"github.com/Dobefu/topdown-adventure-game/internal/animation"
)

var (
	frameCount int
)

const (
	NUM_FRAMES = 16
)

func (p *Player) handleAnimations() {
	frameCount += 1
	prevAnimationState := p.animationState

	// Change the animation frame every 4 game ticks.
	if (frameCount % 4) == 0 {
		p.frameIndex += 1

		if p.frameIndex >= NUM_FRAMES {
			p.frameIndex = 0
		}
	}

	angle := p.velocity.AngleDegrees()

	if p.velocity.IsZero() {
		// Idle state.
		p.animationState = animation.AnimationState(int(p.animationState) % 8)
	} else if p.velocity.Magnitude() >= 2 {
		// Running state.
		p.animationState = animation.AnimationState(int((angle+22.5)/45)%8 + 16)
	} else {
		// Walking state.
		p.animationState = animation.AnimationState(int((angle+22.5)/45)%8 + 8)
	}

	// After changing the animation state, reset the animation index.
	if p.animationState != prevAnimationState {
		p.frameIndex = 0
	}
}
