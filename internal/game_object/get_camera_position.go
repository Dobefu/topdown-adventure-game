package game_object

import (
	"github.com/Dobefu/vectors"
)

func (g *GameObject) GetCameraPosition() (position *vectors.Vector3) {
	return g.GetPosition()
}
