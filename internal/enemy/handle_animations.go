package enemy

import (
	"github.com/Dobefu/topdown-adventure-game/internal/animation"
	"github.com/Dobefu/topdown-adventure-game/internal/state"
)

const (
	// RunningThreshold is the threshold at which the running animation plays.
	RunningThreshold = 1.75
)

func (e *Enemy) handleAnimations() {
	e.frameCount++
	prevAnimationState := e.animationState

	// Change the animation frame every 3 game ticks.
	if (e.frameCount % 3) == 0 {
		e.frameIndex++

		if e.frameIndex >= NumFrames {
			e.frameIndex = 0

			if e.state == state.StateHurt {
				e.state = state.StateDefault
			}
		}
	}

	angle := e.velocity.AngleDegrees()

	if e.state == state.StateHurt {
		// Hurt state.
		e.animationState = animation.State(
			int(e.animationState)%8 + int(animation.StateOffsetHurt),
		)
	} else if e.state == state.StateJump {
		// Jump state.
		e.animationState = animation.State(
			int(e.animationState)%8 + int(animation.StateOffsetJump),
		)
	} else if e.velocity.IsZero() {
		// Idle state.
		e.animationState = animation.State(
			int(e.animationState)%8 + int(animation.StateOffsetIdle),
		)
	} else if e.velocity.Magnitude() < RunningThreshold {
		// Walking state.
		e.animationState = animation.State(
			int((angle+22.5)/45)%8 + int(animation.StateOffsetWalk),
		)
	} else {
		// Running state.
		e.animationState = animation.State(
			int((angle+22.5)/45)%8 + int(animation.StateOffsetRun),
		)
	}

	prevCategory := int(prevAnimationState) / 8
	currentCategory := int(e.animationState) / 8

	// Only reset the animation index when the animation category changes.
	// This prevents the animation from resetting when just the direction changes
	// within the same category.
	if prevCategory != currentCategory {
		e.frameIndex = 0
	}
}
