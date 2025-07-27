// Package game provides the main game.
package game

import (
	"fmt"
	"log"
	"math"

	"github.com/Dobefu/topdown-adventure-game/internal/input"
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/Dobefu/topdown-adventure-game/internal/scene"
	"github.com/Dobefu/topdown-adventure-game/internal/storage"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	ebitengine_input "github.com/quasilyte/ebitengine-input"
	"github.com/setanarut/kamera/v2"
)

const (
	// VirtualWidth is the virtual width to draw the game at.
	VirtualWidth = 640
	// VirtualHeight is the virtual height to draw the game at.
	VirtualHeight = 360

	// CameraSmoothing is the smooth factor of the camera. Range: 0.01 - 1.
	CameraSmoothing = .1
)

// Game struct provides the game.
type Game struct {
	isDebugEnabled bool
	isDebugActive  bool

	audioContext *audio.Context

	scale float64

	input *ebitengine_input.Handler

	scene interfaces.Scene

	screenWidth  int
	screenHeight int

	cachedLayerImages []*ebiten.Image
}

// NewGame creates a new game instance.
func NewGame(isDebugEnabled bool) (g *Game) {
	g = &Game{
		isDebugEnabled: isDebugEnabled,
	}

	if isDebugEnabled {
		val, err := storage.GetOption("isDebugActive", false)

		if err != nil {
			log.Fatal(err)
		}

		g.isDebugActive = val
	}

	g.audioContext = audio.NewContext(48000)

	g.input = input.Input.NewHandler(0, input.UIKeymap)

	g.SetScene(&scene.MainMenuScene{})

	return g
}

// GetIsDebugEnabled gets whether or not debugging is enabled.
func (g *Game) GetIsDebugEnabled() (isDebugEnabled bool) {
	return g.isDebugEnabled
}

// GetIsDebugActive gets whether or not debugging is active.
func (g *Game) GetIsDebugActive() (isDebugActive bool) {
	return g.isDebugActive
}

// GetScene gets the current active scene.
func (g *Game) GetScene() (scene interfaces.Scene) {
	return g.scene
}

// SetScene sets the currently active scene.
func (g *Game) SetScene(scene interfaces.Scene) {
	g.scene = scene

	camera := kamera.NewCamera(
		16,
		16,
		float64(g.screenWidth),
		float64(g.screenHeight),
	)

	g.scene.SetGame(g)
	g.scene.SetCamera(camera)

	camera.ShakeEnabled = true
	camera.SmoothType = kamera.SmoothDamp
	camera.SmoothOptions = kamera.DefaultSmoothOptions()
	camera.SmoothOptions.SmoothDampTimeX = CameraSmoothing
	camera.SmoothOptions.SmoothDampTimeY = CameraSmoothing

	widthScale := float64(g.screenWidth) / VirtualWidth
	heightScale := float64(g.screenHeight) / VirtualWidth
	camera.ZoomFactor = math.Min(widthScale, heightScale)

	g.scene.Init()
	g.scene.InitUI()
	g.scene.InitPauseScreenUI()

	ui := g.scene.GetUI()
	ui.ChangeFocus(widget.FOCUS_SOUTH)

	// Update the UI immediately, so the new focused element gets set.
	// Otherwise, two inputs will be needed in order to change focus on a controller.
	ui.Update()

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

// GetAudioContext gets the current audio context.
func (g *Game) GetAudioContext() (audioContext *audio.Context) {
	return g.audioContext
}

// Update handles updating of all active gameobjects in a scene.
func (g *Game) Update() (err error) {
	input.Input.Update()
	g.UpdateUIInput()

	camera := g.scene.GetCamera()

	widthScale := float64(g.screenWidth) / VirtualWidth
	heightScale := float64(g.screenHeight) / VirtualWidth
	g.scale = math.Min(widthScale, heightScale)

	if g.isDebugEnabled && g.input.ActionIsJustPressed(input.ActionToggleDebug) {
		g.isDebugActive = !g.isDebugActive
		err = storage.SetOption("isDebugActive", fmt.Sprintf("%v", g.isDebugActive))

		if err != nil {
			return err
		}
	}

	if g.scene == nil {
		return nil
	}

	cameraTarget := g.scene.GetCameraTarget()

	if camera != nil && cameraTarget != nil {
		cameraTargetPosition := cameraTarget.GetCameraPosition()

		camera.LookAt(
			cameraTargetPosition.X,
			cameraTargetPosition.Y,
		)
	}

	if g.scene.GetCanPause() && g.input.ActionIsJustPressed(input.ActionPause) {
		g.scene.SetIsPaused(!g.scene.GetIsPaused())
	}

	if g.scene.GetIsPaused() {
		g.scene.GetPauseScreenUI().Update()

		return nil
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

// DrawFinalScreen draws the final screen.
func (g *Game) DrawFinalScreen(screen ebiten.FinalScreen, offscreen *ebiten.Image, geoM ebiten.GeoM) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM = geoM

	screen.DrawImage(offscreen, op)
}

// Layout determines the size of the rendered game.
func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return VirtualWidth, VirtualHeight
}
