package scene

import (
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
)

type Scene struct {
	interfaces.Scene

	gameObjects []interfaces.GameObject
	ui          *ebitenui.UI
}

func (s *Scene) Init() {
	s.InitUI()
}

func (s *Scene) InitUI() {
	s.ui = &ebitenui.UI{
		Container: widget.NewContainer(),
	}
}

func (s *Scene) GetGameObjects() []interfaces.GameObject {
	return s.gameObjects
}

func (s *Scene) AddGameObject(gameObject interfaces.GameObject) {
	s.gameObjects = append(s.gameObjects, gameObject)
}
