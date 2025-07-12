package scene

import (
	"image/color"

	"github.com/Dobefu/topdown-adventure-game/internal/fonts"
	"github.com/Dobefu/topdown-adventure-game/internal/ui"
	"github.com/ebitenui/ebitenui/widget"
)

type MainMenuScene struct {
	Scene
}

func (s *MainMenuScene) Init() {
	s.Scene.Init()
}

func (s *MainMenuScene) InitUI() {
	s.Scene.InitUI()

	btnStart := ui.NewButton(
		widget.ButtonOpts.TextLabel("Start"),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			s.GetGame().SetScene(&OverworldScene{})
		}),

		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),
	)

	btnOptions := ui.NewButton(
		widget.ButtonOpts.TextLabel("Options"),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			s.GetGame().SetScene(&OptionsScene{})
		}),

		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),
	)

	container := ui.NewContainer()

	container.AddChild(widget.NewText(
		widget.TextOpts.Text("Title\n\n", fonts.FontDefaultXxl, color.White),

		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	))

	btnContainer := ui.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)
	container.AddChild(btnContainer)

	btnContainer.AddChild(btnStart)
	btnContainer.AddChild(btnOptions)

	btnStart.AddFocus(widget.FOCUS_SOUTH, btnOptions)
	btnOptions.AddFocus(widget.FOCUS_NORTH, btnStart)

	s.ui.Container.AddChild(container)
}
