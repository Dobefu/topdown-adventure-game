package ui

import (
	"github.com/ebitenui/ebitenui/widget"
)

// NewRowContainer creates a new EbitenUI container with a row layout.
func NewRowContainer(
	direction widget.Direction,
	spacing int,
	paddingBlock int,
	paddingInline int,
	opts ...widget.ContainerOpt,
) *widget.Container {
	defaultOpts := []widget.ContainerOpt{
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(direction),
				widget.RowLayoutOpts.Spacing(spacing),
				widget.RowLayoutOpts.Padding(widget.Insets{
					Top:    paddingBlock,
					Left:   paddingInline,
					Right:  paddingInline,
					Bottom: paddingBlock,
				}),
			),
		),
	}

	return NewContainer(append(defaultOpts, opts...)...)
}
