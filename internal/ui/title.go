package ui

import (
	"image/color"

	"github.com/Dobefu/topdown-adventure-game/internal/fonts"
	"github.com/ebitenui/ebitenui/widget"
)

// NewTitle creates a new EbitenUI text widget with styling for a title.
func NewTitle(text string) *widget.Text {
	return widget.NewText(
		widget.TextOpts.Text(text, fonts.FontDefaultXl, color.White),

		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)
}
