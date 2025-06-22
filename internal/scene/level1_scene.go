package scene

import (
	"github.com/Dobefu/topdown-adventure-game/internal/player"
	"github.com/Dobefu/vectors"
)

type Level1Scene struct {
	Scene
}

func (s *Level1Scene) Init() {
	s.AddGameObject(
		player.NewPlayer(vectors.Vector3{
			X: 10,
			Y: 10,
			Z: 0,
		}))
}
