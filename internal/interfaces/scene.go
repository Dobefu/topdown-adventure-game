package interfaces

import (
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
	GetCollisionTile(x, y int) int
	AddGameObject(gameObject GameObject)
}
