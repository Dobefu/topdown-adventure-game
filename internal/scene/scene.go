package scene

import (
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
)

type Scene struct {
	interfaces.Scene

	gameObjects []interfaces.GameObject
}

func (s *Scene) Init() {
	// noop
}

func (s *Scene) GetGameObjects() []interfaces.GameObject {
	return s.gameObjects
}

func (s *Scene) AddGameObject(gameObject interfaces.GameObject) {
	s.gameObjects = append(s.gameObjects, gameObject)
}
