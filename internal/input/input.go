package input

import (
	input "github.com/quasilyte/ebitengine-input"
)

var (
	Input  input.System
	Keymap input.Keymap
)

const (
	ActionMoveLeft = iota
	ActionMoveRight
	ActionMoveUp
	ActionMoveDown
	ActionMoveAnalog

	ActionAimMouse
	ActionAimAnalog
	ActionShoot

	ActionToggleDebug
)

func init() {
	Input = input.System{}
	Input.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})

	Keymap = input.Keymap{
		ActionMoveLeft:   {input.KeyGamepadLeft, input.KeyLeft, input.KeyA, input.KeyH},
		ActionMoveRight:  {input.KeyGamepadRight, input.KeyRight, input.KeyD, input.KeyL},
		ActionMoveUp:     {input.KeyGamepadUp, input.KeyUp, input.KeyW, input.KeyK},
		ActionMoveDown:   {input.KeyGamepadDown, input.KeyDown, input.KeyS, input.KeyJ},
		ActionMoveAnalog: {input.KeyGamepadLStickMotion},

		ActionAimMouse:  {input.KeyMouseRight},
		ActionAimAnalog: {input.KeyGamepadRStickMotion},
		ActionShoot:     {input.KeyMouseLeft, input.KeySpace, input.KeyGamepadR2},

		ActionToggleDebug: {input.KeyF5},
	}
}
