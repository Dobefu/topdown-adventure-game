package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
)

type Game interface {
}

type game struct {
	Game

	world donburi.World
}

func NewGame() *game {
	return &game{
		world: donburi.NewWorld(),
	}
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, screen.Bounds().Max.String())
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
