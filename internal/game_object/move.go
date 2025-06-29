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

	target := vectors.Vector2{
		X: pos.X + velocity.X,
		Y: pos.Y + velocity.Y,
	}

	topLeft, _ := scene.GetCollisionTile(velocity, vectors.Vector2{X: target.X, Y: target.Y})
	topRight, _ := scene.GetCollisionTile(velocity, vectors.Vector2{X: target.X + 31, Y: target.Y})
	bottomLeft, _ := scene.GetCollisionTile(velocity, vectors.Vector2{X: target.X, Y: target.Y + 31})
	bottomRight, _ := scene.GetCollisionTile(velocity, vectors.Vector2{X: target.X + 31, Y: target.Y + 31})

	return topLeft == 0 && topRight == 0 && bottomLeft == 0 && bottomRight == 0
}
