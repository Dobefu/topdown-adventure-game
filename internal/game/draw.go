package game

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *game) Draw(screen *ebiten.Image) {
	screenSize := screen.Bounds().Size()
	g.screenWidth = screenSize.X
	g.screenHeight = screenSize.Y

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
		screen.Fill(color.Black)

		err := sceneMapRenderer.RenderLayer(0)

		if err != nil {
			log.Fatal(err)
		}

		g.cachedLayerImages[0].WritePixels(sceneMapRenderer.Result.Pix)
		camera.Draw(g.cachedLayerImages[0], &ebiten.DrawImageOptions{}, screen)
		sceneMapRenderer.Clear()

		gameObjects := g.scene.GetGameObjects()

		for _, gameObject := range gameObjects {
			if !gameObject.GetIsActive() {
				continue
			}

			gameObject.DrawBelow(screen)
		}

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

		for _, gameObject := range gameObjects {
			if !gameObject.GetIsActive() {
				continue
			}

			gameObject.DrawUI(g.cachedUIImg)
		}

		UIImgOptions := &ebiten.DrawImageOptions{}
		UIImgOptions.GeoM.Scale(widthScale, heightScale)
		screen.DrawImage(g.cachedUIImg, UIImgOptions)
	}

	gameObjects := g.scene.GetGameObjects()

	for _, gameObject := range gameObjects {
		if !gameObject.GetIsActive() {
			continue
		}

		gameObject.Draw(screen)
	}

	if sceneMap != nil {
		err := sceneMapRenderer.RenderLayer(2)

		if err != nil {
			log.Fatal(err)
		}

		g.cachedLayerImages[2].WritePixels(sceneMapRenderer.Result.Pix)
		camera.Draw(g.cachedLayerImages[2], &ebiten.DrawImageOptions{}, screen)
		sceneMapRenderer.Clear()
	}

	g.scene.GetUI().Draw(screen)
}
