package ui

import (
	"github.com/ebitenui/ebitenui/widget"
)

func NewContainer(opts ...widget.ContainerOpt) *widget.Container {
	defaultOpts := []widget.ContainerOpt{
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(8),
			),
		),
	}

	return widget.NewContainer(append(defaultOpts, opts...)...)
}
