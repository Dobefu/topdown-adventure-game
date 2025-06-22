package game

import (
	"image/color"
	"log"

	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/topdown-adventure-game/internal/scene"
	"github.com/hajimehoshi/ebiten/v2"
)

type game struct {
	interfaces.Game

	scene interfaces.Scene
}

func NewGame() (g *game) {
	g = &game{}
	g.SetScene(&scene.MainMenuScene{})

	return g
}

func (g *game) GetScene() (scene interfaces.Scene) {
	return g.scene
}

func (g *game) SetScene(scene interfaces.Scene) {
	g.scene = scene

	g.scene.SetGame(g)

	g.scene.Init()
	g.scene.InitUI()
}

func (g *game) Update() (err error) {
	input.Input.Update()

	if g.scene == nil {
		return nil
	}

	g.scene.GetUI().Update()

	gameObjects := g.scene.GetGameObjects()

	for _, gameObject := range gameObjects {
		if !gameObject.GetIsActive() {
			continue
		}

		err = gameObject.Update()

		if err != nil {
			return err
		}
	}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	if g.scene == nil {
		return
	}

	sceneMap, sceneMapRenderer := g.scene.GetSceneMapData()

	if sceneMap != nil {
		err := sceneMapRenderer.RenderLayer(0)

		if err != nil {
			log.Fatal(err)
		}

		err = sceneMapRenderer.RenderLayer(1)

		if err != nil {
			log.Fatal(err)
		}

		screen.Fill(color.Black)
		screen.DrawImage(ebiten.NewImageFromImage(sceneMapRenderer.Result), nil)
		sceneMapRenderer.Clear()
	} else {
		screen.Clear()
	}

	g.scene.GetUI().Draw(screen)

	gameObjects := g.scene.GetGameObjects()

	for _, gameObject := range gameObjects {
		if !gameObject.GetIsActive() {
			continue
		}

		gameObject.Draw(screen)
	}
}

func (g *game) DrawFinalScreen(screen ebiten.FinalScreen, offscreen *ebiten.Image, geoM ebiten.GeoM) {
	screen.DrawImage(offscreen, nil)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
