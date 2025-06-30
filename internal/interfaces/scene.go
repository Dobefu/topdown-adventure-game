package interfaces

import (
	"github.com/Dobefu/vectors"
	"github.com/ebitenui/ebitenui"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
	"github.com/setanarut/kamera/v2"
)

type Scene interface {
	Init()
	InitUI()
	InitSceneMap(path string)
	SetGame(game Game)
	GetGameObjects() []GameObject
	GetSceneMapData() (sceneMap *tiled.Map, sceneMapRenderer *render.Renderer)
	GetUI() *ebitenui.UI
	SetCamera(camera *kamera.Camera)
	GetCamera() *kamera.Camera
	SetCameraTarget(camera GameObject)
	GetCameraTarget() GameObject
	GetCollisionTile(velocity vectors.Vector3, position vectors.Vector2) int
	AddGameObject(gameObject GameObject)
}
