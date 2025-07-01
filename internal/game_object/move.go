package game_object

import (
	"math"

	"github.com/Dobefu/vectors"
)

func (g *GameObject) Move(
	velocity vectors.Vector3,
) (newVelocity vectors.Vector3, hasCollided bool) {
	pos := g.GetPosition()

	if g.canMoveTo(velocity) {
		pos.Add(velocity)
		g.SetPosition(*pos)

		return velocity, false
	}

	if g.canMoveTo(vectors.Vector3{X: velocity.X, Y: 0, Z: 0}) {
		pos.Add(vectors.Vector3{X: velocity.X, Y: 0, Z: 0})
		g.SetPosition(*pos)
	} else if velocity.X != 0 {
		maxX := g.findMaxMovement(vectors.Vector3{X: velocity.X, Y: 0, Z: 0})

		if maxX != 0 {
			pos.Add(vectors.Vector3{X: maxX, Y: 0, Z: 0})
			g.SetPosition(*pos)
		}
	}

	pos = g.GetPosition()

	if g.canMoveTo(vectors.Vector3{X: 0, Y: velocity.Y, Z: 0}) {
		pos.Add(vectors.Vector3{X: 0, Y: velocity.Y, Z: 0})
	} else if velocity.Y != 0 {
		maxY := g.findMaxMovement(vectors.Vector3{X: 0, Y: velocity.Y, Z: 0})

		if maxY != 0 {
			pos.Add(vectors.Vector3{X: 0, Y: maxY, Z: 0})
		}
	}

	g.SetPosition(*pos)

	return velocity, true
}

func (g *GameObject) findMaxMovement(velocity vectors.Vector3) float64 {
	var target float64

	if velocity.X != 0 {
		target = velocity.X
	} else {
		target = velocity.Y
	}

	var minDistance float64 = 0
	maxDistance := math.Abs(target)
	threshold := 0.1

	for maxDistance-minDistance > threshold {
		center := (minDistance + maxDistance) / 2

		var testVelocity vectors.Vector3

		if velocity.X != 0 {
			if target > 0 {
				testVelocity = vectors.Vector3{X: center, Y: 0, Z: 0}
			} else {
				testVelocity = vectors.Vector3{X: -center, Y: 0, Z: 0}
			}
		} else {
			if target > 0 {
				testVelocity = vectors.Vector3{X: 0, Y: center, Z: 0}
			} else {
				testVelocity = vectors.Vector3{X: 0, Y: -center, Z: 0}
			}
		}

		if g.canMoveTo(testVelocity) {
			minDistance = center
		} else {
			maxDistance = center
		}
	}

	if target > 0 {
		return minDistance
	} else {
		return -minDistance
	}
}

func (g *GameObject) canMoveTo(velocity vectors.Vector3) bool {
	pos := g.GetPosition()
	scene := *g.GetScene()

	target := vectors.Vector2{
		X: pos.X + velocity.X,
		Y: pos.Y + velocity.Y,
	}

	topLeft := scene.GetCollisionTile(velocity, vectors.Vector2{X: target.X + 4, Y: target.Y + 23})
	topRight := scene.GetCollisionTile(velocity, vectors.Vector2{X: target.X + 27, Y: target.Y + 23})
	bottomLeft := scene.GetCollisionTile(velocity, vectors.Vector2{X: target.X + 4, Y: target.Y + 31})
	bottomRight := scene.GetCollisionTile(velocity, vectors.Vector2{X: target.X + 27, Y: target.Y + 31})

	return topLeft == 0 && topRight == 0 && bottomLeft == 0 && bottomRight == 0
}
