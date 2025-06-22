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
)

func init() {
	Input = input.System{}
	Input.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})

	Keymap = input.Keymap{
		ActionMoveLeft:  {input.KeyGamepadLeft, input.KeyGamepadLStickLeft, input.KeyLeft, input.KeyA, input.KeyH},
		ActionMoveRight: {input.KeyGamepadRight, input.KeyGamepadLStickRight, input.KeyRight, input.KeyD, input.KeyL},
		ActionMoveUp:    {input.KeyGamepadUp, input.KeyGamepadLStickUp, input.KeyUp, input.KeyW, input.KeyK},
		ActionMoveDown:  {input.KeyGamepadDown, input.KeyGamepadLStickDown, input.KeyDown, input.KeyS, input.KeyJ},
	}
}
