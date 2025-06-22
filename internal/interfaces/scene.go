package interfaces

import (
	"github.com/ebitenui/ebitenui"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

type Scene interface {
	Init()
	InitUI()
	InitSceneMap(path string)
	SetGame(game Game)
	GetGameObjects() []GameObject
	GetSceneMapData() (sceneMap *tiled.Map, sceneMapRenderer *render.Renderer)
	GetUI() *ebitenui.UI
	AddGameObject(gameObject GameObject)
}
