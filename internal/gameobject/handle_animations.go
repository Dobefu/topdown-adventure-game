package gameobject

import (
	"github.com/Dobefu/topdown-adventure-game/internal/animation"
	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/topdown-adventure-game/internal/state"
)

const (
	// RunningThreshold is the threshold at which the running animation plays.
	RunningThreshold = 1.75
)

// HandleAnimations handles the animations of a gameobject.
func HandleAnimations(
	obj *GameObject,
) {
	obj.GameFrameCount++
	prevAnimationState := obj.AnimationState

	// Change the animation frame every 3 game ticks.
	if (obj.GameFrameCount % 3) == 0 {
		obj.FrameIndex++

		if obj.FrameIndex >= obj.NumFrames {
			obj.FrameIndex = 0

			if obj.State == state.StateHurt {
				obj.State = state.StateDefault
			}
		}
	}

	angle := obj.Velocity.AngleDegrees()

	if obj.State == state.StateHurt {
		// Hurt state.
		obj.AnimationState = animation.State(
			int(obj.AnimationState)%8 + int(animation.StateOffsetHurt),
		)
	} else if obj.State == state.StateJump {
		// Jump state.
		obj.AnimationState = animation.State(
			int(obj.AnimationState)%8 + int(animation.StateOffsetJump),
		)
	} else if obj.Velocity.IsZero() {
		// Idle state.
		obj.AnimationState = animation.State(
			int(obj.AnimationState)%8 + int(animation.StateOffsetIdle),
		)
	} else if obj.Velocity.Magnitude() < RunningThreshold {
		// Walking state.
		obj.AnimationState = animation.State(
			int((angle+22.5)/45)%8 + int(animation.StateOffsetWalk),
		)
	} else {
		// Running state.
		obj.AnimationState = animation.State(
			int((angle+22.5)/45)%8 + int(animation.StateOffsetRun),
		)
	}

	if obj.Input != nil {
		if _, ok := obj.Input.PressedActionInfo(input.ActionAimAnalog); ok ||
			obj.Input.ActionIsPressed(input.ActionAimMouse) {

			cameraPos := *obj.GetCameraPosition()
			// obj.Position.Add(vectors.Vector3{
			// 	X: float64(obj.FrameWidth) / 2,
			// 	Y: float64(obj.FrameHeight) / 2,
			// 	Z: 0,
			// })

			cameraPos.Sub(obj.Position)
			angle := cameraPos.AngleDegrees()

			// Aiming state.
			obj.AnimationState = animation.State(
				int((angle+22.5)/45)%8 + int(animation.StateOffsetAim),
			)
		}
	}

	prevCategory := int(prevAnimationState) / 8
	currentCategory := int(obj.AnimationState) / 8

	// Only reset the animation index when the animation category changes.
	// This prevents the animation from resetting when just the direction changes
	// within the same category.
	if prevCategory != currentCategory {
		obj.FrameIndex = 0
	}
}
