package scene

import (
	"image/color"

	"github.com/Dobefu/topdown-adventure-game/internal/fonts"
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

	sliderVolumeContainer, sliderVolume := ui.NewSlider()

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
