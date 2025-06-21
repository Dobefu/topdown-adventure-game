package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, screen.Bounds().Max.String())
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowTitle("Top-down Adventure Game")

	gameOptions := ebiten.RunGameOptions{}

	err := ebiten.RunGameWithOptions(&Game{}, &gameOptions)

	if err != nil {
		log.Fatal(err)
	}
}
