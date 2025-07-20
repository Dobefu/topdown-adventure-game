package gameobject

import (
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
)

// GetScene gets the current scene that the game object is in.
func (g *GameObject) GetScene() *interfaces.Scene {
	return g.scene
}
