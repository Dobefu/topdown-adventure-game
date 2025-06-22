package game

import (
	"image/color"
	"log"
	"math"

	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/topdown-adventure-game/internal/scene"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/setanarut/kamera/v2"
)

type game struct {
	interfaces.Game

	scene interfaces.Scene

	screenWidth  int
	screenHeight int
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

	camera := kamera.NewCamera(0, 0, float64(g.screenWidth), float64(g.screenHeight))
	g.scene.SetGame(g)
	g.scene.SetCamera(camera)

	widthScale := float64(g.screenWidth) / 640
	heightScale := float64(g.screenHeight) / 360
	camera.ZoomFactor = math.Min(widthScale, heightScale)

	g.scene.Init()
	g.scene.InitUI()
}

func (g *game) Update() (err error) {
	input.Input.Update()

	if g.scene == nil {
		return nil
	}

	camera := g.scene.GetCamera()
	cameraTarget := g.scene.GetCameraTarget()

	if camera != nil && cameraTarget != nil {
		camera.LookAt(cameraTarget.GetPosition().X, cameraTarget.GetPosition().Y)
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
	g.screenWidth = screen.Bounds().Size().X
	g.screenHeight = screen.Bounds().Size().Y

	if g.scene == nil {
		return
	}

	camera := g.scene.GetCamera()
	camera.SetSize(float64(g.screenWidth), float64(g.screenHeight))

	widthScale := float64(g.screenWidth) / 640
	heightScale := float64(g.screenHeight) / 360
	camera.ZoomFactor = math.Min(widthScale, heightScale)

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

		camera.Draw(ebiten.NewImageFromImage(sceneMapRenderer.Result), &ebiten.DrawImageOptions{}, screen)
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
