package interfaces

import "github.com/ebitenui/ebitenui"

type Scene interface {
	Init()
	InitUI()
	SetGame(game Game)
	GetGameObjects() []GameObject
	GetUI() *ebitenui.UI
	AddGameObject(gameObject GameObject)
}
