package game_object

import (
	"github.com/Dobefu/vectors"
)

func (g *GameObject) GetPosition() (position *vectors.Vector3) {
	return &g.position
}
