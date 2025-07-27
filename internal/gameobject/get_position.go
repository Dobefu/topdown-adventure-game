package gameobject

import (
	"github.com/Dobefu/vectors"
)

// GetPosition gets the current position of the game object.
func (g *GameObject) GetPosition() (position *vectors.Vector3) {
	return &g.Position
}
