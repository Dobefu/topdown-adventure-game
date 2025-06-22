package ui

import (
	"image/color"
	"log"

	"github.com/Dobefu/topdown-adventure-game/internal/fonts"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

var (
	buttonImage *widget.ButtonImage
)

func init() {
	var err error
	buttonImage, err = loadButtonImage()

	if err != nil {
		log.Fatal(err)
	}
}

func NewButton(opts ...widget.ButtonOpt) *widget.Button {
	defaultOpts := []widget.ButtonOpt{
		widget.ButtonOpts.TextPadding(widget.Insets{
			Top:    16,
			Left:   24,
			Right:  24,
			Bottom: 16,
		}),
		widget.ButtonOpts.TextFace(fonts.FontDefaultMd),
		widget.ButtonOpts.TextColor(&widget.ButtonTextColor{
			Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
		}),
		widget.ButtonOpts.Image(buttonImage),
	}

	return widget.NewButton(append(defaultOpts, opts...)...)
}

func loadButtonImage() (*widget.ButtonImage, error) {
	idle := image.NewBorderedNineSliceColor(
		color.NRGBA{R: 170, G: 170, B: 180, A: 255},
		color.NRGBA{90, 90, 90, 255},
		3,
	)

	hover := image.NewBorderedNineSliceColor(
		color.NRGBA{R: 130, G: 130, B: 150, A: 255},
		color.NRGBA{70, 70, 70, 255},
		3,
	)

	pressed := image.NewAdvancedNineSliceColor(
		color.NRGBA{R: 130, G: 130, B: 150, A: 255},
		image.NewBorder(3, 2, 2, 2, color.NRGBA{70, 70, 70, 255}),
	)

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}
