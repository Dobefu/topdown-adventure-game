package ui

import (
	"bytes"
	"image"
	"image/color"
	"log"

	"github.com/Dobefu/topdown-adventure-game/internal/fonts"
	imageUI "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	_ "embed"
)

var (
	//go:embed img/button-idle.png
	buttonIdleImgBytes []byte

	//go:embed img/button-hover.png
	buttonHoverImgBytes []byte

	//go:embed img/button-pressed.png
	buttonPressedImgBytes []byte

	buttonImg *widget.ButtonImage
)

func init() {
	var err error
	buttonImg, err = loadButtonImage()

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
			Idle: color.NRGBA{0xff, 0xff, 0xff, 0xff},
		}),
		widget.ButtonOpts.Image(buttonImg),
	}

	return widget.NewButton(append(defaultOpts, opts...)...)
}

func loadButtonImage() (*widget.ButtonImage, error) {
	imgIdle, _, err := image.Decode(bytes.NewReader(buttonIdleImgBytes))
	if err != nil {
		return nil, err
	}

	imgHover, _, err := image.Decode(bytes.NewReader(buttonHoverImgBytes))
	if err != nil {
		return nil, err
	}

	imgPressed, _, err := image.Decode(bytes.NewReader(buttonPressedImgBytes))
	if err != nil {
		return nil, err
	}

	idle := imageUI.NewNineSliceSimple(ebiten.NewImageFromImage(imgIdle), 12, 40)
	hover := imageUI.NewNineSliceSimple(ebiten.NewImageFromImage(imgHover), 12, 40)
	pressed := imageUI.NewNineSliceSimple(ebiten.NewImageFromImage(imgPressed), 12, 40)

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}
