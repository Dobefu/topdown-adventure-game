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
	Init()
	InitUI()
	InitPauseScreenUI()
	InitSceneMap(path string)
	SetGame(game Game)
	GetGame() (game Game)
	GetGameObjects() []GameObject
	GetActiveGameObjects() []GameObject
	GetSceneMapData() (sceneMap *tiled.Map, sceneMapRenderer *render.Renderer)
	GetUI() *ebitenui.UI
	GetPauseScreenUI() *ebitenui.UI
	SetCanPause(canPause bool)
	GetCanPause() (canPause bool)
	SetIsPaused(isPaused bool)
	GetIsPaused() (isPaused bool)
	SetCamera(camera *kamera.Camera)
	GetCamera() *kamera.Camera
	SetCameraTarget(camera GameObject)
	GetCameraTarget() GameObject
	GetCollisionTile(velocity vectors.Vector3, position vectors.Vector2) int
	AddGameObject(gameObject GameObject)
	RemoveGameObject(gameObject GameObject)
}
