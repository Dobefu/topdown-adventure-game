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

	btnBack := ui.NewButton(
		widget.ButtonOpts.TextLabel("Back"),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			s.GetGame().SetScene(&MainMenuScene{})
		}),

		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)

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
		widget.TextOpts.Text("Options\n\n", fonts.FontDefaultXxl, color.White),

		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	))

	container.AddChild(btnBack)

	s.ui.Container.AddChild(container)
}
