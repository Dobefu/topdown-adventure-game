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
	} else if p.velocity.Magnitude() >= 1.75 {
		// Running state.
		p.animationState = animation.AnimationState(int((angle+22.5)/45)%8 + 16)
	} else {
		// Walking state.
		p.animationState = animation.AnimationState(int((angle+22.5)/45)%8 + 8)
	}

	prevCategory := int(prevAnimationState) / 8
	currentCategory := int(p.animationState) / 8

	// Only reset the animation index when the animation category changes.
	// This prevents the animation from resetting when just the direction changes
	// within the same category.
	if prevCategory != currentCategory {
		p.frameIndex = 0
	}
}
