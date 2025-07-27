package gameobject

import (
	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/topdown-adventure-game/internal/state"
	"github.com/Dobefu/topdown-adventure-game/internal/tiledata"
	"github.com/Dobefu/vectors"
	ebitengine_input "github.com/quasilyte/ebitengine-input"
)

// MovementConfig defines the configuration for movement behaviour.
type MovementConfig struct {
	VelocityDamping  float64
	StopThreshold    float64
	Acceleration     float64
	MaxSpeed         float64
	RunningThreshold float64
}

// DefaultMovementConfig gets the default movement configuration.
func DefaultMovementConfig() MovementConfig {
	return MovementConfig{
		VelocityDamping:  .9,
		StopThreshold:    .15,
		Acceleration:     .6,
		MaxSpeed:         4,
		RunningThreshold: 2,
	}
}

// HandleMovement handles the movement logic.
func HandleMovement(
	obj interfaces.MovementHandler,
	velocity *vectors.Vector3,
	rawInputVelocity *vectors.Vector3,
	inputHandler *ebitengine_input.Handler,
	currentState *state.State,
	config MovementConfig,
) {
	if rawInputVelocity != nil {
		HandleInput(rawInputVelocity, inputHandler, currentState, config)
		velocity.Add(*rawInputVelocity)
	}

	// Dampen the velocity.
	velocity.Mul(vectors.Vector3{X: config.VelocityDamping, Y: config.VelocityDamping, Z: config.VelocityDamping})

	// If the velocity magnitude is very low, set it to zero.
	// This allows the idle animations to work.
	if velocity.Magnitude() < config.StopThreshold {
		velocity.Mul(vectors.Vector3{X: 0, Y: 0, Z: 1})
	}

	pos := *obj.GetPosition()

	// Apply gravity.
	if pos.Z > 0 {
		velocity.Z -= config.Acceleration / 2
	}

	handleMovementState(currentState, velocity, pos, inputHandler, config)

	x1, y1, x2, y2 := obj.GetCollisionRect()
	newVelocity, _, collidedTiles := obj.MoveWithCollisionRect(*velocity, x1, y1, x2, y2)
	*velocity = newVelocity

	// Handle ledge jumping if state is provided.
	handleLedgeJump(currentState, collidedTiles, velocity)
}

// HandleInput handles movement input.
func HandleInput(
	rawInputVelocity *vectors.Vector3,
	inputHandler *ebitengine_input.Handler,
	currentState *state.State,
	config MovementConfig,
) {
	if inputHandler == nil || rawInputVelocity == nil {
		return
	}

	rawInputVelocity.Clear()

	if currentState != nil && *currentState == state.StateDefault {
		if info, ok := inputHandler.PressedActionInfo(input.ActionMoveAnalog); ok {
			rawInputVelocity.Add(vectors.Vector3{
				X: info.Pos.X,
				Y: info.Pos.Y,
				Z: 0,
			})
		} else {
			if inputHandler.ActionIsPressed(input.ActionMoveLeft) {
				rawInputVelocity.X--
			}

			if inputHandler.ActionIsPressed(input.ActionMoveRight) {
				rawInputVelocity.X++
			}

			if inputHandler.ActionIsPressed(input.ActionMoveUp) {
				rawInputVelocity.Y--
			}

			if inputHandler.ActionIsPressed(input.ActionMoveDown) {
				rawInputVelocity.Y++
			}
		}
	}

	if rawInputVelocity.Magnitude() > 0 {
		rawInputVelocity.Normalize()
		rawInputVelocity.Mul(vectors.Vector3{X: config.Acceleration, Y: config.Acceleration, Z: 1})
	}
}

func handleMovementState(
	currentState *state.State,
	velocity *vectors.Vector3,
	pos vectors.Vector3,
	inputHandler *ebitengine_input.Handler,
	config MovementConfig,
) {
	if currentState == nil {
		velocity.ClampMagnitude(config.MaxSpeed)
		return
	}

	if *currentState == state.StateJump && pos.Z <= 0 {
		*currentState = state.StateDefault
	}

	if *currentState == state.StateDefault {
		if inputHandler != nil {
			if _, ok := inputHandler.PressedActionInfo(input.ActionAimAnalog); ok ||
				inputHandler.ActionIsPressed(input.ActionAimMouse) {

				velocity.ClampMagnitude(config.RunningThreshold)
			} else {
				velocity.ClampMagnitude(config.MaxSpeed)
			}
		} else {
			velocity.ClampMagnitude(config.MaxSpeed)
		}
	}
}

func handleLedgeJump(
	currentState *state.State,
	collidedTiles []int,
	velocity *vectors.Vector3,
) {
	if currentState == nil {
		return
	}

	for _, collidedTile := range collidedTiles {
		if collidedTile == tiledata.TileCollisionLedgeVertical &&
			*currentState != state.StateJump &&
			velocity.Y != 0 {
			*currentState = state.StateJump

			velocity.X = 0
			velocity.Normalize()

			velocity.Y *= 8
			velocity.Z = 4
			break
		}

		if collidedTile == tiledata.TileCollisionLedgeHorizontal &&
			*currentState != state.StateJump &&
			velocity.X != 0 {
			*currentState = state.StateJump

			velocity.Y = 0
			velocity.Normalize()

			velocity.X *= 8
			velocity.Z = 4
			break
		}
	}
}
