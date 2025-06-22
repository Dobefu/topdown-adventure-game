package game_object

import (
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
)

func (g *GameObject) GetScene() *interfaces.Scene {
	return g.scene
}
