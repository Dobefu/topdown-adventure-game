// Package scene houses both the Scene struct and the scenes in the game.
package scene

import (
	"log"
	"math"
	"slices"

	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/topdown-adventure-game/internal/ui"
	"github.com/Dobefu/vectors"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
	"github.com/setanarut/kamera/v2"
)

// Scene defines a base scene instance.
// This is meant to be embedded, it should not be used directly.
type Scene struct {
	game interfaces.Game

	camera           *kamera.Camera
	cameraTarget     *interfaces.GameObject
	gameObjects      []interfaces.GameObject
	sceneMap         *tiled.Map
	sceneMapRenderer *render.Renderer
	ui               *ebitenui.UI
	pauseScreenUI    *ebitenui.UI
	canPause         bool
	isPaused         bool
}

// Init initializes the base scene.
func (s *Scene) Init() {
	s.InitUI()
	s.InitPauseScreenUI()
}

// InitUI initializes the base scene UI.
func (s *Scene) InitUI() {
	s.ui = &ebitenui.UI{
		Container: ui.NewAnchorContainer(0, 0),
	}
}

// InitPauseScreenUI initializes the pause screen UI.
func (s *Scene) InitPauseScreenUI() {
	s.pauseScreenUI = &ebitenui.UI{
		Container: ui.NewAnchorContainer(0, 0),
	}
}

// InitSceneMap initializes the scene map.
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

// SetGame sets the game instance.
func (s *Scene) SetGame(game interfaces.Game) {
	s.game = game
}

// GetGame gets the game instance.
func (s *Scene) GetGame() (game interfaces.Game) {
	return s.game
}

// GetGameObjects gets all the game objects in the scene.
func (s *Scene) GetGameObjects() []interfaces.GameObject {
	return s.gameObjects
}

// GetActiveGameObjects gets all the active game objects in the scene.
func (s *Scene) GetActiveGameObjects() (gameObjects []interfaces.GameObject) {
	for _, gameObject := range s.gameObjects {
		if !gameObject.GetIsActive() {
			continue
		}

		gameObjects = append(gameObjects, gameObject)
	}

	return gameObjects
}

// GetSceneMapData gets the map data of the scene.
func (s *Scene) GetSceneMapData() (sceneMap *tiled.Map, sceneMapRenderer *render.Renderer) {
	return s.sceneMap, s.sceneMapRenderer
}

// GetUI gets the UI of the scene.
func (s *Scene) GetUI() *ebitenui.UI {
	return s.ui
}

// GetPauseScreenUI gets the pause screen UI of the scene.
func (s *Scene) GetPauseScreenUI() *ebitenui.UI {
	return s.pauseScreenUI
}

// SetCanPause sets whether or not the scene can be paused.
func (s *Scene) SetCanPause(canPause bool) {
	s.canPause = canPause
}

// GetCanPause gets whether or not the scene can be paused.
func (s *Scene) GetCanPause() (canPause bool) {
	return s.canPause
}

// SetIsPaused sets whether or not the scene is paused.
func (s *Scene) SetIsPaused(isPaused bool) {
	s.isPaused = isPaused

	if isPaused && s.pauseScreenUI.GetFocusedWidget() == nil {
		s.pauseScreenUI.ChangeFocus(widget.FOCUS_NEXT)
	}
}

// GetIsPaused gets the current pause state of the scene.
func (s *Scene) GetIsPaused() (isPaused bool) {
	return s.isPaused
}

// SetCamera sets the scene's active camera.
func (s *Scene) SetCamera(camera *kamera.Camera) {
	s.camera = camera
}

// GetCamera gets the scene's active camera.
func (s *Scene) GetCamera() *kamera.Camera {
	return s.camera
}

// SetCameraTarget sets the current camera target of the scene.
func (s *Scene) SetCameraTarget(target interfaces.GameObject) {
	s.cameraTarget = &target
}

// GetCameraTarget gets the current camera target of the scene.
func (s *Scene) GetCameraTarget() interfaces.GameObject {
	if s.cameraTarget == nil {
		return nil
	}

	return *s.cameraTarget
}

// GetCollisionTile gets the collision tile in a scene at a position.
func (s *Scene) GetCollisionTile(velocity vectors.Vector3, position vectors.Vector2) int {
	if s.sceneMap == nil || len(s.sceneMap.Layers) < 4 {
		return 0
	}

	var posX, posY int

	if velocity.X > 0 {
		posX = int(math.Ceil(position.X))
	} else {
		posX = int(math.Floor(position.X))
	}

	if velocity.Y > 0 {
		posY = int(math.Ceil(position.Y))
	} else {
		posY = int(math.Floor(position.Y))
	}

	collisionLayer := s.sceneMap.Layers[3]

	// If the position is out of bounds, assume there's a solid tile.
	if position.X < 0 ||
		position.Y < 0 ||
		posX >= s.sceneMap.Width*s.sceneMap.TileWidth ||
		posY >= s.sceneMap.Height*s.sceneMap.TileHeight {

		return int(s.sceneMap.Tilesets[1].FirstGID)
	}

	tile := collisionLayer.Tiles[(posY/s.sceneMap.TileHeight)*s.sceneMap.Width+(posX/s.sceneMap.TileWidth)]

	if tile == nil {
		return 0
	}

	if tile.Tileset != nil {
		return int(tile.ID + tile.Tileset.FirstGID)
	}

	return int(tile.ID)
}

// AddGameObject adds a game object to the scene.
func (s *Scene) AddGameObject(gameObject interfaces.GameObject) {
	s.gameObjects = append(s.gameObjects, gameObject)
	gameObject.SetScene(s)
	gameObject.Init()
}

// RemoveGameObject removes a game object from the scene.
func (s *Scene) RemoveGameObject(gameObject interfaces.GameObject) {
	for idx, obj := range s.gameObjects {
		if obj.GetID() != gameObject.GetID() {
			continue
		}

		s.gameObjects = slices.Delete(s.gameObjects, idx, idx+1)
		return
	}
}
