package interfaces

import (
	"github.com/Dobefu/vectors"
	"github.com/ebitenui/ebitenui"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
	"github.com/setanarut/kamera/v2"
)

// Scene defines the interface for scene management functionality.
type Scene interface {
	// Initialize the base scene.
	Init()

	// Initialize the scene UI.
	InitUI()

	// Initialize the pause screen UI.
	InitPauseScreenUI()

	// Initialize the scene map.
	InitSceneMap(path string)

	// Set the game instance.
	SetGame(game Game)

	// Get the game instance.
	GetGame() (game Game)

	// Get the game objects.
	GetGameObjects() []GameObject

	// Get the active game objects.
	GetActiveGameObjects() []GameObject

	// Get the map data of the scene.
	GetSceneMapData() (sceneMap *tiled.Map, sceneMapRenderer *render.Renderer)

	// Get the UI of the scene.
	GetUI() *ebitenui.UI

	// Get the pause screen UI of the scene.
	GetPauseScreenUI() *ebitenui.UI

	// Set whether or not the scene can be paused.
	SetCanPause(canPause bool)

	// Get whether or not the scene can be paused.
	GetCanPause() (canPause bool)

	// Set whether or not the scene is paused.
	SetIsPaused(isPaused bool)

	// Get the pause state of the scene.
	GetIsPaused() (isPaused bool)

	// Set the camera of the scene.
	SetCamera(camera *kamera.Camera)

	// Get the camera of the scene.
	GetCamera() *kamera.Camera

	// Set the camera target of the scene.
	SetCameraTarget(camera GameObject)

	// Get the camera target of the scene.
	GetCameraTarget() GameObject

	// Get the collision tile of the scene.
	GetCollisionTile(velocity vectors.Vector3, position vectors.Vector2) int

	// Add a game object to the scene.
	AddGameObject(gameObject GameObject)

	// Remove a game object from the scene.
	RemoveGameObject(gameObject GameObject)
}
