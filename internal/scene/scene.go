package scene

import (
	"log"
	"math"

	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
	"github.com/setanarut/kamera/v2"
)

type Scene struct {
	interfaces.Scene

	Game interfaces.Game

	camera           *kamera.Camera
	cameraTarget     *interfaces.GameObject
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

func (s *Scene) SetCamera(camera *kamera.Camera) {
	s.camera = camera
}

func (s *Scene) GetCamera() *kamera.Camera {
	return s.camera
}

func (s *Scene) SetCameraTarget(target interfaces.GameObject) {
	s.cameraTarget = &target
}

func (s *Scene) GetCameraTarget() interfaces.GameObject {
	if s.cameraTarget == nil {
		return nil
	}

	return *s.cameraTarget
}

func (s *Scene) AddGameObject(gameObject interfaces.GameObject) {
	s.gameObjects = append(s.gameObjects, gameObject)
	gameObject.SetScene(s)
}

func (s *Scene) GetCollisionTile(x float64, y float64) int {
	if s.sceneMap == nil || len(s.sceneMap.Layers) < 2 {
		return 0
	}

	posX := int(math.Ceil(x))
	posY := int(math.Ceil(y))
	collisionLayer := s.sceneMap.Layers[1]

	// If the position is out of bounds, assume there's a solid tile.
	if x < 0 ||
		y < 0 ||
		posX >= s.sceneMap.Width*s.sceneMap.TileWidth ||
		posY >= s.sceneMap.Height*s.sceneMap.TileHeight {

		return int(s.sceneMap.Tilesets[1].FirstGID)
	}

	tile := collisionLayer.Tiles[(posY/s.sceneMap.TileHeight)*s.sceneMap.Width+(posX/s.sceneMap.TileWidth)]

	if tile == nil {
		return 0
	}

	// If the tile has a tileset, return the GID of the tile.
	if tile.Tileset != nil {
		return int(tile.ID + tile.Tileset.FirstGID)
	}

	return int(tile.ID)
}
