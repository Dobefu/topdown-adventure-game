package scene

import (
	"fmt"
	"log/slog"

	"github.com/Dobefu/topdown-adventure-game/internal/storage"
	"github.com/Dobefu/topdown-adventure-game/internal/ui"
	"github.com/ebitenui/ebitenui/widget"
)

// OptionsScene defines a main menu scene instance.
type OptionsScene struct {
	Scene
}

// Init initializes the options scene.
func (s *OptionsScene) Init() {
	s.Scene.Init()
}

// InitUI initializes the options scene UI.
func (s *OptionsScene) InitUI() {
	s.Scene.InitUI()

	currentVolume, err := storage.GetOption("volume", 100)

	if err != nil {
		slog.Error(err.Error())
	}

	sliderVolumeContainer, sliderVolume := ui.NewSlider(
		"Volume",
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
		"Back",
		widget.ButtonOpts.ClickedHandler(func(_ *widget.ButtonClickedEventArgs) {
			s.GetGame().SetScene(&MainMenuScene{})
		}),
	)

	outerContainer := ui.NewRowContainer(widget.DirectionVertical, 16, 0, 0)

	outerContainer.AddChild(ui.NewTitle("Options"))

	innerContainer := ui.NewRowContainer(
		widget.DirectionVertical,
		4,
		0,
		0,
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
