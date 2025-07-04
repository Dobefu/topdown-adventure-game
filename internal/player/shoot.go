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
		velocity := cameraPos.Normalize()
		velocity.Mul(vectors.Vector3{
			X: 10,
			Y: 10,
			Z: 1,
		})

		b.SetPosition(*p.GetPosition())
		b.SetVelocity(velocity)

		// Skip firing if the bullet would remain still, just in case.
		if !velocity.IsZero() {
			b.SetIsActive(true)
			p.shootCooldown = p.shootCooldownMax
		}
		break
	}
}
