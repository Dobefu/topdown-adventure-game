package main

import (
	"log"

	"github.com/Dobefu/topdown-adventure-game/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Top-down Adventure Game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// Set a minimum size limit.
	ebiten.SetWindowSizeLimits(320, 320, int(^uint16(0)), int(^uint16(0)))
	ebiten.SetScreenClearedEveryFrame(false)

	gameOptions := ebiten.RunGameOptions{}

	err := ebiten.RunGameWithOptions(game.NewGame(), &gameOptions)

	if err != nil {
		log.Fatal(err)
	}
}
