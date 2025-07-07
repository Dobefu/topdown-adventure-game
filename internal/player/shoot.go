package player

import "github.com/Dobefu/vectors"

func (p *Player) Shoot() {
	for _, b := range p.bulletPool {
		if b.GetIsActive() {
			continue
		}

		cameraPos := *p.GetCameraPosition()
		pos := *p.GetPosition()
		pos.Add(vectors.Vector3{
			X: FRAME_WIDTH / 2,
			Y: FRAME_HEIGHT / 2,
			Z: 0,
		})

		cameraPos.Sub(pos)
		cameraPos.Normalize()
		cameraPos.Mul(vectors.Vector3{
			X: 10,
			Y: 10,
			Z: 1,
		})

		// Skip firing if the bullet would remain still, just in case.
		if !cameraPos.IsZero() {
			b.SetOwner(&p.GameObject)
			b.SetPosition(*p.GetPosition())
			b.SetVelocity(cameraPos)

			b.SetIsActive(true)
			p.shootCooldown = p.shootCooldownMax
		}
		break
	}
}
