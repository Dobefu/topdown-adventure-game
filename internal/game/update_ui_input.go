package game

import (
	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/ebitenui/ebitenui/widget"
)

func (g *game) UpdateUIInput() {
	ui := g.scene.GetUI()

	if g.input.ActionIsJustPressed(input.ActionMoveLeft) {
		ui.ChangeFocus(widget.FOCUS_WEST)
	}

	if g.input.ActionIsJustPressed(input.ActionMoveRight) {
		ui.ChangeFocus(widget.FOCUS_EAST)
	}

	if g.input.ActionIsJustPressed(input.ActionMoveUp) {
		ui.ChangeFocus(widget.FOCUS_NORTH)
	}

	if g.input.ActionIsJustPressed(input.ActionMoveDown) {
		ui.ChangeFocus(widget.FOCUS_SOUTH)
	}

	if g.input.ActionIsJustPressed(input.ActionClick) {
		focusedWidget := ui.GetFocusedWidget()

		if focusedBtn, ok := focusedWidget.(*widget.Button); ok {
			focusedBtn.Click()
		}
	}
}
