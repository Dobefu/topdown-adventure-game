package game_object

import (
	"math"

	"github.com/Dobefu/vectors"
)

func (g *GameObject) Move(velocity vectors.Vector3) (newVelocity vectors.Vector3) {
	pos := g.GetPosition()

	if g.canMoveTo(velocity) {
		pos.Add(velocity)
		g.SetPosition(*pos)

		return velocity
	}

	if g.canMoveTo(vectors.Vector3{X: velocity.X, Y: 0, Z: 0}) {
		pos.Add(vectors.Vector3{X: velocity.X, Y: 0, Z: 0})
		pos.Y = math.Round(pos.Y/32) * 32
	} else if g.canMoveTo(vectors.Vector3{X: 0, Y: velocity.Y, Z: 0}) {
		pos.Add(vectors.Vector3{X: 0, Y: velocity.Y, Z: 0})
		pos.X = math.Round(pos.X/32) * 32
	} else {
		pos.X = math.Round(pos.X/32) * 32
		pos.Y = math.Round(pos.Y/32) * 32
	}

	g.SetPosition(*pos)

	return velocity
}

func (g *GameObject) canMoveTo(velocity vectors.Vector3) bool {
	pos := g.GetPosition()
	scene := *g.GetScene()

	targetX := pos.X + velocity.X
	targetY := pos.Y + velocity.Y

	topLeft := scene.GetCollisionTile(velocity, targetX, targetY)
	topRight := scene.GetCollisionTile(velocity, targetX+31, targetY)
	bottomLeft := scene.GetCollisionTile(velocity, targetX, targetY+31)
	bottomRight := scene.GetCollisionTile(velocity, targetX+31, targetY+31)

	return topLeft == 0 && topRight == 0 && bottomLeft == 0 && bottomRight == 0
}
