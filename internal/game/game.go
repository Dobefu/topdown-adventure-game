package game

import (
	"fmt"
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

	CAMERA_SMOOTHING = .1
)

type game struct {
	interfaces.Game

	isDebugEnabled bool
	isDebugActive  bool

	input *ebitengine_input.Handler

	scene interfaces.Scene

	screenWidth  int
	screenHeight int

	cachedLayerImages []*ebiten.Image
	cachedUIImg       *ebiten.Image
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

	g.cachedUIImg = ebiten.NewImage(VIRTUAL_WIDTH, VIRTUAL_HEIGHT)

	return g
}

func (g *game) GetScene() (scene interfaces.Scene) {
	return g.scene
}

func (g *game) SetScene(scene interfaces.Scene) {
	g.scene = scene

	camera := kamera.NewCamera(
		-float64(g.screenWidth)/2,
		-float64(g.screenHeight)/2,
		float64(g.screenWidth),
		float64(g.screenHeight),
	)

	g.scene.SetGame(g)
	g.scene.SetCamera(camera)

	camera.ShakeEnabled = true
	camera.SmoothType = kamera.SmoothDamp
	camera.SmoothOptions = kamera.DefaultSmoothOptions()
	camera.SmoothOptions.SmoothDampTimeX = CAMERA_SMOOTHING
	camera.SmoothOptions.SmoothDampTimeY = CAMERA_SMOOTHING

	widthScale := float64(g.screenWidth) / VIRTUAL_WIDTH
	heightScale := float64(g.screenHeight) / VIRTUAL_WIDTH
	camera.ZoomFactor = math.Min(widthScale, heightScale)

	g.scene.Init()
	g.scene.InitUI()

	sceneMap, _ := g.scene.GetSceneMapData()

	if sceneMap != nil {
		for range 3 {
			g.cachedLayerImages = append(
				g.cachedLayerImages,
				ebiten.NewImage(
					sceneMap.Width*sceneMap.TileWidth,
					sceneMap.Height*sceneMap.TileHeight,
				),
			)
		}
	}
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
		cameraTargetPosition := cameraTarget.GetCameraPosition()

		camera.LookAt(
			cameraTargetPosition.X-camera.Width/2,
			cameraTargetPosition.Y-camera.Height/2,
		)
	}

	g.scene.GetUI().Update()

	activeGameObjects := g.scene.GetActiveGameObjects()

	for _, gameObject := range activeGameObjects {
		err = gameObject.Update()

		if err != nil {
			return err
		}
	}

	return nil
}

func (g *game) DrawFinalScreen(screen ebiten.FinalScreen, offscreen *ebiten.Image, geoM ebiten.GeoM) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM = geoM

	screen.DrawImage(offscreen, op)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
