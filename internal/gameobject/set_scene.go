package gameobject

import (
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
)

// SetScene sets the scene that the game object is currently in.
func (g *GameObject) SetScene(scene interfaces.Scene) {
	g.scene = &scene
}
