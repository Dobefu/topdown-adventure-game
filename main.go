package main

import (
	"log"

	"github.com/Dobefu/topdown-adventure-game/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Top-down Adventure Game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	ebiten.SetWindowSize(1280, 720)

	ebiten.SetScreenClearedEveryFrame(false)

	gameOptions := ebiten.RunGameOptions{
		DisableHiDPI: true,
	}

	err := ebiten.RunGameWithOptions(game.NewGame(), &gameOptions)

	if err != nil {
		log.Fatal(err)
	}
}
