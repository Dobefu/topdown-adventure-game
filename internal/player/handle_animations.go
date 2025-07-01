package player

import (
	"github.com/Dobefu/topdown-adventure-game/internal/animation"
	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/vectors"
)

var (
	frameCount int
)

const (
	NUM_FRAMES = 16

	RUNNING_THRESHOLD = 1.75
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

	if _, ok := p.input.PressedActionInfo(input.ActionAimAnalog); ok ||
		p.input.ActionIsPressed(input.ActionAimMouse) {

		cameraPos := *p.GetCameraPosition()
		pos := *p.GetPosition()
		pos.Add(vectors.Vector3{
			X: FRAME_WIDTH / 2,
			Y: FRAME_HEIGHT / 2,
			Z: 0,
		})

		cameraPos.Sub(pos)
		angle := cameraPos.AngleDegrees()

		// Aiming state.
		p.animationState = animation.AnimationState(
			int((angle+22.5)/45)%8 + int(animation.AnimationStateOffsetAim),
		)
	} else if p.velocity.IsZero() {
		// Idle state.
		p.animationState = animation.AnimationState(
			int(p.animationState)%8 + int(animation.AnimationStateOffsetIdle),
		)
	} else if p.velocity.Magnitude() < RUNNING_THRESHOLD {
		// Walking state.
		p.animationState = animation.AnimationState(
			int((angle+22.5)/45)%8 + int(animation.AnimationStateOffsetWalk),
		)
	} else {
		// Running state.
		p.animationState = animation.AnimationState(
			int((angle+22.5)/45)%8 + int(animation.AnimationStateOffsetRun),
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
