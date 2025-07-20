package gameobject

import (
	"github.com/Dobefu/vectors"
)

// GetCameraPosition gets the position that a camera should be following.
func (g *GameObject) GetCameraPosition() (position *vectors.Vector3) {
	return g.GetPosition()
}
