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
	s.InitSceneMap("maps/overworld.tmx")

	player := player.NewPlayer(vectors.Vector3{X: 0, Y: 0, Z: 0})

	s.AddGameObject(player)
	s.SetCameraTarget(player)
}

func (s *OverworldScene) InitUI() {
	s.Scene.InitUI()
}
