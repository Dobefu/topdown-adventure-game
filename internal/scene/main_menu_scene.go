package scene

import (
	"os"
	"runtime"

	"github.com/Dobefu/topdown-adventure-game/internal/ui"
	"github.com/ebitenui/ebitenui/widget"
)

// MainMenuScene defines a main menu scene instance.
type MainMenuScene struct {
	Scene
}

// Init initializes the main menu scene.
func (s *MainMenuScene) Init() {
	s.Scene.Init()
}

// InitUI initializes the main menu scene UI.
func (s *MainMenuScene) InitUI() {
	s.Scene.InitUI()

	btnStart := ui.NewButton(
		"Start",
		widget.ButtonOpts.ClickedHandler(func(_ *widget.ButtonClickedEventArgs) {
			s.GetGame().SetScene(&OverworldScene{})
		}),
	)

	btnOptions := ui.NewButton(
		"Options",
		widget.ButtonOpts.ClickedHandler(func(_ *widget.ButtonClickedEventArgs) {
			s.GetGame().SetScene(&OptionsScene{})
		}),
	)

	outerContainer := ui.NewRowContainer(widget.DirectionVertical, 16, 0, 0)

	outerContainer.AddChild(ui.NewTitle("Title"))

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
	innerContainer.AddChild(btnStart)
	innerContainer.AddChild(btnOptions)

	if runtime.GOOS != "js" {
		innerContainer.AddChild(
			ui.NewButton(
				"Quit",
				widget.ButtonOpts.ClickedHandler(func(_ *widget.ButtonClickedEventArgs) {
					os.Exit(0)
				}),
			),
		)
	}

	btnStart.AddFocus(widget.FOCUS_SOUTH, btnOptions)
	btnOptions.AddFocus(widget.FOCUS_NORTH, btnStart)

	s.ui.Container.AddChild(outerContainer)
}
