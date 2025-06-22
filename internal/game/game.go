package game

import (
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/topdown-adventure-game/internal/scene"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game interface {
	ebiten.Game

	GetScene() (scene interfaces.Scene)
	SetScene(scene interfaces.Scene)
}

type game struct {
	Game

	scene interfaces.Scene
}

func NewGame() (g *game) {
	g = &game{}
	g.SetScene(&scene.Level1Scene{})

	return g
}

func (g *game) SetScene(scene interfaces.Scene) {
	g.scene = scene

	g.scene.Init()
}

func (g *game) Update() (err error) {
	if g.scene == nil {
		return nil
	}

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
