package ui

import (
	"fmt"
	"image/color"
	"log"

	"github.com/Dobefu/topdown-adventure-game/internal/fonts"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

var (
	handleImg *widget.ButtonImage
)

func init() {
	var err error
	handleImg, err = loadButtonImage()

	if err != nil {
		log.Fatal(err)
	}
}

// NewSlider creates a new EbitenUI slider with some default options.
func NewSlider(
	label string,
	currentValue int,
	changedHandler func(args *widget.SliderChangedEventArgs),
	opts ...widget.SliderOpt,
) (*widget.Container, *widget.Slider) {
	var sliderLabel *widget.Label
	var sliderText *widget.Label

	defaultOpts := []widget.SliderOpt{
		widget.SliderOpts.Direction(widget.DirectionHorizontal),
		widget.SliderOpts.MinMax(0, 100),
		widget.SliderOpts.InitialCurrent(currentValue),
		widget.SliderOpts.Images(
			&widget.SliderTrackImage{
				Idle:  image.NewNineSliceColor(color.NRGBA{128, 128, 128, 255}),
				Hover: image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			},
			handleImg,
		),
		widget.SliderOpts.FixedHandleSize(16),
		widget.SliderOpts.TrackOffset(0),
		widget.SliderOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(150, 16),
		),
		widget.SliderOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
		widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
			changedHandler(args)

			sliderText.Label = fmt.Sprintf("%d%%", args.Current)
		}),
	}

	slider := widget.NewSlider(append(defaultOpts, opts...)...)

	sliderLabel = widget.NewLabel(
		widget.LabelOpts.Text(
			label,
			fonts.FontDefaultMd,
			&widget.LabelColor{Idle: color.White},
		),
	)

	sliderText = widget.NewLabel(
		widget.LabelOpts.TextOpts(
			widget.TextOpts.Position(
				widget.TextPositionEnd,
				widget.TextPositionCenter,
			),
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.MinSize(45, 0),
				widget.WidgetOpts.LayoutData(
					widget.RowLayoutData{
						Position: widget.RowLayoutPositionCenter,
					},
				),
			),
		),
		widget.LabelOpts.Text(
			fmt.Sprintf("%d%%", slider.Current),
			fonts.FontDefaultMd,
			&widget.LabelColor{Idle: color.White},
		),
	)

	outerContainer := NewRowContainer(
		widget.DirectionVertical,
		0,
		0,
		0,
	)

	innerContainer := NewRowContainer(widget.DirectionHorizontal, 0, 0, 0)

	outerContainer.AddChild(sliderLabel)
	outerContainer.AddChild(innerContainer)

	innerContainer.AddChild(slider)
	innerContainer.AddChild(sliderText)

	return outerContainer, slider
}
