// Package input provides mapping for input actions.
package input

import (
	input "github.com/quasilyte/ebitengine-input"
)

var (
	// Input is a singleton that handles the "ebitengine-input" input system.
	Input input.System
	// PlayerKeymap houses the keymap for player controls.
	PlayerKeymap input.Keymap
	// UIKeymap houses the keymap for UI controls.
	UIKeymap input.Keymap
)

const (
	// ActionMoveLeft handles movement to the left.
	ActionMoveLeft = iota
	// ActionMoveRight handles movement to the right.
	ActionMoveRight
	// ActionMoveUp handles upward movement.
	ActionMoveUp
	// ActionMoveDown handles downward movement.
	ActionMoveDown
	// ActionMoveAnalog handles movement with the analog stick.
	ActionMoveAnalog

	// ActionAimMouse handles aiming with the mouse.
	ActionAimMouse
	// ActionAimAnalog handles aiming with the analog stick.
	ActionAimAnalog
	// ActionShoot handles shooting.
	ActionShoot

	// ActionConfirm handles the UI confirmation button.
	ActionConfirm
	// ActionPause handles the pause button.
	ActionPause

	// ActionToggleDebug handles toggling of the debug overlay.
	ActionToggleDebug
)

func init() {
	Input = input.System{}
	Input.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})

	PlayerKeymap = input.Keymap{
		ActionMoveLeft:   {input.KeyGamepadLeft, input.KeyLeft, input.KeyA, input.KeyH},
		ActionMoveRight:  {input.KeyGamepadRight, input.KeyRight, input.KeyD, input.KeyL},
		ActionMoveUp:     {input.KeyGamepadUp, input.KeyUp, input.KeyW, input.KeyK},
		ActionMoveDown:   {input.KeyGamepadDown, input.KeyDown, input.KeyS, input.KeyJ},
		ActionMoveAnalog: {input.KeyGamepadLStickMotion},

		ActionAimMouse:  {input.KeyMouseRight},
		ActionAimAnalog: {input.KeyGamepadRStickMotion},
		ActionShoot:     {input.KeyMouseLeft, input.KeySpace, input.KeyGamepadR2},
	}

	UIKeymap = input.Keymap{
		ActionMoveLeft:  {input.KeyGamepadLeft, input.KeyGamepadLStickLeft, input.KeyLeft, input.KeyA, input.KeyH},
		ActionMoveRight: {input.KeyGamepadRight, input.KeyGamepadLStickRight, input.KeyRight, input.KeyD, input.KeyL},
		ActionMoveUp:    {input.KeyGamepadUp, input.KeyGamepadLStickUp, input.KeyUp, input.KeyW, input.KeyK},
		ActionMoveDown:  {input.KeyGamepadDown, input.KeyGamepadLStickDown, input.KeyDown, input.KeyS, input.KeyJ},

		ActionConfirm: {input.KeyGamepadA, input.KeyEnter},
		ActionPause:   {input.KeyGamepadStart, input.KeyEscape},

		ActionToggleDebug: {input.KeyF5},
	}
}
