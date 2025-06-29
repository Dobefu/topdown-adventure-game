package game

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/topdown-adventure-game/internal/scene"
	"github.com/Dobefu/topdown-adventure-game/internal/storage"
	"github.com/hajimehoshi/ebiten/v2"
	ebitengine_input "github.com/quasilyte/ebitengine-input"
	"github.com/setanarut/kamera/v2"
)

const (
	VIRTUAL_WIDTH  = 640
	VIRTUAL_HEIGHT = 360

	CAMERA_SMOOTHING = .2
)

type game struct {
	interfaces.Game

	isDebugEnabled bool
	isDebugActive  bool

	input *ebitengine_input.Handler

	scene interfaces.Scene

	screenWidth  int
	screenHeight int
}

func NewGame(isDebugEnabled bool) (g *game) {
	g = &game{
		isDebugEnabled: isDebugEnabled,
	}

	if isDebugEnabled {
		val, err := storage.GetOption("isDebugActive")

		if err != nil {
			log.Fatal(err)
		}

		g.isDebugActive = val == "true"
	}

	g.input = input.Input.NewHandler(255, input.Keymap)

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

	camera.SmoothType = kamera.Lerp
	camera.SmoothOptions = kamera.DefaultSmoothOptions()
	camera.SmoothOptions.LerpSpeedX = CAMERA_SMOOTHING
	camera.SmoothOptions.LerpSpeedY = CAMERA_SMOOTHING

	widthScale := float64(g.screenWidth) / VIRTUAL_WIDTH
	heightScale := float64(g.screenHeight) / VIRTUAL_WIDTH
	camera.ZoomFactor = math.Min(widthScale, heightScale)

	g.scene.Init()
	g.scene.InitUI()
}

func (g *game) Update() (err error) {
	input.Input.Update()

	if g.input.ActionIsJustPressed(input.ActionToggleDebug) {
		g.isDebugActive = !g.isDebugActive
		err = storage.SetOption("isDebugActive", fmt.Sprintf("%v", g.isDebugActive))

		if err != nil {
			return err
		}
	}

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

	widthScale := float64(g.screenWidth) / VIRTUAL_WIDTH
	heightScale := float64(g.screenHeight) / VIRTUAL_HEIGHT
	camera.ZoomFactor = math.Min(widthScale, heightScale)

	sceneMap, sceneMapRenderer := g.scene.GetSceneMapData()

	if sceneMap != nil {
		err := sceneMapRenderer.RenderLayer(0)

		if err != nil {
			log.Fatal(err)
		}

		if g.isDebugActive {
			_ = sceneMapRenderer.RenderLayer(2)
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

	if sceneMap != nil {
		err := sceneMapRenderer.RenderLayer(1)

		if err != nil {
			log.Fatal(err)
		}

		camera.Draw(ebiten.NewImageFromImage(sceneMapRenderer.Result), &ebiten.DrawImageOptions{}, screen)
		sceneMapRenderer.Clear()
	}
}

func (g *game) DrawFinalScreen(screen ebiten.FinalScreen, offscreen *ebiten.Image, geoM ebiten.GeoM) {
	screen.DrawImage(offscreen, nil)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
