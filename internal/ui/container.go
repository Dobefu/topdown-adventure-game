package ui

import (
	"github.com/ebitenui/ebitenui/widget"
)

// NewContainer creates a new EbitenUI container with some default options.
func NewContainer(
	opts ...widget.ContainerOpt,
) *widget.Container {
	defaultOpts := []widget.ContainerOpt{
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
	}

	return widget.NewContainer(append(defaultOpts, opts...)...)
}
