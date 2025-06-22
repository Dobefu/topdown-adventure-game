package game_object

import (
	"github.com/Dobefu/topdown-adventure-game/internal/interfaces"
)

func (g *GameObject) SetScene(scene interfaces.Scene) {
	g.scene = &scene
}
