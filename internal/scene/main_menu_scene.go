package scene

import (
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

	s.ui.Container.AddChild(ui.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		widget.ButtonOpts.TextLabel("Start"),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			s.Game.SetScene(&Level1Scene{})
		}),
	))
}
