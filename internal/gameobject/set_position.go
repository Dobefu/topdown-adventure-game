package gameobject

import (
	"github.com/Dobefu/vectors"
)

// SetPosition sets the current position of a game object.
func (g *GameObject) SetPosition(position vectors.Vector3) {
	if g.position.Z == position.Z {
		g.position = position
		return
	}

	g.position = position
}
