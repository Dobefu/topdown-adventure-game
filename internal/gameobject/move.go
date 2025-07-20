package gameobject

import (
	"math"

	"github.com/Dobefu/vectors"
)

func (g *GameObject) Move(
	velocity vectors.Vector3,
) {
	(*g.GetPosition()).Add(velocity)
}

func (g *GameObject) MoveWithCollisionRect(
	velocity vectors.Vector3,
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
) (newVelocity vectors.Vector3, hasCollided bool) {
	pos := g.GetPosition()

	if g.canMoveTo(velocity, x1, y1, x2, y2) {
		pos.Add(velocity)

		if pos.Z < 0 {
			pos.Z = 0
			velocity.Z = 0
		}

		g.SetPosition(*pos)

		return velocity, false
	}

	if g.canMoveTo(vectors.Vector3{X: velocity.X, Y: 0, Z: 0}, x1, y1, x2, y2) {
		pos.Add(vectors.Vector3{X: velocity.X, Y: 0, Z: 0})
		g.SetPosition(*pos)
	} else if velocity.X != 0 {
		maxX := g.findMaxMovement(
			vectors.Vector3{X: velocity.X, Y: 0, Z: 0},
			x1,
			y1,
			x2,
			y2,
		)

		if maxX != 0 {
			pos.Add(vectors.Vector3{X: maxX, Y: 0, Z: 0})
			g.SetPosition(*pos)
		}
	}

	pos = g.GetPosition()

	if g.canMoveTo(vectors.Vector3{X: 0, Y: velocity.Y, Z: 0}, x1, y1, x2, y2) {
		pos.Add(vectors.Vector3{X: 0, Y: velocity.Y, Z: 0})
	} else if velocity.Y != 0 {
		maxY := g.findMaxMovement(
			vectors.Vector3{X: 0, Y: velocity.Y, Z: 0},
			x1,
			y1,
			x2,
			y2,
		)

		if maxY != 0 {
			pos.Add(vectors.Vector3{X: 0, Y: maxY, Z: 0})
		}
	}

	g.SetPosition(*pos)

	return velocity, true
}

func (g *GameObject) findMaxMovement(
	velocity vectors.Vector3,
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
) float64 {
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

		testVelocity := getTestVelocityFromVelocity(velocity, target, center)

		if g.canMoveTo(testVelocity, x1, y1, x2, y2) {
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

func (g *GameObject) canMoveTo(
	velocity vectors.Vector3,
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
) bool {
	pos := g.GetPosition()
	scene := *g.GetScene()

	target := vectors.Vector2{
		X: pos.X + velocity.X,
		Y: pos.Y + velocity.Y,
	}

	topLeft := scene.GetCollisionTile(velocity, vectors.Vector2{X: target.X + x1, Y: target.Y + y1})
	if topLeft != 0 {
		return false
	}

	topRight := scene.GetCollisionTile(velocity, vectors.Vector2{X: target.X + x2, Y: target.Y + y1})
	if topRight != 0 {
		return false
	}

	bottomLeft := scene.GetCollisionTile(velocity, vectors.Vector2{X: target.X + x1, Y: target.Y + y2})
	if bottomLeft != 0 {
		return false
	}

	bottomRight := scene.GetCollisionTile(velocity, vectors.Vector2{X: target.X + x2, Y: target.Y + y2})

	return bottomRight == 0
}

func getTestVelocityFromVelocity(
	velocity vectors.Vector3,
	target float64,
	center float64,
) (testVelocity vectors.Vector3) {
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

	return testVelocity
}
