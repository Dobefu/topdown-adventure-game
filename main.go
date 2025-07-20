// Package main provides the entrypoint for the game.
package main

import (
	"log"

	"github.com/Dobefu/topdown-adventure-game/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

// Main is the main function.
func Main(isDebugEnabled bool) {
	ebiten.SetWindowTitle("Top-down Adventure Game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	ebiten.SetWindowSize(1280, 720)

	ebiten.SetScreenClearedEveryFrame(false)

	gameOptions := ebiten.RunGameOptions{}

	err := ebiten.RunGameWithOptions(game.NewGame(isDebugEnabled), &gameOptions)

	if err != nil {
		log.Fatal(err)
	}
}
