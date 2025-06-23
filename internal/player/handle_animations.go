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

	normalizedVelocity := p.velocity.Normalize()

	if normalizedVelocity.Y < 0 && normalizedVelocity.X >= normalizedVelocity.Y {
		p.animationState = animation.AnimationStateWalkingUp
	}

	if normalizedVelocity.Y > 0 && normalizedVelocity.X <= normalizedVelocity.Y {
		p.animationState = animation.AnimationStateWalkingDown
	}

	if normalizedVelocity.X < 0 && normalizedVelocity.X <= normalizedVelocity.Y {
		p.animationState = animation.AnimationStateWalkingLeft
	}

	if normalizedVelocity.X > 0 && normalizedVelocity.X >= normalizedVelocity.Y {
		p.animationState = animation.AnimationStateWalkingRight
	}

	if p.velocity.IsZero() {
		p.animationState = animation.AnimationState(int(p.animationState) % 4)
	}

	if p.animationState != prevAnimationState {
		p.frameIndex = 0
	}
}
