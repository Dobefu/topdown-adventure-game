package ui

import (
	"github.com/ebitenui/ebitenui/widget"
)

// NewAnchorContainer creates a new EbitenUI container with an anchor layout.
func NewAnchorContainer(
	paddingBlock int,
	paddingInline int,
	opts ...widget.ContainerOpt,
) *widget.Container {
	defaultOpts := []widget.ContainerOpt{
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(
				widget.AnchorLayoutOpts.Padding(widget.Insets{
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
