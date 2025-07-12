package game

import (
	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/ebitenui/ebitenui/widget"
)

func (g *game) UpdateUIInput() {
	ui := g.scene.GetUI()

	if g.input.ActionIsPressed(input.ActionMoveLeft) {
		ui.ChangeFocus(widget.FOCUS_WEST)
	}

	if g.input.ActionIsPressed(input.ActionMoveRight) {
		ui.ChangeFocus(widget.FOCUS_EAST)
	}

	if g.input.ActionIsPressed(input.ActionMoveUp) {
		ui.ChangeFocus(widget.FOCUS_NORTH)
	}

	if g.input.ActionIsPressed(input.ActionMoveDown) {
		ui.ChangeFocus(widget.FOCUS_SOUTH)
	}

	if g.input.ActionIsPressed(input.ActionClick) {
		focusedWidget := ui.GetFocusedWidget()

		if focusedWidget != nil {
			focusedWidget.(*widget.Button).Click()
		}
	}
}
