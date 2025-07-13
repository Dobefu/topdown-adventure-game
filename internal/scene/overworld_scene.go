package scene

import (
	"image/color"

	"github.com/Dobefu/topdown-adventure-game/internal/enemy"
	"github.com/Dobefu/topdown-adventure-game/internal/player"
	"github.com/Dobefu/topdown-adventure-game/internal/ui"
	"github.com/Dobefu/vectors"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type OverworldScene struct {
	Scene
}

func (s *OverworldScene) Init() {
	s.canPause = true

	s.Scene.Init()
	s.InitSceneMap("maps/overworld.tmx")

	player := player.NewPlayer(vectors.Vector3{X: 0, Y: 0, Z: 0})
	s.AddGameObject(player)
	s.SetCameraTarget(player)

	enemy := enemy.NewEnemy(vectors.Vector3{X: 160, Y: 240, Z: 0})
	s.AddGameObject(enemy)
}

func (s *OverworldScene) InitUI() {
	s.Scene.InitUI()
}

func (s *OverworldScene) InitPauseScreenUI() {
	s.Scene.InitPauseScreenUI()

	btnContinue := ui.NewButton(
		"Continue",
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			s.SetIsPaused(false)
		}),
	)

	btnMainMenu := ui.NewButton(
		"Main Menu",
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			s.game.SetScene(&MainMenuScene{})
		}),
	)

	outerContainer := ui.NewContainer(
		64,
		32,
		32,
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(color.RGBA{R: 32, G: 32, B: 32, A: 128}),
		),
	)

	outerContainer.AddChild(ui.NewTitle("Paused"))

	innerContainer := ui.NewContainer(
		16,
		0,
		0,
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)
	outerContainer.AddChild(innerContainer)

	innerContainer.AddChild(btnContinue)
	innerContainer.AddChild(btnMainMenu)

	btnContinue.AddFocus(widget.FOCUS_SOUTH, btnMainMenu)
	btnMainMenu.AddFocus(widget.FOCUS_NORTH, btnContinue)

	s.pauseScreenUi.Container.AddChild(outerContainer)
}
