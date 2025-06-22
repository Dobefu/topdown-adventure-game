package scene

import (
	"log"

	"github.com/Dobefu/topdown-adventure-game/internal/player"
	"github.com/Dobefu/vectors"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

type OverworldScene struct {
	Scene
}

func (s *OverworldScene) Init() {
	s.Scene.Init()

	sceneMap, err := tiled.LoadFile("maps/overworld.tmx", tiled.WithFileSystem(mapsFS))

	if err != nil {
		log.Fatal(err)
	}

	s.sceneMap = sceneMap
	sceneMapRenderer, err := render.NewRendererWithFileSystem(s.sceneMap, mapsFS)

	if err != nil {
		log.Fatal(err)
	}

	s.sceneMapRenderer = sceneMapRenderer

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
