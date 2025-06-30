package player

import "github.com/Dobefu/vectors"

func (p *Player) GetCameraPosition() (position *vectors.Vector3) {
	position = p.GetPosition()

	return position
}
