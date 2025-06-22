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

	container := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			),
		),
	)

	container.AddChild(widget.NewText(
		widget.TextOpts.Text("Title\n\n", fonts.FontDefaultXxl, color.White),

		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	))

	container.AddChild(ui.NewButton(
		widget.ButtonOpts.TextLabel("Start"),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			s.Game.SetScene(&Level1Scene{})
		}),

		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	))

	s.ui.Container.AddChild(container)
}
