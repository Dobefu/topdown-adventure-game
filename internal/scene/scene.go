package scene

import (
	"log"

	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

type Scene struct {
	interfaces.Scene

	Game interfaces.Game

	gameObjects      []interfaces.GameObject
	sceneMap         *tiled.Map
	sceneMapRenderer *render.Renderer
	ui               *ebitenui.UI
}

func (s *Scene) Init() {
	s.InitUI()
}

func (s *Scene) InitUI() {
	s.ui = &ebitenui.UI{
		Container: widget.NewContainer(
			widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		),
	}
}

func (s *Scene) InitSceneMap(path string) {
	sceneMap, err := tiled.LoadFile(path, tiled.WithFileSystem(mapsFS))

	if err != nil {
		log.Fatal(err)
	}

	s.sceneMap = sceneMap
	sceneMapRenderer, err := render.NewRendererWithFileSystem(s.sceneMap, mapsFS)

	if err != nil {
		log.Fatal(err)
	}

	s.sceneMapRenderer = sceneMapRenderer
}

func (s *Scene) SetGame(game interfaces.Game) {
	s.Game = game
}

func (s *Scene) GetGameObjects() []interfaces.GameObject {
	return s.gameObjects
}

func (s *Scene) GetSceneMapData() (sceneMap *tiled.Map, sceneMapRenderer *render.Renderer) {
	return s.sceneMap, s.sceneMapRenderer
}

func (s *Scene) GetUI() *ebitenui.UI {
	return s.ui
}

func (s *Scene) AddGameObject(gameObject interfaces.GameObject) {
	s.gameObjects = append(s.gameObjects, gameObject)
}
