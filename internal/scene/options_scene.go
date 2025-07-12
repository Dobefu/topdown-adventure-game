package scene

import (
	"fmt"
	"image/color"
	"log/slog"
	"strconv"

	"github.com/Dobefu/topdown-adventure-game/internal/fonts"
	"github.com/Dobefu/topdown-adventure-game/internal/storage"
	"github.com/Dobefu/topdown-adventure-game/internal/ui"
	"github.com/ebitenui/ebitenui/widget"
)

type OptionsScene struct {
	Scene
}

func (s *OptionsScene) Init() {
	s.Scene.Init()
}

func (s *OptionsScene) InitUI() {
	s.Scene.InitUI()

	currentVolume, err := getCurrentVolume()

	if err != nil {
		slog.Error(err.Error())
	}

	sliderVolumeContainer, sliderVolume := ui.NewSlider(
		currentVolume,
		func(args *widget.SliderChangedEventArgs) {
			if args.Dragging {
				return
			}

			err := storage.SetOption("volume", fmt.Sprintf("%d", args.Current))

			if err != nil {
				slog.Error(err.Error())
			}
		},
	)

	btnBack := ui.NewButton(
		widget.ButtonOpts.TextLabel("Back"),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			s.GetGame().SetScene(&MainMenuScene{})
		}),
	)

	outerContainer := ui.NewContainer(64)

	outerContainer.AddChild(widget.NewText(
		widget.TextOpts.Text("Options", fonts.FontDefaultXxl, color.White),

		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	))

	innerContainer := ui.NewContainer(
		16,
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)
	outerContainer.AddChild(innerContainer)

	innerContainer.AddChild(sliderVolumeContainer)
	innerContainer.AddChild(btnBack)

	sliderVolume.AddFocus(widget.FOCUS_SOUTH, btnBack)
	btnBack.AddFocus(widget.FOCUS_NORTH, sliderVolume)

	s.ui.Container.AddChild(outerContainer)
}

func getCurrentVolume() (currentVolume int, err error) {
	volumeOption, err := storage.GetOption("volume")

	if err != nil {
		return 100, err
	}

	parsedVolumeOption, err := strconv.ParseInt(volumeOption, 10, 64)

	if err != nil {
		return 100, err
	}

	currentVolume = int(parsedVolumeOption)

	return currentVolume, nil
}
