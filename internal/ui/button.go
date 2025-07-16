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

func NewButton(text string, opts ...widget.ButtonOpt) *widget.Button {
	defaultOpts := []widget.ButtonOpt{
		widget.ButtonOpts.TextLabel(text),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Top:    8,
			Left:   24,
			Right:  24,
			Bottom: 8,
		}),
		widget.ButtonOpts.TextFace(fonts.FontDefaultSm),
		widget.ButtonOpts.TextColor(&widget.ButtonTextColor{
			Idle: color.NRGBA{0xff, 0xff, 0xff, 0xff},
		}),
		widget.ButtonOpts.Image(buttonImg),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),
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

	return &widget.ButtonImage{
		Idle:    imageUI.NewNineSliceSimple(ebiten.NewImageFromImage(imgIdle), 12, 40),
		Hover:   imageUI.NewNineSliceSimple(ebiten.NewImageFromImage(imgHover), 12, 40),
		Pressed: imageUI.NewNineSliceSimple(ebiten.NewImageFromImage(imgPressed), 12, 40),
	}, nil
}
