package game

import (
	"image/color"
	"log"
	"math"

	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
	"github.com/hajimehoshi/ebiten/v2"
)

// Draw handles drawing of all active gameobjects in a scene.
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	screenSize := screen.Bounds().Size()
	g.screenWidth = screenSize.X
	g.screenHeight = screenSize.Y

	if g.scene == nil {
		return
	}

	camera := g.scene.GetCamera()
	camera.SetSize(float64(g.screenWidth), float64(g.screenHeight))

	widthScale := float64(g.screenWidth) / VirtualWidth
	heightScale := float64(g.screenHeight) / VirtualHeight
	camera.ZoomFactor = math.Min(widthScale, heightScale)

	activeGameObjects := g.scene.GetActiveGameObjects()
	sceneMap, sceneMapRenderer := g.scene.GetSceneMapData()

	if sceneMap != nil {
		err := sceneMapRenderer.RenderLayer(0)

		if err != nil {
			log.Fatal(err)
		}

		g.cachedLayerImages[0].WritePixels(sceneMapRenderer.Result.Pix)
		camera.Draw(g.cachedLayerImages[0], &ebiten.DrawImageOptions{}, screen)
		sceneMapRenderer.Clear()

		g.drawGameObjectsBelow(screen, activeGameObjects)

		err = sceneMapRenderer.RenderLayer(1)

		if err != nil {
			log.Fatal(err)
		}

		if g.isDebugActive {
			_ = sceneMapRenderer.RenderLayer(3)
		}

		g.cachedLayerImages[1].WritePixels(sceneMapRenderer.Result.Pix)
		camera.Draw(g.cachedLayerImages[1], &ebiten.DrawImageOptions{}, screen)
		sceneMapRenderer.Clear()
	}

	g.drawGameObjects(screen, activeGameObjects)
	g.drawGameObjectsAbove(screen, activeGameObjects)

	if sceneMap != nil {
		err := sceneMapRenderer.RenderLayer(2)

		if err != nil {
			log.Fatal(err)
		}

		g.cachedLayerImages[2].WritePixels(sceneMapRenderer.Result.Pix)
		camera.Draw(g.cachedLayerImages[2], &ebiten.DrawImageOptions{}, screen)
		sceneMapRenderer.Clear()
	}

	g.drawGameObjectsUI(screen, activeGameObjects)
	g.scene.GetUI().Draw(screen)

	if g.scene.GetIsPaused() {
		g.scene.GetPauseScreenUI().Draw(screen)
	}
}

func (g *Game) drawGameObjectsBelow(
	screen *ebiten.Image,
	activeGameObjects []interfaces.GameObject,
) {
	for _, gameObject := range activeGameObjects {
		gameObject.DrawBelow(screen)
	}
}

func (g *Game) drawGameObjects(
	screen *ebiten.Image,
	activeGameObjects []interfaces.GameObject,
) {
	for _, gameObject := range activeGameObjects {
		gameObject.Draw(screen)
	}
}

func (g *Game) drawGameObjectsAbove(
	screen *ebiten.Image,
	activeGameObjects []interfaces.GameObject,
) {
	for _, gameObject := range activeGameObjects {
		gameObject.DrawAbove(screen)
	}
}

func (g *Game) drawGameObjectsUI(
	screen *ebiten.Image,
	activeGameObjects []interfaces.GameObject,
) {
	for _, gameObject := range activeGameObjects {
		gameObject.DrawUI(screen)
	}
}
