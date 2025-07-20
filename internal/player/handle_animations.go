package player

import (
	"github.com/Dobefu/topdown-adventure-game/internal/animation"
	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/topdown-adventure-game/internal/state"
	"github.com/Dobefu/vectors"
)

const (
	// RunningThreshold is the threshold at which the running animation plays.
	RunningThreshold = 1.75
)

func (p *Player) handleAnimations() {
	p.frameCount++
	prevAnimationState := p.animationState

	// Change the animation frame every 3 game ticks.
	if (p.frameCount % 3) == 0 {
		p.frameIndex++

		if p.frameIndex >= NumFrames {
			p.frameIndex = 0

			if p.state == state.StateHurt {
				p.state = state.StateDefault
			}
		}
	}

	angle := p.velocity.AngleDegrees()

	if p.state == state.StateHurt {
		// Hurt state.
		p.animationState = animation.State(
			int(p.animationState)%8 + int(animation.StateOffsetHurt),
		)
	} else if _, ok := p.input.PressedActionInfo(input.ActionAimAnalog); ok ||
		p.input.ActionIsPressed(input.ActionAimMouse) {

		cameraPos := *p.GetCameraPosition()
		pos := *p.GetPosition()
		pos.Add(vectors.Vector3{
			X: FrameWidth / 2,
			Y: FrameHeight / 2,
			Z: 0,
		})

		cameraPos.Sub(pos)
		angle := cameraPos.AngleDegrees()

		// Aiming state.
		p.animationState = animation.State(
			int((angle+22.5)/45)%8 + int(animation.StateOffsetAim),
		)
	} else if p.velocity.IsZero() {
		// Idle state.
		p.animationState = animation.State(
			int(p.animationState)%8 + int(animation.StateOffsetIdle),
		)
	} else if p.velocity.Magnitude() < RunningThreshold {
		// Walking state.
		p.animationState = animation.State(
			int((angle+22.5)/45)%8 + int(animation.StateOffsetWalk),
		)
	} else {
		// Running state.
		p.animationState = animation.State(
			int((angle+22.5)/45)%8 + int(animation.StateOffsetRun),
		)
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
