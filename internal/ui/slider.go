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

func NewSlider(
	currentValue int,
	changedHandler func(args *widget.SliderChangedEventArgs),
	opts ...widget.SliderOpt,
) (*widget.Container, *widget.Slider) {
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
		widget.SliderOpts.FixedHandleSize(32),
		widget.SliderOpts.TrackOffset(0),
		widget.SliderOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(200, 32),
		),
		widget.SliderOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
		widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
			changedHandler(args)

			sliderText.Label = fmt.Sprintf("% 4d%%", args.Current)
		}),
	}

	slider := widget.NewSlider(append(defaultOpts, opts...)...)

	sliderText = widget.NewLabel(
		widget.LabelOpts.TextOpts(widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionCenter,
		}))),
		widget.LabelOpts.Text(
			fmt.Sprintf("% 4d%%", slider.Current),
			fonts.FontDefaultSm,
			&widget.LabelColor{Idle: color.White},
		),
	)

	container := NewContainer(
		16,
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			),
		),
	)

	container.AddChild(slider)
	container.AddChild(sliderText)

	return container, slider
}
