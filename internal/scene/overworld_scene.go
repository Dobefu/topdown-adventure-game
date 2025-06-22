package scene

import (
	"github.com/Dobefu/topdown-adventure-game/internal/player"
	"github.com/Dobefu/vectors"
)

type OverworldScene struct {
	Scene
}

func (s *OverworldScene) Init() {
	s.Scene.Init()

	s.AddGameObject(
		player.NewPlayer(vectors.Vector3{
			X: 10,
			Y: 10,
			Z: 0,
		}))
}

func (s *OverworldScene) InitUI() {
	s.Scene.InitUI()
}
